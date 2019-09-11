#!/usr/bin/env bash

set -euo pipefail

mkdir -pv xkcd

for x in {1..2201}
do
    curl --output "xkcd/$x.json" "https://xkcd.com/$x/info.0.json"
done
