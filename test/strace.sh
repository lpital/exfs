#! /bin/bash

rm -r func.txt
touch func.txt
cat $1|
uniq |
sort |
uniq > func.txt
