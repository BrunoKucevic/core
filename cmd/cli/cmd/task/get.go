package task

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/theopenlane/core/cmd/cli/cmd"
	"github.com/theopenlane/core/pkg/openlaneclient"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing task",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	command.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "task id to query")
	getCmd.Flags().BoolP("include-completed", "c", false, "include completed tasks")
}

// get an existing task in the platform
func get(ctx context.Context) error {
	// attempt to setup with token, otherwise fall back to JWT with session
	client, err := cmd.TokenAuth(ctx, cmd.Config)
	if err != nil || client == nil {
		// setup http client
		client, err = cmd.SetupClientWithAuth(ctx)
		cobra.CheckErr(err)
		defer cmd.StoreSessionCookies(client)
	}
	// filter options
	id := cmd.Config.String("id")

	// if an task ID is provided, filter on that task, otherwise get all
	if id != "" {
		o, err := client.GetTaskByID(ctx, id)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	includeCompleted := cmd.Config.Bool("include-completed")
	if includeCompleted {
		// get all will be filtered for the authorized organization(s)
		o, err := client.GetAllTasks(ctx)
		cobra.CheckErr(err)

		return consoleOutput(o)

	}

	o, err := client.GetTasks(ctx, nil, nil, &openlaneclient.TaskWhereInput{
		CompletedIsNil: &includeCompleted,
	})

	return consoleOutput(o)
}
