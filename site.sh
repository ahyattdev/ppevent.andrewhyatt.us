#!/bin/bash

set -e

if [ event_changed.go ] ; then
	# Generate site
	go run generate.go > docs/index.html
	# Commit it
	git commit -am "Event changed"
	# Push it
	git push
fi
