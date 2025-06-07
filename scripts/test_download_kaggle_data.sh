#!/bin/bash

##chmod +x scripts/test_download_kaggle_data.sh



# list of expected files to make test more readable
expected_files=(
    "../data/links.csv"
    "../data/movies_metadata.csv"
    "../data/ratings.csv"
    "../data/keywords.csv"
    "../data/credits.csv"
    "../data/links_small.csv"
    "../data/ratings_small.csv"
)

allPassed=true

for file in ${expected_files[@]}; do
    if [ -f "$file" ]; then
        echo "PASS: File $file exists."
    else
        echo "FAIL: File $file does not exist."
        allPassed=false
    fi
done

if $allPassed; then
    echo "All expected files are present."
    exit 0
else
    echo "Some expected files are missing."
    exit 1
fi