#!/bin/bash

for dir in */; do
  dir=$(basename "$dir")
  make -C "$dir" clean
done
