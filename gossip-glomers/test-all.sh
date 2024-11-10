#!/bin/bash
set -o pipefail

output_path=/tmp/maelstrom-test-all-output-$(date +"%H-%M-%S")
echo -e "\033[0;34m🌀🌀🌀 tail -f $output_path to see output 🌀🌀🌀\033[0m"

prompt_cleanup() {
  read -p "Delete $output_path? (y/n) [Default - y]: " answer
  if [[ "$answer" != "n" ]]; then
    echo "rm $output_path"
    rm $output_path
  fi
}

sigint_cleanup() {
  echo "Received SIGINT, deleting $output_path" 
  echo "rm $output_path"
  rm $output_path
  exit 1
}
trap sigint_cleanup INT

for dir in */; do
  dir=$(basename "$dir")
  echo -e "\033[0;34mTesting $dir...\033[0m"
  output=$(make -C "$dir" test 2>&1 | tee -a $output_path)
  status=$?
  if [ "$status" -ne 0 ]; then
    echo -e "\033[0;31mTests for $dir failed!\033[0m"
    prompt_cleanup
    exit 1
  fi
  echo -e "\033[0;32mTests for $dir passed!\033[0m"
done

echo -e "\033[0;34m🌀🌀🌀 All tests passed! 🌀🌀🌀\033[0m"
prompt_cleanup
