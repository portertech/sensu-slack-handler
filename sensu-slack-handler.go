package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/bluele/slack"
	"github.com/sensu/sensu-go/types"
)

func main() {
	hookURL := flag.String("hook-url", "", "The Slack Webhook URL to use when sending notifications.")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	stdin, _ := reader.ReadString('\n')

	rawEvent := []byte(stdin)
	event := &types.Event{}

	if err := json.Unmarshal(rawEvent, event); err != nil {
		fmt.Errorf("%v", err)
		os.Exit(2)
	}

	hook := slack.NewWebHook(*hookURL)

	err := hook.PostMessage(&slack.WebHookPostPayload{
		Text: event.Check.Output,
		Attachments: []*slack.Attachment{
			{Text: "danger", Color: "danger"},
		},
	})

	if err != nil {
		fmt.Errorf("%v", err)
		os.Exit(2)
	}
}
