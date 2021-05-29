package client

import (
	"net/http"

	"github.com/puppetlabs/relay/pkg/errors"
	"github.com/puppetlabs/relay/pkg/model"
)

func (c *Client) GetProfileSettings() (*model.GetProfileSettingsEnvelope, errors.Error) {
	resp := &model.GetProfileSettingsEnvelope{}

	if err := c.Request(
		WithPath("/auth/profile/settings"),
		WithResponseInto(&resp)); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) UpdateProfileSettings(us *model.UserSettings) errors.Error {
	req := &model.PutProfileSettingsEnvelope{
		UserSettings: us,
	}

	if err := c.Request(
		WithMethod(http.MethodPut),
		WithPath("/auth/profile/settings"),
		WithBody(req),
	); err != nil {
		return err
	}

	return nil
}
