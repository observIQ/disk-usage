package main

import (
    "github.com/BlueMedoraPublic/disk-usage/alert/slack"
)

func initAlert() error {
    return initAlertSlack()
}

func initAlertSlack() error {
    globalConfig.alert = slack.Slack{
        Channel: slackChannel,
        HookURL: slackHookURL,
    }
    return globalConfig.alert.Init()
}
