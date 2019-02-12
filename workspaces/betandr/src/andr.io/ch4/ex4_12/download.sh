#!/bin/bash
for i in {1..2083}
do
   curl -H "Accept: application/json" "https://xkcd.com/$i/info.0.json" > ./index/$i.json
   sleep 0.5
done
