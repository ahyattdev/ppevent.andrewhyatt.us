#!/bin/bash

set -e

date

cd "$(dirname "$0")"

git config user.name "Andrew Hyatt"
git config user.email "4400272+ahyattdev@users.noreply.github.com"

export GO111MODULE=on

go get

if go run event_changed.go; then
	echo "Generating site"
	# Generate site
	go run generate.go > docs/index.html
	# Commit it
	git add index.html
	git commit -m "Event changed"
	# Push it
	git push
else
	echo "Nothing to do"
fi
