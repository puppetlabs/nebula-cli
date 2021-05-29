package model

type ListUserWorkflowSubscriptions struct {
	Workflows []*UserWorkflowSubscriptionsSummary `json:"workflows"`
}

type UserWorkflowSubscriptionsSummary struct {
	*WorkflowIdentifier
	Subscriptions *UserWorkflowSubscriptions `json:"subscriptions"`
}

type UserWorkflowSubscriptions struct {
	Subscribe bool     `json:"subscribe"`
	Custom    []string `json:"custom"`
}

type ListUserWorkflowSubscriptionsEnvelope struct {
	*ListUserWorkflowSubscriptions
}

type GetUserWorkflowSubscriptionsEnvelope struct {
	*UserWorkflowSubscriptions
}

type PutUserWorkflowSubscriptionsEnvelope struct {
	*UserWorkflowSubscriptions
}
