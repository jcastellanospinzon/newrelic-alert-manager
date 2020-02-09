package newrelic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fpetkovski/newrelic-operator/internal"
	"github.com/fpetkovski/newrelic-operator/pkg/alert_policies/domain"
	"github.com/go-logr/logr"
	"github.com/opentracing/opentracing-go"
	"io/ioutil"
)

type nrqlConditionRepository struct {
	client internal.NewrelicClient
	log    logr.Logger
}

func newNrqlConditionRepository(log logr.Logger, client internal.NewrelicClient) *nrqlConditionRepository {
	return &nrqlConditionRepository{
		client: client,
		log:    log,
	}
}

func (repository nrqlConditionRepository) getConditions(policyId int64) (*domain.NrqlConditionList, error) {
	repository.log.Info("Getting conditions for policy", "PolicyId", policyId)

	endpoint := fmt.Sprintf("alerts_nrql_conditions.json?policy_id=%d", policyId)
	response, err := repository.client.Get(endpoint)
	if err != nil {
		return nil, err
	}

	var conditionList domain.NrqlConditionList
	err = json.NewDecoder(response.Body).Decode(&conditionList)
	if err != nil {
		return nil, err
	}

	return &conditionList, nil
}

func (repository nrqlConditionRepository) saveConditions(ctx context.Context, policy *domain.AlertPolicy) error {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "save-nrql-alert-conditions-newrelic")
	defer sp.Finish()

	existingConditions, err := repository.getConditions(*policy.Policy.Id)
	if err != nil {
		return err
	}

	newConditionsSet := domain.NewNrqlConditionSetFromSlice(policy.NrqlConditions)
	for _, condition := range existingConditions.Condition {
		if newConditionsSet.Contains(condition) {
			continue
		}

		err := repository.deleteConditions(*condition.Id)
		if err != nil {
			return err
		}
	}

	existingConditionSet := domain.NewNrqlConditionSet(*existingConditions)
	for _, newCondition := range policy.NrqlConditions {
		if existingConditionSet.Contains(newCondition.Condition) {
			continue
		}
		err := repository.saveCondition(*policy.Policy.Id, newCondition)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repository nrqlConditionRepository) deleteConditions(conditionId int64) error {
	repository.log.Info("Deleting condition", "ConditionId", conditionId)

	endpoint := fmt.Sprintf("alerts_nrql_conditions/%d.json", conditionId)
	_, err := repository.client.Delete(endpoint)

	return err
}

func (repository nrqlConditionRepository) saveCondition(policyId int64, condition *domain.NrqlCondition) error {
	repository.log.Info("Saving condition", "Policy Id", policyId, "NrqlConditionBody", condition)
	payload, err := json.Marshal(&condition)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("alerts_nrql_conditions/policies/%d.json", policyId)
	response, err := repository.client.PostJson(endpoint, payload)
	if response != nil && response.StatusCode >= 300 {
		responseContent, _ := ioutil.ReadAll(response.Body)
		return errors.New(string(responseContent))
	}

	if err != nil {
		return err
	}

	return nil
}
