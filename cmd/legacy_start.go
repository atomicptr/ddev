package cmd

import (
	"fmt"
	"log"

	"github.com/drud/bootstrap/cli/local"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// LegacyStartCmd represents the stop command
var LegacyStartCmd = &cobra.Command{
	Use:   "start [app_name] [environment_name]",
	Short: "Start an application's local services.",
	Long:  `Start will turn on the local containers that were previously stopped for an app.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := local.LegacyApp{
			Name:        activeApp,
			Environment: activeDeploy,
			Template:    local.LegacyComposeTemplate,
		}

		err := app.Start()
		if err != nil {
			log.Println(err)
			Failed("Failed to start site.")
		}

		err = app.Config()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Waiting for site readiness. This may take a couple minutes...")
		siteURL, err := app.Wait()
		if err != nil {
			log.Println(err)
			Failed("Site failed to achieve readiness.")
		}

		color.Cyan("Successfully started %s %s", activeApp, activeDeploy)
		if siteURL != "" {
			color.Cyan("Your application can be reached at: %s", siteURL)
		}

	},
}

func init() {

	LegacyCmd.AddCommand(LegacyStartCmd)

}