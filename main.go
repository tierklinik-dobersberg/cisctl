package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"

	idm "github.com/tierklinik-dobersberg/cis-idm/cmds/userctl/cmds"
	roster "github.com/tierklinik-dobersberg/rosterd/cmds/rosterctl/cmds"
)

func dumpConfig(root *cli.Root) *cobra.Command {
	return &cobra.Command{
		Use:  "dump-config",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			root.Print(root)
		},
	}
}

func main() {
	root := cli.New("cisctl")

	root.AddCommand(
		dumpConfig(root),

		// IDM Commands
		idm.GetLoginCommand(root),
		idm.GetProfileCommand(root),
		idm.GetRoleCommand(root),
		idm.GetUsersCommand(root),
		idm.GetSendNotificationCommand(root),
		idm.GetRegisterUserCommand(root),

		// Roster commands
		roster.WorkShiftCommand(root),
		roster.WorkTimeCommand(root),
		roster.ConstraintCommand(root),
		roster.RosterCommand(root),
		roster.OffTimeCommand(root),
	)

	if err := root.ExecuteContext(root.Context()); err != nil {
		logrus.Fatal(err)
	}
}
