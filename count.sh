count=$(find . -type f -o -type d -not -path '*/\.*' | wc -l)


result=$((count * 5))

printf "\tTotal files * 5: %d\t\n" "$result"