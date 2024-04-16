#!/bin/sh

dir=$(basename "$(pwd)")

if [ "$dir" != 'leetcode' ]; then
    echo "Your current directory is $dir"
    echo "This command is meant to be run in the leetcode directory. Please run this command in the correct directory."
    exit 1
fi

if [ "$#" -ne 1 ]; then
echo "Invalid arguments received. Please try again with correct arguments."
elif [ -d "$1" ]; then
  echo "Error: directory already exists."
else
  mkdir "$1" \
  && touch "$1"/solution.go \
  && echo "package _$1" >> "$1"/solution.go
fi

echo "Directory $1 successfully created"