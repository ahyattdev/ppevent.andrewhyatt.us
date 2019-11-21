#!/bin/bash

set -e

date

cd "$(dirname "$0")"

git config user.name "Andrew Hyatt"
git config user.email "4400272+ahyattdev@users.noreply.github.com"

export PATH="/opt/go/bin/:$PATH"

export GO111MODULE=on

go get

go run generate.go

git add docs/index.html
git commit -m "Event changed"
git push
