#!/bin/bash
for i in {1..100}; do
    printf -v url "http://localhost:8000/?xmax=-1.444&xmin=-2.444&ymax=3&ymin=2&zoom=$i"
    printf -v file "./pngs/%010d.png" $i
    curl $url > $file
done
