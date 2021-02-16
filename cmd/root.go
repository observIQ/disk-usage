package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/BlueMedoraPublic/disk-usage/internal/alert"
	"github.com/BlueMedoraPublic/disk-usage/internal/disk"
	"github.com/BlueMedoraPublic/disk-usage/internal/lock"
	"github.com/BlueMedoraPublic/disk-usage/internal/pkg/host"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const version string = "3.0.1"

// flags
var (
	v         bool // print version and exit
	dryrun    bool
	threshold int
	hostname  string
	alertType string
	logLevel  string

	// slack
	slackChannel string
	slackHookURL string
)

// Execute is the main function, will run disk-usage
func Execute() {
	if v {
		fmt.Println("disk-usage version:", version)
		os.Exit(0)
	}

	if err := execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func execute() error {
	c, err := initConfig()
	if err != nil {
		return err
	}
	return c.Run()
}

func init() {
	flag.BoolVar(&v, "version", false, "Print version")
	flag.BoolVar(&dryrun, "dryrun", false, "Run without sending alerts")
	flag.IntVar(&threshold, "t", 85, "Disk usage percentage that should trigger an alert")
	flag.StringVar(&alertType, "alert-type", "slack", "Alert type to use. Defaults to slack for backwards compatability, falls back on Stdout if slack params are not set")
	flag.StringVar(&hostname, "hostname", "", "Set the hostname")
	flag.StringVar(&logLevel, "log-level", "info", "Set log level (error, warning, info, trace)")

	// slack
	flag.StringVar(&slackChannel, "c", "", "Slack channel")
	flag.StringVar(&slackHookURL, "slack-url", "", "Slack webhook url")
	flag.Parse()
}

func initConfig() (disk.Config, error) {
	initLog()

	if hostname == "" {
		h, err := os.Hostname()
		if err != nil {
			log.Error("could not determine hostname", err)
			hostname = "unknown"
		}
		hostname = h
	}

	ip, err := host.PrimaryAddress()
	if err != nil {
		log.Error("could not determine ip address", err)
	}

	if err := validateFlags(); err != nil {
		return disk.Config{}, err
	}

	a, err := initAlert()
	if err != nil {
		return disk.Config{}, err
	}

	l, err := initLock()
	if err != nil {
		return disk.Config{}, err
	}

	return disk.Config{
		Threshold: threshold,
		Alert:     a,
		Lock:      l,
		Host: disk.System{
			Name:    hostname,
			Address: ip,
		},
	}, nil
}

func initLog() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Error(errors.Wrap(err, "Invalid log level set, using INFO"))
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

// initAlert sets the alert type. Default to slack if slackHookURL is set, for
// backwards compatability. Fall back on stdout if slack is not set.
func initAlert() (alert.Alert, error) {
	if alertType == "stdout" || dryrun {
		return alert.Stdout()
	}

	if alertType == "slack" && slackHookURL != "" {
		return alert.Slack(slackChannel, slackHookURL)
	}

	if alertType == "slack" && slackHookURL == "" {
		return alert.Stdout()
	}

	return nil, fmt.Errorf(fmt.Sprintf("failed to set alert type %s", alertType))
}

func initLock() (lock.Lock, error) {
	if dryrun {
		return lock.Null(), nil
	}

	// const defined in root_unix.go / root_windows.go
	return lock.File(lockpath)
}

func validateFlags() error {
	return nil
}
