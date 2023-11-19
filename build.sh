#!/bin/sh
START=$(pwd)
alias compile='go120 build'

cd rootexec/ && compile && echo 'Built: rootexec'
cd $START

cd rootd/ && compile && echo 'Built: rootd'
cd $START

cd rootctl/ && compile && echo 'Built: rootctl'
cd $START

exit 0