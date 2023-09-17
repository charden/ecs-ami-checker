package slack

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"os"
)

func NotifySlack(amiId string) {
	webhookUrl := os.Getenv("WEBHOOK_URL")

	payload := slack.Payload{
		Text:      os.Getenv("AMI_PARAMETER_PATH") + "に新しいAMIが追加されました\n最新のAMI ID: " + amiId,
		Username:  os.Getenv("SLACK_USERNAME"),
		Channel:   os.Getenv("SLACK_CHANNEL"),
		IconEmoji: os.Getenv("SLACK_ICON_EMOJI"),
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
