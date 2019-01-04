#!/bin/bash

set -e

go run event_changed.go

if  [ $? ]
then
	# Generate site
	go run generate.go > docs/index.html
	# Commit it
	git commit -am "Event changed"
	# Push it
	git push
fi
