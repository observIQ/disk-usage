#!/bin/sh

# keep this script in the root level of the repo, it may be retrieved
# programmatically by deployment systems

# web hook url is the first argument
SLACK_HOOK=${1}

# slack channel should be second argument
CHANNEL=${2}

# disk consumption percentage that will trigger the alert
THRESHOLD=80

# how often to check disk usage, in minutes
INTERVAL=10

# deploy
cd /usr/local/bin
wget https://github.com/BlueMedoraPublic/disk-usage/releases/download/3.1.0/disk-usage-linux-amd64
chmod +x /usr/local/bin/disk-usage-linux-amd64

# setup cronjob
(crontab -l 2>/dev/null; echo "*/${INTERVAL} * * * * /usr/local/bin/disk-usage-linux-amd64 -t ${THRESHOLD} -c ${CHANNEL}" -slack-url ${SLACK_HOOK}) | crontab -
