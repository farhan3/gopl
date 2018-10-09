#!/bin/bash

# clean binary files
for d in $(find . -path ./.git -prune -o -type d -print); do
	echo $d
	cd $d
	go clean
	cd -
done
