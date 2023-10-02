package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"

	idm "github.com/tierklinik-dobersberg/cis-idm/cmds/userctl/cmds"
	roster "github.com/tierklinik-dobersberg/rosterd/cmds/rosterctl/cmds"
)

func main() {
	root := cli.New("cisctl")

	root.AddCommand(
		// IDM Commands
		idm.GetLoginCommand(root),
		idm.GetProfileCommand(root),
		idm.GetRoleCommand(root),
		idm.GetUsersCommand(root),
		idm.GetSendNotificationCommand(root),

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
