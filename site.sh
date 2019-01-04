#!/bin/bash

set -e

if go run event_changed.go; then
	echo "Generating site"
	# Generate site
	go run generate.go > docs/index.html
	# Commit it
	git commit -am "Event changed"
	# Push it
	git push
else
	echo "Nothing to do"
fi
