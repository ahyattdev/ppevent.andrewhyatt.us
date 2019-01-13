#!/bin/bash

set -e

date

cd ~/git/ppevent.andrewhyatt.us

if /usr/local/bin/go run event_changed.go; then
	echo "Generating site"
	# Generate site
	/usr/local/bin/go run generate.go > docs/index.html
	# Commit it
	git commit -am "Event changed"
	# Push it
	git push
else
	echo "Nothing to do"
fi
