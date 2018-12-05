#!/bin/sh

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
wget https://github.com/BlueMedoraPublic/disk-usage/releases/download/1.0.1/disk-usage.bin
chmod +x /usr/local/bin/disk-usage.bin

# setup cronjob
(crontab -l 2>/dev/null; echo "*/${INTERVAL} * * * * /usr/local/bin/disk-usage.bin -t ${THRESHOLD} -s -c ${CHANNEL}" -slack-url ${SLACK_HOOK}) | crontab -
