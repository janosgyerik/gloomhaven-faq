#!/usr/bin/env bash
#
# Simple wrapper to get Go dependencies and execute scaper.
# The new content is produced on stdout.
#

set -euo pipefail

# example: https://boardgamegeek.com/thread/1897763/official-faq-game-no-rules-questions-please
url=$1

cd "$(dirname "$0")"

deps=(
    "github.com/golang/protobuf/proto"
    "github.com/gocolly/colly"
)
for dep in "${deps[@]}"; do
    [[ -d "$GOPATH/src/$dep" ]] || go get "$dep"
done

go run scrape.go "$url"
