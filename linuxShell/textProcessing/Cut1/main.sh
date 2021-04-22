#!/bin/bash

while IFS= read -r LINE || [[ -n "$LINE" ]]; do
    printf "%s\n" `echo "$LINE" | head -c 3 | tail -c 1`
done