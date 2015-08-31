#!/bin/bash
##! @description: usage
##! @version: 1
##! @author: crackcell <tanmenglong@gmail.com>
##! @date:   Thu Aug 27 20:53:23 2015

echo "step1"
hadoop jar /usr/lib/hadoop-mapreduce/hadoop-streaming-2.6.0-cdh5.4.5.jar \
       -input /tmp/hpipe/example/wordcount/input/part-* \
       -output /tmp/hpipe/example/wordcount/output/step1 \
       -mapper cat \
       -reducer "wc -l"
