// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicy) DeepCopyInto(out *AlertPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicy.
func (in *AlertPolicy) DeepCopy() *AlertPolicy {
	if in == nil {
		return nil
	}
	out := new(AlertPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicyList) DeepCopyInto(out *AlertPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicyList.
func (in *AlertPolicyList) DeepCopy() *AlertPolicyList {
	if in == nil {
		return nil
	}
	out := new(AlertPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicySpec) DeepCopyInto(out *AlertPolicySpec) {
	*out = *in
	if in.NrqlConditions != nil {
		in, out := &in.NrqlConditions, &out.NrqlConditions
		*out = make([]NrqlCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ApmConditions != nil {
		in, out := &in.ApmConditions, &out.ApmConditions
		*out = make([]ApmCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InfraConditions != nil {
		in, out := &in.InfraConditions, &out.InfraConditions
		*out = make([]InfraCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicySpec.
func (in *AlertPolicySpec) DeepCopy() *AlertPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AlertPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertPolicyStatus) DeepCopyInto(out *AlertPolicyStatus) {
	*out = *in
	if in.NewrelicPolicyId != nil {
		in, out := &in.NewrelicPolicyId, &out.NewrelicPolicyId
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertPolicyStatus.
func (in *AlertPolicyStatus) DeepCopy() *AlertPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AlertPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApmCondition) DeepCopyInto(out *ApmCondition) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ConditionScope != nil {
		in, out := &in.ConditionScope, &out.ConditionScope
		*out = new(string)
		**out = **in
	}
	if in.Entities != nil {
		in, out := &in.Entities, &out.Entities
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	out.CriticalThreshold = in.CriticalThreshold
	if in.WarningThreshold != nil {
		in, out := &in.WarningThreshold, &out.WarningThreshold
		*out = new(Threshold)
		**out = **in
	}
	if in.UserDefined != nil {
		in, out := &in.UserDefined, &out.UserDefined
		*out = new(UserDefined)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApmCondition.
func (in *ApmCondition) DeepCopy() *ApmCondition {
	if in == nil {
		return nil
	}
	out := new(ApmCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailNotificationChannel) DeepCopyInto(out *EmailNotificationChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailNotificationChannel.
func (in *EmailNotificationChannel) DeepCopy() *EmailNotificationChannel {
	if in == nil {
		return nil
	}
	out := new(EmailNotificationChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmailNotificationChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailNotificationChannelFactory) DeepCopyInto(out *EmailNotificationChannelFactory) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailNotificationChannelFactory.
func (in *EmailNotificationChannelFactory) DeepCopy() *EmailNotificationChannelFactory {
	if in == nil {
		return nil
	}
	out := new(EmailNotificationChannelFactory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailNotificationChannelList) DeepCopyInto(out *EmailNotificationChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EmailNotificationChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailNotificationChannelList.
func (in *EmailNotificationChannelList) DeepCopy() *EmailNotificationChannelList {
	if in == nil {
		return nil
	}
	out := new(EmailNotificationChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmailNotificationChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailNotificationChannelSpec) DeepCopyInto(out *EmailNotificationChannelSpec) {
	*out = *in
	if in.PolicySelector != nil {
		in, out := &in.PolicySelector, &out.PolicySelector
		*out = make(labels.Set, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailNotificationChannelSpec.
func (in *EmailNotificationChannelSpec) DeepCopy() *EmailNotificationChannelSpec {
	if in == nil {
		return nil
	}
	out := new(EmailNotificationChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraCondition) DeepCopyInto(out *InfraCondition) {
	*out = *in
	out.CriticalThreshold = in.CriticalThreshold
	if in.WarningThreshold != nil {
		in, out := &in.WarningThreshold, &out.WarningThreshold
		*out = new(InfraThreshold)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraCondition.
func (in *InfraCondition) DeepCopy() *InfraCondition {
	if in == nil {
		return nil
	}
	out := new(InfraCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraThreshold) DeepCopyInto(out *InfraThreshold) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraThreshold.
func (in *InfraThreshold) DeepCopy() *InfraThreshold {
	if in == nil {
		return nil
	}
	out := new(InfraThreshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationChannelStatus) DeepCopyInto(out *NotificationChannelStatus) {
	*out = *in
	if in.NewrelicChannelId != nil {
		in, out := &in.NewrelicChannelId, &out.NewrelicChannelId
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationChannelStatus.
func (in *NotificationChannelStatus) DeepCopy() *NotificationChannelStatus {
	if in == nil {
		return nil
	}
	out := new(NotificationChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NrqlCondition) DeepCopyInto(out *NrqlCondition) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	out.AlertThreshold = in.AlertThreshold
	if in.WarningThreshold != nil {
		in, out := &in.WarningThreshold, &out.WarningThreshold
		*out = new(Threshold)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NrqlCondition.
func (in *NrqlCondition) DeepCopy() *NrqlCondition {
	if in == nil {
		return nil
	}
	out := new(NrqlCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackNotificationChannel) DeepCopyInto(out *SlackNotificationChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackNotificationChannel.
func (in *SlackNotificationChannel) DeepCopy() *SlackNotificationChannel {
	if in == nil {
		return nil
	}
	out := new(SlackNotificationChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlackNotificationChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackNotificationChannelFactory) DeepCopyInto(out *SlackNotificationChannelFactory) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackNotificationChannelFactory.
func (in *SlackNotificationChannelFactory) DeepCopy() *SlackNotificationChannelFactory {
	if in == nil {
		return nil
	}
	out := new(SlackNotificationChannelFactory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackNotificationChannelList) DeepCopyInto(out *SlackNotificationChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SlackNotificationChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackNotificationChannelList.
func (in *SlackNotificationChannelList) DeepCopy() *SlackNotificationChannelList {
	if in == nil {
		return nil
	}
	out := new(SlackNotificationChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlackNotificationChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackNotificationChannelSpec) DeepCopyInto(out *SlackNotificationChannelSpec) {
	*out = *in
	if in.PolicySelector != nil {
		in, out := &in.PolicySelector, &out.PolicySelector
		*out = make(labels.Set, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackNotificationChannelSpec.
func (in *SlackNotificationChannelSpec) DeepCopy() *SlackNotificationChannelSpec {
	if in == nil {
		return nil
	}
	out := new(SlackNotificationChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Threshold) DeepCopyInto(out *Threshold) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Threshold.
func (in *Threshold) DeepCopy() *Threshold {
	if in == nil {
		return nil
	}
	out := new(Threshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDefined) DeepCopyInto(out *UserDefined) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDefined.
func (in *UserDefined) DeepCopy() *UserDefined {
	if in == nil {
		return nil
	}
	out := new(UserDefined)
	in.DeepCopyInto(out)
	return out
}