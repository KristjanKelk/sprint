directory="theDirectory/left/down/beginning"
file_to_print="README"
cd "$directory" && cat "$file_to_print" || echo "Error: Directory or file not found."