package client

import (
	"net/http"

	"github.com/puppetlabs/relay/pkg/errors"
	"github.com/puppetlabs/relay/pkg/model"
)

func (c *Client) ListUserNotifications() (*model.GetNotificationsEnvelope, errors.Error) {
	resp := &model.GetNotificationsEnvelope{}

	if err := c.Request(
		WithPath("/api/notifications"),
		WithResponseInto(&resp)); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) MarkUserNotificationsAsDone(notificationIDs []string) errors.Error {
	req := &model.PostAllNotificationDoneEnvelope{
		IDs: &notificationIDs,
	}

	if err := c.Request(
		WithMethod(http.MethodPost),
		WithPath("/api/notifications/done"),
		WithBody(req),
	); err != nil {
		return err
	}

	return nil
}
