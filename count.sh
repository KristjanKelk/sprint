#!/bin/bash

# Count the number of files and directories
count=$(find . -type f -o -type d | grep -v '/\.' | wc -l)

# Multiply the count by 5
result=$((count * 5))

# Create the usage string
usage="Usage: "

# Use 'printf' to format and print the result
printf "\t%s%d\t\n" "$usage" "$result"