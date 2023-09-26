#!/bin/bash

for dir in */; do
  dir=$(basename "$dir")
  echo -e "\033[0;34m🌀🌀🌀 Testing $dir with Maelstrom... 🌀🌀🌀\033[0m"
  output=$(make -C "$dir" test 2>&1)
  status=$?
  if [ "$status" -ne 0 ]; then
    echo "$output"
    echo -e "\033[0;31m🌀🌀🌀 Tests for $dir failed! 🌀🌀🌀\033[0m"
    exit 1
  fi
  echo -e "\033[0;32m🌀🌀🌀 Tests for $dir passed! 🌀🌀🌀\033[0m"
done
