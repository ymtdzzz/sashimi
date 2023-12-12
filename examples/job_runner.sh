#!/bin/bash

TRIGGER_CMD=$1

if [ -z "$TRIGGER_CMD" ]; then
    echo "Usage: $0 \"[job split command]\""
    exit 1
fi

echo "Start job triggered by scheduler such as cron or AWS CloudWatch Events, etc..."

CMDS_STR=$(eval $TRIGGER_CMD)

IFS=',' read -ra ADDR <<< "$CMDS_STR"
for i in "${ADDR[@]}"; do
    echo "Start splitted job on a container!"
    echo "Running command: $i"
    eval ./$i
    echo "Done!"
done
