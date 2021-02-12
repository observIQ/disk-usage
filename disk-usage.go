package main
import (
	"os"
	"flag"

	log "github.com/golang/glog"
	"github.com/BlueMedoraPublic/disk-usage/alert"
	"github.com/BlueMedoraPublic/disk-usage/lockfile"
)

const version string  = "3.0.0"

type GlobalConfig struct {
	Threshold int
	Hostname string
	Dryrun bool

	// alert interface
	alert alert.Alert
}

var returnVersion bool		   // Flag returns the version and then exits
var drives        []string     // Global var stores list of drives
var globalConfig  GlobalConfig

// slack
var (
	slackHookURL string
	slackChannel string
)

func main() {
	if returnVersion {
		log.Info(version)
		os.Exit(0)
	}

	if err := initAlert(); err != nil {
		log.Error(err)
		os.Exit(1)
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

	flag.StringVar(&slackChannel, "c", "#some_channel", "Pass a slack channel")
	flag.StringVar(&slackHookURL, "slack-url", "https://hooks.slack.com/services/somehook", "Pass a slack hooks URL")

	// glog flags
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")

	flag.Parse()
}


func sendAlert(message string, newLock bool) error {
	if globalConfig.Dryrun {
		log.Info("Dry run, skipping alert")
		return nil
	}
	if err := globalConfig.alert.Send(message); err != nil {
		return err
	}
	log.Info("Alert sent: " + message)
	return nil
}

func handleLock(createLock, createAlert bool, message string) error {
	// If disk usage is healthy, and lock exists, clear it
	// by removing the lock
	if createLock == false && lockfile.Exists(lockPath()) {
		if err := sendAlert(message + " disk usage is healthy", true); err != nil {
			return err
		}
		return lockfile.RemoveLock(lockPath())
	}

	// If disk usage is not healthy and lockfile does not exist,
	// fire off an alert
	if createLock == true && !lockfile.Exists(lockPath()) {
		if err := sendAlert(message, false); err != nil {
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
