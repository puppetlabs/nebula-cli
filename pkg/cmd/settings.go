package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

func newSettingsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Manage your settings",
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(newViewProfileSettings())
	cmd.AddCommand(newUpdateUserSubscriptions())

	return cmd
}

func newViewProfileSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "View current profile settings",
		Args:  cobra.MaximumNArgs(1),
		RunE:  doViewProfileSettings,
	}

	return cmd
}

func newUpdateUserSubscriptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [subscription] [channel] [enabled]",
		Short: "Update user subscription settings",
		Args:  cobra.MaximumNArgs(3),
		RunE:  doUpdateUserSubscriptions,
	}

	return cmd
}

func doViewProfileSettings(cmd *cobra.Command, args []string) error {
	Dialog.Progress("Getting profile settings...")

	ups, cwerr := Client.GetProfileSettings()
	if cwerr != nil {
		return cwerr
	}

	Dialog.Infof("Digest enabled: %t", ups.Digest.Enabled)
	Dialog.Infof("Subscriptions enabled:")
	Dialog.Infof("* Approvals:")
	Dialog.Infof("  * Application: %t", ups.Subscriptions.Workflows.Approval.Application)
	Dialog.Infof("  * Email: %t", ups.Subscriptions.Workflows.Approval.Email)
	Dialog.Infof("* Failures:")
	Dialog.Infof("  * Application: %t", ups.Subscriptions.Workflows.Failure.Application)
	Dialog.Infof("  * Email: %t", ups.Subscriptions.Workflows.Failure.Email)
	Dialog.Infof("* Success:")
	Dialog.Infof("  * Application: %t", ups.Subscriptions.Workflows.Success.Application)
	Dialog.Infof("  * Email: %t", ups.Subscriptions.Workflows.Success.Email)

	return nil
}

func doUpdateUserSubscriptions(cmd *cobra.Command, args []string) error {
	Dialog.Progress("Updating user subscription settings...")

	// FIXME Add error messaging here
	if len(args) < 2 {
		return nil
	}

	ups, cwerr := Client.GetProfileSettings()
	if cwerr != nil {
		return cwerr
	}

	st := args[0]
	ct := args[1]
	e, err := strconv.ParseBool(args[2])
	if err != nil {
		return err
	}

	us := ups.UserSettings

	switch st {
	case "approvals":
		switch ct {
		case "application":
			us.Subscriptions.Workflows.Approval.Application = e
		case "email":
			us.Subscriptions.Workflows.Approval.Email = e
		}
	case "failures":
		switch ct {
		case "application":
			us.Subscriptions.Workflows.Failure.Application = e
		case "email":
			us.Subscriptions.Workflows.Failure.Email = e
		}
	case "success":
		switch ct {
		case "application":
			us.Subscriptions.Workflows.Success.Application = e
		case "email":
			us.Subscriptions.Workflows.Success.Email = e
		}

	}

	cwerr = Client.UpdateProfileSettings(us)
	if cwerr != nil {
		return cwerr
	}

	return nil
}
