#!/bin/bash

set -e

echo "Looking at previous hash"

go run event_changed.go

status=$?
echo "Exit status is $status"
if echo hi; then
	echo "Generating site site"
	# Generate site
	go run generate.go > docs/index.html
	# Commit it
	git commit -am "Event changed"
	# Push it
	git push
else
	echo "Nothing to do"
fi
