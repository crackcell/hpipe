     _______         __
    |   |   |.-----.|__|.-----.-----.
    |       ||  _  ||  ||  _  |  -__|
    |___|___||   __||__||   __|_____|
             |__|       |__|

# Hpipe

Hpipe is a workflow engine supporting hybrid workflows with built-in support for
Hadoop Streaming and Spark.

It is useful in many production scenarios such as offline data processing of
click-through-rate prediction for online advertising.

## Quick Start

1. Compile

    $ make
    $ ./output

2. Run example

    ./output/bin/hpipe-run -v -p example/wordcount -f examples/wordcount/wordcount.dot --namenode hpc0:8020 --jar /usr/lib/hadoop-mapreduce/hadoop-streaming.jar
