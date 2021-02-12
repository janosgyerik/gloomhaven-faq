#!/usr/bin/env bash
#
# Simple wrapper to update (rewrite) README.md with the latest content,
# scraped from some source.
#

set -euo pipefail

cd "$(dirname "$0")"

url=https://boardgamegeek.com/thread/1897763/official-faq-game-no-rules-questions-please

./tools/bgg/scrape.sh "$url" > README.md

git diff
