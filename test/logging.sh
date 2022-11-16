#!/bin/bash
./../openmedia-check -i "$1" | tee log.json | jq -c 'select(.status=="FAILURE") | .data.file + " " + .data.week' | tr -d "\"" | tee moves.txt
# | xargs mv
