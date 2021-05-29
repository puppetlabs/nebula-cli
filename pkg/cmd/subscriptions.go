package cmd

import (
	"github.com/spf13/cobra"
)

func newSubscriptionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Manage your Relay subscriptions",
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(newListUserWorkflowSubscriptions())
	cmd.AddCommand(newSubscribeUserWorkflow())
	cmd.AddCommand(newUnsubscribeUserWorkflow())

	return cmd
}

func newListUserWorkflowSubscriptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List workflow subscriptions",
		Args:  cobra.MaximumNArgs(1),
		RunE:  doListUserWorkflowSubscriptions,
	}

	return cmd
}

func newSubscribeUserWorkflow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe [workflow name]",
		Short: "Subscribe to workflow",
		Args:  cobra.MaximumNArgs(1),
		RunE:  doSubscribeUserWorkflow,
	}

	return cmd
}

func newUnsubscribeUserWorkflow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unsubscribe [workflow name]",
		Short: "Unsubscribe to workflow",
		Args:  cobra.MaximumNArgs(1),
		RunE:  doUnsubscribeUserWorkflow,
	}

	return cmd
}

func doListUserWorkflowSubscriptions(cmd *cobra.Command, args []string) error {
	Dialog.Progress("Listing workflow subscriptions...")

	uws, cwerr := Client.ListUserWorkflowSubscriptions()
	if cwerr != nil {
		return cwerr
	}

	for _, wf := range uws.Workflows {
		if wf.Subscriptions != nil {
			Dialog.Infof("%s %t %s", wf.Name, wf.Subscriptions.Subscribe, wf.Subscriptions.Custom)
		}
	}

	return nil
}

func doSubscribeUserWorkflow(cmd *cobra.Command, args []string) error {
	name, err := getWorkflowName(args)

	if err != nil {
		return err
	}

	Dialog.Progress("Subscribing...")

	uws, cwerr := Client.UpdateUserWorkflowSubscription(name, true, nil)
	if cwerr != nil {
		return cwerr
	}

	if uws.UserWorkflowSubscriptions != nil {
		Dialog.Infof("%t %s", uws.Subscribe, uws.Custom)
	}

	return nil
}

func doUnsubscribeUserWorkflow(cmd *cobra.Command, args []string) error {
	name, err := getWorkflowName(args)

	if err != nil {
		return err
	}

	Dialog.Progress("Unsubscribing...")

	uws, cwerr := Client.UpdateUserWorkflowSubscription(name, false, nil)
	if cwerr != nil {
		return cwerr
	}

	if uws.UserWorkflowSubscriptions != nil {
		Dialog.Infof("%t %s", uws.Subscribe, uws.Custom)
	}

	return nil
}
