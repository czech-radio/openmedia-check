#!/bin/bash
./../openmedia-check -i "$1" | tee log.json \
#  | jq -c '.data | select(.status=="FAILURE") | "mv " + .file + " ../" + .week' \
  | tee moves.txt
