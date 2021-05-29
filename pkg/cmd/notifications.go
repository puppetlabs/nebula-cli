package cmd

import (
	"github.com/puppetlabs/relay/pkg/format"
	"github.com/spf13/cobra"
)

func newNotificationsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "notifications",
		Short: "Manage your notifications",
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(newListUserNotifications())
	cmd.AddCommand(newClearAllReadUserNotifications())

	return cmd
}

func newListUserNotifications() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List notifications",
		Args:  cobra.MaximumNArgs(1),
		RunE:  doListUserNotifications,
	}

	return cmd
}

func newClearAllReadUserNotifications() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear all read notifications",
		Args:  cobra.MaximumNArgs(1),
		RunE:  doClearAllReadUserNotifications,
	}

	return cmd
}

func doListUserNotifications(cmd *cobra.Command, args []string) error {
	Dialog.Progress("Listing notifications...")

	uns, cwerr := Client.ListUserNotifications()
	if cwerr != nil {
		return cwerr
	}

	for _, un := range uns.Notifications {
		wfn := un.Fields["workflow_name"]
		rn := int64(un.Fields["run_number"].(float64))
		read := un.Read

		prefix := "* "
		if read {
			prefix = "  "
		}

		link := format.GuiLink(Config, "/workflows/%s/runs/%d/graph", wfn, rn)
		switch un.Type {
		case "workflow.failed":
			Dialog.Infof("%s Workflow failed: %s @ %s", prefix, wfn, link)
		case "workflow.succeeded":
			Dialog.Infof("%s Workflow succeeded: %s @ %s", prefix, wfn, link)
		case "step.approval":
			sn := un.Fields["step_name"]
			Dialog.Infof("%s Step %s needs approvals: %s @ %s", prefix, sn, wfn, link)
		}
	}

	return nil
}

func doClearAllReadUserNotifications(cmd *cobra.Command, args []string) error {
	Dialog.Progress("Clearing notifications...")

	uns, cwerr := Client.ListUserNotifications()
	if cwerr != nil {
		return cwerr
	}

	nids := make([]string, 0)
	for _, un := range uns.Notifications {
		if un.Read {
			nids = append(nids, un.ID)
		}
	}

	if len(nids) > 0 {
		cwerr := Client.MarkUserNotificationsAsDone(nids)
		if cwerr != nil {
			return cwerr
		}
	}

	return nil
}
