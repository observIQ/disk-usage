#!/bin/sh

./disk-usage -t 1
OUT=$(jq .alerted /tmp/suppress | jq length)
if [ "$OUT" -eq 0 ]; then
  echo "expected state to not be empty"
  exit 1
fi

./disk-usage -t 99
OUT=$(jq .alerted /tmp/suppress | jq length)
if [ "$OUT" -eq 0 ]; then
    echo "expected state to be empty"
  exit 1
fi

exit 0
