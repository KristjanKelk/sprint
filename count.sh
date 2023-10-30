#!/bin/bash

# Count the number of files and directories
count=$(find . -type f -o -type d | grep -v '*/\.*' | wc -l)

# Multiply the count by 5
result=$((count * 5))


# Use 'printf' to format and print the result
printf "\tTotal files * 5: %d\t\n" "$result"