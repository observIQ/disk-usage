package main
import (
	"os"
	"flag"

	log "github.com/golang/glog"
	"github.com/BlueMedoraPublic/disk-usage/alert/slack"
	"github.com/BlueMedoraPublic/disk-usage/lockfile"
)

const version string  = "2.0.0"

type GlobalConfig struct {
	Threshold int
	Hostname string
	Slack SlackConfig
	Dryrun bool
}

type SlackConfig struct {
	Url     string
	Channel string
}

var returnVersion bool		   // Flag returns the version and then exits
var drives        []string     // Global var stores list of drives
var globalConfig  GlobalConfig
var slackConfig   SlackConfig

func main() {
	if returnVersion {
		log.Info(version)
		os.Exit(0)
	}

	// if we cannot determine the hostname, set the hostname
	// and print the error. We still want to attempt to alert
	// TODO: Find other identifieable information such as
	// an ip address
	var err error
	globalConfig.Hostname, err = os.Hostname()
	if err != nil {
		log.Error("could not determine hostname", err)
		globalConfig.Hostname = "unknown"
	}

	if err := getMountpoints(); err != nil {
		log.Error("", err)
		os.Exit(1)
	}

	if err := getUsage(); err != nil {
		log.Error("", err)
		os.Exit(1)
	}
}


func init() {
	flag.BoolVar(&returnVersion, "version", false, "Get current version")

	flag.BoolVar(&globalConfig.Dryrun, "dryrun", false, "Run without sending alerts")
	flag.IntVar(&globalConfig.Threshold, "t", 85, "Pass a threshold as an integer")

	flag.StringVar(&globalConfig.Slack.Channel, "c", "#some_channel", "Pass a slack channel")
	flag.StringVar(&globalConfig.Slack.Url, "slack-url", "https://hooks.slack.com/services/somehook", "Pass a slack hooks URL")

	// glog flags
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")

	flag.Parse()
}


func alert(message string, newLock bool) error {
	if globalConfig.Dryrun {
		log.Info("Dry run, skipping alert")
		return nil
	}
	return slackAlert(message)
}

func slackAlert(m string) error {
	alert := slack.Alert{
		Message: m,
		Channel: globalConfig.Slack.Channel,
		URL: globalConfig.Slack.Url,
	}
	log.Info("slack alert sent: " + m)
	return alert.Send()
}

func handleLock(createLock, createAlert bool, message string) error {
	// If disk usage is healthy, and lock exists, clear it
	// by removing the lock
	if createLock == false && lockfile.Exists(lockPath()) {
		if err := alert(message + " disk usage is healthy", true); err != nil {
			return err
		}
		return lockfile.RemoveLock(lockPath())
	}

	// If disk usage is not healthy and lockfile does not exist,
	// fire off an alert
	if createLock == true && !lockfile.Exists(lockPath()) {
		if err := alert(message, false); err != nil {
			return err
		}
		return lockfile.CreateLock(lockPath())
	}

	if createLock == true && lockfile.Exists(lockPath()) {
		log.Info("Lock exists, skipping alert.")
		return nil
	}

	return nil
}
