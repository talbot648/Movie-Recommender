#!/bin/bash

# Creating data folder
mkdir -p ../data

# Download and unzip the dataset and put into data folder
kaggle datasets download rounakbanik/the-movies-dataset --unzip -p ../data

# displays once download is done
echo "Dataset downloaded and unzipped into data folder."

## makes the script executable
# chmod +x scripts/download_kaggle_data.sh

