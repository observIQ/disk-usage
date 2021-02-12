package cmd

import (
    "os"
    "fmt"
    "flag"

    "github.com/BlueMedoraPublic/disk-usage/internal/alert"
    "github.com/BlueMedoraPublic/disk-usage/internal/disk"
    "github.com/BlueMedoraPublic/disk-usage/internal/lock"

    log "github.com/golang/glog"
)

const version string  = "3.0.0"

// flags
var (
    v bool // print version and exit
    dryrun bool
    threshold int
    hostname string
    alertType string

    // slack
    slackChannel  string
    slackHookURL  string
)

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

    // slack
    flag.StringVar(&slackChannel, "c", "", "Slack channel")
    flag.StringVar(&slackHookURL, "slack-url", "", "Slack webhook urlL")

    // glog flags
    flag.Set("logtostderr", "true")
    flag.Set("stderrthreshold", "WARNING")

    flag.Parse()
}

func initConfig() (disk.Config, error) {
    if err := initHostname(); err != nil {
        return disk.Config{}, err
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
        Hostname: hostname,
        Alert: a,
        Lock: l,
    }, nil
}

func initHostname() error {
    if hostname == "" {
        h, err := os.Hostname()
        if err != nil {
            log.Error("could not determine hostname", err)
            hostname = "unknown"
        }
        hostname = h
    }
    return nil
}

// initAlert sets the alert type. Default to slack if slackHookURL is set, for
// backwards compatability. Fall back on stdout if slack is not set.
func initAlert() (alert.Alert, error) {
    if dryrun {
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
    // const defined in root_unix.go / root_windows.go
    return lock.File(lockpath)
}

func validateFlags() error {
    return nil
}
