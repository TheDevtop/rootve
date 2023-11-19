#!/bin/sh
START=$(pwd)
alias compile='go120'

# Build rootexec
cd rootexec/ && compile && echo 'Built: rootexec'
cd $START

# Build rootd
cd rootd/ && compile && echo 'Built: rootd'
cd $START

# Build rootctl
cd rootctl/ && compile && echo 'Built: rootctl'
cd $START

exit 0