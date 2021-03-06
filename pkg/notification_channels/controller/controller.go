package controller

import (
	"github.com/personio/newrelic-alert-manager/internal"
	iov1alpha1 "github.com/personio/newrelic-alert-manager/pkg/apis/alerts/v1alpha1"
	"github.com/personio/newrelic-alert-manager/pkg/notification_channels/domain"
	"github.com/personio/newrelic-alert-manager/pkg/notification_channels/infrastructure/k8s"
	"github.com/personio/newrelic-alert-manager/pkg/notification_channels/infrastructure/newrelic"
	"github.com/go-logr/logr"
	"github.com/operator-framework/operator-sdk/pkg/predicate"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller-notification-channel")

// Reconcile reconciles a NotificationChannel object
type Reconcile struct {
	k8s      *k8s.Client
	logr     logr.Logger
	scheme   *runtime.Scheme
	newrelic *newrelic.ChannelRepository
}

func Add(mgr manager.Manager, controllerName string, channelType iov1alpha1.NotificationChannel, channelFactory iov1alpha1.ChannelFactory) error {
	k8sClient := k8s.NewClient(log, mgr.GetClient(), channelFactory)
	reconciler := newReconciler(mgr, k8sClient)

	// Create a new controller
	c, err := controller.New(controllerName, mgr, controller.Options{Reconciler: reconciler})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: channelType}, &handler.EnqueueRequestForObject{}, predicate.GenerationChangedPredicate{})
	if err != nil {
		return err
	}

	mapFn := handler.ToRequestsFunc(
		func(a handler.MapObject) []reconcile.Request {
			channels, err := k8sClient.GetChannels()
			if err != nil {
				log.Error(err, "Unable to list all channels")
			}
			requests := make([]reconcile.Request, channels.Size())
			for i, namespacedName := range channels.GetNamespacedNames() {
				requests[i] = reconcile.Request{NamespacedName: namespacedName}
			}

			return requests
		})

	err = c.Watch(&source.Kind{Type: &iov1alpha1.AlertPolicy{}}, &handler.EnqueueRequestsFromMapFunc{
		ToRequests: mapFn,
	}, k8s.StatusChangedPredicate{})
	if err != nil {
		return err
	}

	return nil
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, k8sClient *k8s.Client) reconcile.Reconciler {
	newrelicClient := internal.NewNewrelicClient(
		log,
		"https://api.newrelic.com/v2",
		os.Getenv("NEWRELIC_ADMIN_KEY"),
	)
	repository := newrelic.NewChannelRepository(log, newrelicClient)
	return &Reconcile{
		logr:     log,
		k8s:      k8sClient,
		scheme:   mgr.GetScheme(),
		newrelic: repository,
	}
}

// blank assignment to verify that Reconcile implements reconcile.Reconciler
var _ reconcile.Reconciler = &Reconcile{}

func (r *Reconcile) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling NotificationChannel")

	// Fetch the NotificationChannel instance
	instance, err := r.k8s.GetChannel(request.NamespacedName)
	if err != nil {
		if errors.IsNotFound(err) {
			return internal.NewReconcileResult(nil)
		}
		r.logr.Error(err, "Error reading object, requeueing request")
		return internal.NewReconcileResult(err)
	}

	policies, err := r.k8s.GetPolicies(instance)
	if err != nil {
		r.logr.Error(err, "Error getting policies for channel, requeueing request")
		return internal.NewReconcileResult(err)
	}

	channel := instance.NewChannel(policies)
	channel.Channel.Configuration.PreviousVersion = instance.GetStatus().NewrelicConfigVersion

	if iov1alpha1.IsDeleted(instance) {
		return r.deleteChannel(*channel, instance)
	} else {
		err = r.k8s.SetFinalizer(instance)
		if err != nil {
			reqLogger.Error(err, "Error setting finalizer on channel")
			return internal.NewReconcileResult(err)
		}

		configVersion := channel.Channel.Configuration.Version()
		instance.SetStatus(iov1alpha1.NewChannelPending(channel.Channel.Id, configVersion))
		err := r.k8s.UpdateChannelStatus(instance)
		if err != nil {
			return internal.NewReconcileResult(err)
		}

		err = r.newrelic.Save(channel)
		if err != nil {
			instance.SetStatus(iov1alpha1.NewChannelError(channel.Channel.Id, err))
			statusErr := r.k8s.UpdateChannelStatus(instance)
			if statusErr != nil {
				return internal.NewReconcileResult(statusErr)
			}

			reqLogger.Error(err, "Error saving notification channel")
			return internal.NewReconcileResult(err)
		}

		instance.SetStatus(iov1alpha1.NewChannelReady(channel.Channel.Id, configVersion))
		err = r.k8s.UpdateChannelStatus(instance)
		if err != nil {
			return internal.NewReconcileResult(err)
		}

		reqLogger.Info("Finished reconciling")
		return internal.NewReconcileResult(nil)
	}
}

func (r *Reconcile) deleteChannel(channel domain.NotificationChannel, instance iov1alpha1.NotificationChannel) (reconcile.Result, error) {
	err := r.newrelic.Delete(channel)
	if err != nil {
		r.logr.Error(err, "Error deleting policy")
		return reconcile.Result{}, err
	}

	err = r.k8s.DeleteChannel(instance)
	if err != nil {
		r.logr.Error(err, "Error updating resource")
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}
