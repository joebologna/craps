#!/bin/bash

unzip -o rolls-1024.zip
find roll-? -type f -name '*.png' | while read a; do
    echo -n .
    magick "${a}" -resize 384x384 "${a}"
done
echo
