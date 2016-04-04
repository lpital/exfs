#! /bin/bash

bin/hdfs dfs -rm -r input
bin/hdfs dfs -mkdir input
bin/hdfs dfs -rm -r output
bin/hdfs dfs -mkdir output
bin/hdfs dfs -put etc/hadoop/* input
bin/hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-2.7.0.jar grep input output 'dfs[a-z.]+'
rm -r output
mkdir output
bin/hdfs dfs -get output output
cat output/*

