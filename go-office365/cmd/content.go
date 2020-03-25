package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/devodev/go-graph/office365"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCommandContent())
}

func newCommandContent() *cobra.Command {
	var (
		pubIdentifier string
		startTime     string
		endTime       string
	)

	cmd := &cobra.Command{
		Use:   "content [content-type]",
		Short: "List content that is available to be fetched for the provided content-type.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// command line args
			ctArg := args[0]

			// validate args
			if !office365.ContentTypeValid(ctArg) {
				fmt.Println("ContentType invalid")
				return
			}
			ct, err := office365.GetContentType(ctArg)
			if err != nil {
				fmt.Println(err)
				return
			}

			// parse optional args
			if pubIdentifier == "" {
				pubIdentifier = config.Credentials.ClientID
			}
			startTime := parseDate(startTime)
			endTime := parseDate(endTime)

			client := office365.NewClientAuthenticated(&config.Credentials)
			content, err := client.Subscriptions.Content(context.Background(), pubIdentifier, ct, startTime, endTime)
			if err != nil {
				fmt.Printf("error getting content: %s\n", err)
				return
			}
			for _, u := range content {
				userData, err := json.Marshal(u)
				if err != nil {
					fmt.Printf("error marshalling content: %s\n", err)
					continue
				}
				fmt.Println(string(userData))
			}
		},
	}
	cmd.Flags().StringVar(&pubIdentifier, "identifier", "", "Publisher Identifier")
	cmd.Flags().StringVar(&startTime, "start", "", "Start time")
	cmd.Flags().StringVar(&endTime, "end", "", "End time")

	return cmd
}

func parseDate(param string) time.Time {
	formats := []string{
		office365.RequestDateFormat,
		office365.RequestDatetimeFormat,
		office365.RequestDatetimeLargeFormat,
	}
	for _, format := range formats {
		parsed, err := time.Parse(format, param)
		if err == nil {
			return parsed
		}
	}
	return time.Time{}
}