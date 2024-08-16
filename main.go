package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"

	pbx "github.com/tierklinik-dobersberg/3cx-support/cmd/callctl/cmds"
	calendar "github.com/tierklinik-dobersberg/cis-cal/cmds/ciscalctl/cmds"
	idm "github.com/tierklinik-dobersberg/cis-idm/cmds/idmctl/cmds"
	comments "github.com/tierklinik-dobersberg/comment-service/cmds/client/cmds"
	customer "github.com/tierklinik-dobersberg/customer-service/cmds/customercli/cmds"
	roster "github.com/tierklinik-dobersberg/rosterd/cmds/rosterctl/cmds"
)

func dumpConfig(root *cli.Root) *cobra.Command {
	return &cobra.Command{
		Use:  "dump-config",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			root.Print(root.Config())
		},
	}
}

func main() {
	root := cli.New("cisctl")

	customerCmd := &cobra.Command{
		Use:     "customers",
		Aliases: []string{"customer"},
	}

	customerCmd.AddCommand(
		customer.GetSearchCommand(root),
		customer.GetUpdateCustomerCommand(root),
	)

	root.AddCommand(
		dumpConfig(root),

		// IDM Commands
		idm.GetLoginCommand(root),
		idm.GetProfileCommand(root),
		idm.GetRoleCommand(root),
		idm.GetUsersCommand(root),
		idm.GetSendNotificationCommand(root),
		idm.GetRegisterUserCommand(root),
		idm.GenerateVAPIDKeys(),

		// Roster commands
		roster.WorkShiftCommand(root),
		roster.WorkTimeCommand(root),
		roster.ConstraintCommand(root),
		roster.RosterCommand(root),
		roster.OffTimeCommand(root),

		// Comments command
		comments.CommentsCommand(root),
		comments.ScopeCommand(root),

		// Calendar commands
		calendar.GetCalendarCommand(root),
		calendar.GetEventsCommand(root),
		calendar.GetHolidayCommand(root),

		// Call commands
		pbx.GetCallLogCommand(root),
		pbx.GetOnDutyCommand(root),
		pbx.GetInboundNumbersCommand(root),

		// Customer Commands
		customerCmd,
	)

	if err := root.ExecuteContext(root.Context()); err != nil {
		logrus.Fatal(err)
	}
}
