#!/bin/sh

./disk-usage -t 1
if [ ! -f /tmp/suppress ]; then
  echo "expected lock file to exist"
  exit 1
fi

./disk-usage -t 99
if [ -f /tmp/suppress ]; then
    echo "expected lock file to not exist"
    exit 1
fi

exit 0
