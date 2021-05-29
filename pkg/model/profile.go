package model

type UserDigestSettings struct {
	Enabled bool `json:"enabled"`
}

type UserSubscriptionChannels struct {
	Application bool `json:"app"`
	Email       bool `json:"email"`
}

type UserWorkflowSubscriptionSettings struct {
	Approval *UserSubscriptionChannels `json:"approval"`
	Failure  *UserSubscriptionChannels `json:"failure"`
	Success  *UserSubscriptionChannels `json:"success"`
}

type UserSubscriptionSettings struct {
	Workflows *UserWorkflowSubscriptionSettings `json:"workflows"`
}

type UserSettings struct {
	Digest        *UserDigestSettings       `json:"digest"`
	Subscriptions *UserSubscriptionSettings `json:"subscriptions"`
}

type GetProfileSettingsEnvelope struct {
	*UserSettings
}

type PutProfileSettingsEnvelope struct {
	*UserSettings
}
