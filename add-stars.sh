#! /bin/bash

set -e

NUMBER_TO_ADD=$1

README_FILE=README.md
PATTERN="message=\d"

# Find original number of stars and add the required number of new stars
REPLACEMENT=$(grep -E -o -e $PATTERN $README_FILE | awk -v add="$NUMBER_TO_ADD" '{split($0,a,"="); print a[1] "="  a[2] + add}')

# Replace the number of stars in file
# perl is used here because sed is not portable in MacOS
perl -i -pe"s/$PATTERN/$REPLACEMENT/g" $README_FILE
