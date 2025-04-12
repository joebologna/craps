#!/bin/bash

# Input file (your high-resolution image)
input_image="craps-icon.png"

# Output directory
output_dir="AppIcons"
mkdir -p "$output_dir"

# Resize and save images
magick "$input_image" -resize 40x40 "$output_dir/icon-40x40@2x.png"
magick "$input_image" -resize 58x58 "$output_dir/icon-29x29@2x.png"
magick "$input_image" -resize 87x87 "$output_dir/icon-29x29@3x.png"
magick "$input_image" -resize 80x80 "$output_dir/icon-40x40@2x.png"
magick "$input_image" -resize 120x120 "$output_dir/icon-60x60@2x.png"
magick "$input_image" -resize 152x152 "$output_dir/icon-76x76@2x.png"
magick "$input_image" -resize 167x167 "$output_dir/icon-83.5x83.5@2x.png"
magick "$input_image" -resize 1024x1024 "$output_dir/icon-1024x1024.png"

echo "Icons resized and saved in $output_dir"
