package main
import (
	"os"
	"flag"
	"fmt"
	"bytes"
	"net/http"

	"gopkg.in/gomail.v2"
)


const version string  = "1.0.1"


type GlobalConfig struct {
	Threshold int
	UseSlack  bool
}

type EmailConfig struct {
	Recipient string
	Sender    string
	Server    string
	Port	  int
}

type SlackConfig struct {
	Url     string
	Channel string
}


var returnVersion bool		   // Flag returns the version and then exits
var dryrun 		  bool         // When true, no alerts are sent
var drives        []string     // Global var stores list of drives
var globalConfig  GlobalConfig
var emailConfig   EmailConfig
var slackConfig   SlackConfig


func main() {
	if returnVersion {
		getVersion()
		os.Exit(0)
	}

	getMountpoints()
	getUsage()
}


func init() {
	flag.BoolVar(&returnVersion, "version", false, "Get current version")

	flag.BoolVar(&dryrun, "dryrun", false, "Run without sending alerts")
	flag.IntVar(&globalConfig.Threshold, "t", 85, "Pass a threshold as an integer")

	flag.BoolVar(&globalConfig.UseSlack, "s", false, "Enable slack by passing 'true'")
	flag.StringVar(&slackConfig.Channel, "c", "#some_channel", "Pass a slack channel")
	flag.StringVar(&slackConfig.Url, "slack-url", "https://hooks.slack.com/services/somehook", "Pass a slack hooks URL")

	flag.StringVar(&emailConfig.Recipient, "r", "email@domain.com", "Pass an email recipient")
	flag.StringVar(&emailConfig.Server, "smtp-server", "smtp.domain.localnet", "Pass an smtp server hostname")
	flag.IntVar(&emailConfig.Port, "smtp-port", 25, "Pass an smtp listening port")

	flag.Parse()

	emailConfig.Sender = "admin@" + getHostname()
}


func alert(message string, lock bool) bool {
	var result bool

	if dryrun {
		fmt.Println("Dry run, skipping alert")
		return true
	}

	if lock == true && lockExists(lockpath) {
		fmt.Println("Lock exists, skipping alert.")
		return true
	}

	if globalConfig.UseSlack != true {
		result = alertEmail(message)
	} else {
		result = alertSlack(message)
	}

	// If alert sent, create lock file only if lock == true
	if result == true  {
		fmt.Println("Alert sent")
		if lock == true {
			createLock(lockpath)
		}
		return true

	} else {
		fmt.Println("Failed to send alert!")
		return false
	}
}


func alertSlack(message string) bool {
	var json []byte = []byte(
		`
		{
			"channel": "` + slackConfig.Channel + `",
			"text":"` + message + `"
		}
		`)

	req, err := http.NewRequest("POST", slackConfig.Url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// If error, send email instead
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(err)
		fmt.Println("Problem with slack, falling back to email.")

		return alertEmail(message)

	} else if resp.StatusCode == 200 {
		return true

	} else {
		return false
	}
}


func alertEmail(message string) bool {
	s := getHostname() + " disk usage"
	d := gomail.NewDialer(emailConfig.Server, emailConfig.Port, "", "")
	m := gomail.NewMessage()

	m.SetHeader("From", emailConfig.Sender)
	m.SetHeader("To", emailConfig.Recipient)
	m.SetHeader("Subject", s)
	m.SetBody("text/html", message)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}


func getHostname() string {
  hostname, err := os.Hostname()
  if err != nil {
	  fmt.Println("Failed to get hostname. . . Returning something..")
	  return "Could luck, could not determine my hostname."
  }
  return hostname
}


func getVersion() {
	fmt.Printf(version)
}
