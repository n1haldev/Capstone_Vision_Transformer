#!/bin/bash

while IFS= read -r plant
do
  # Run the Python script
  python google_images_downloader.py "$plant"

done < plants.txt

