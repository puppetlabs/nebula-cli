package model

type NotificationIdentifier struct {
	ID string `json:"id"`
}

type NotificationSummary struct {
	*NotificationIdentifier
	*Lifecycle

	Type       string   `json:"type"`
	Attributes []string `json:"attributes,omitempty"`
}

type GetNotificationEnvelope struct {
	*NotificationSummary

	State  string                 `json:"state"`
	Done   bool                   `json:"done"`
	Read   bool                   `json:"read"`
	Fields map[string]interface{} `json:"fields"`
}

type GetNotificationsEnvelope struct {
	Notifications []*GetNotificationEnvelope `json:"notifications"`
}

type PostAllNotificationDoneEnvelope struct {
	IDs *[]string `json:"ids,omitempty"`
}
