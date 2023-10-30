
# Create files a, !, \, and "
touch a '!' '\' '"'

# Create a directory called `
mkdir '`'

# Copy the file ! into `
cp '!' '`'

# Check the MOVE_A environment variable
if [ "$MOVE_A" == "yes" ]; then
    # Move the file a into the ` directory
    mv a '`'
elif [ "$MOVE_A" == "no" ]; then
    # Remove the file a
    rm a
fi