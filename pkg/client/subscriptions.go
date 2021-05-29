package client

import (
	"fmt"
	"net/http"

	"github.com/puppetlabs/relay/pkg/errors"
	"github.com/puppetlabs/relay/pkg/model"
)

func (c *Client) ListUserWorkflowSubscriptions() (*model.ListUserWorkflowSubscriptionsEnvelope, errors.Error) {
	resp := &model.ListUserWorkflowSubscriptionsEnvelope{}

	if err := c.Request(
		WithPath("/api/subscriptions/workflows"),
		WithResponseInto(&resp)); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) UpdateUserWorkflowSubscription(name string, subscribe bool, custom []string) (*model.GetUserWorkflowSubscriptionsEnvelope, errors.Error) {
	params := &model.PutUserWorkflowSubscriptionsEnvelope{
		UserWorkflowSubscriptions: &model.UserWorkflowSubscriptions{
			Subscribe: subscribe,
			Custom:    custom,
		},
	}

	response := &model.GetUserWorkflowSubscriptionsEnvelope{}

	if err := c.Request(
		WithMethod(http.MethodPut),
		WithPath(fmt.Sprintf("/api/subscriptions/workflows/%s", name)),
		WithBody(params),
		WithResponseInto(response),
	); err != nil {
		return nil, err
	}

	return response, nil
}
