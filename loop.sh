# Ensure an argument is provided
if [ "$#" -eq 0 ]; then
  echo "Please provide the number of times to run the loop as an argument."
  exit 1
fi

# Get the number of times to run the loop (limited to 100)
num_times=$1
if [ "$num_times" -gt 100 ]; then
  num_times=100
fi

# Loop to print the text and the Nth time it has done so
for ((i=1; i<=$num_times; i++)); do
  echo "$i times I've printed kristjankelk"
done