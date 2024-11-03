#!/bin/bash

output_path=/tmp/maelstrom-test-all-output-$(date +"%H-%M-%S")
echo -e "\033[0;34m🌀🌀🌀 tail -f $output_path to see output 🌀🌀🌀\033[0m"

for dir in */; do
  dir=$(basename "$dir")
  echo -e "\033[0;34m🌀🌀🌀 Testing $dir with Maelstrom... 🌀🌀🌀\033[0m"
  output=$(make -C "$dir" test 2>&1 | tee $output_path)
  status=$?
  if [ "$status" -ne 0 ]; then
    echo "$output"
    echo -e "\033[0;31m🌀🌀🌀 Tests for $dir failed! 🌀🌀🌀\033[0m"
    exit 1
  fi
  echo -e "\033[0;32m🌀🌀🌀 Tests for $dir passed! 🌀🌀🌀\033[0m"
done

read -p "Delete $output_path? (y/n) [Default - y]: " answer
if [[ "$answer" != "n" ]]; then
  echo "rm $output_path"
  rm $output_path
fi
