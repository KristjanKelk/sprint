#!/bin/bash

# Use 'find' to list all unhidden files in the current directory and its subdirectories
usage="Number of unhidden files: "
file_count=$(find . -type f -not -path '*/\.*' | wc -l)

result=$(($file_count * 5))


# Display the result
printf "\t$usage\t%d\n\n" "$result"
