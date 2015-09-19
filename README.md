     _______         __
    |   |   |.-----.|__|.-----.-----.
    |       ||  _  ||  ||  _  |  -__|
    |___|___||   __||__||   __|_____|
             |__|       |__|

# Hpipe
[![Build Status](https://travis-ci.org/crackcell/hpipe.svg?branch=master)](https://travis-ci.org/crackcell/hpipe)
[![Coverage Status](https://coveralls.io/repos/crackcell/hpipe/badge.svg?branch=master&service=github)](https://coveralls.io/github/crackcell/hpipe?branch=master)

Hpipe is a workflow engine supporting hybrid workflows with built-in support for
Hadoop Streaming and Hive.

It is useful in many production scenarios such as offline data processing of
click-through-rate prediction for online advertising.

## Features

- Integrate various types of jobs into a uniformed workflow.
- Job status tracking mechanism: record job status and auto-resume from break point.
- Workflow descriptive language based on [Graphviz](http://graphviz.org/) being easy to edit and visualize.

## Quick Start

1. Compile

> make

2. Run example

> ./output/bin/hpipe-run -v -p $PWD/examples/wordcount -f ./examples/wordcount/wordcount.dot --namenode namenode:8020 --sqlite ./status.db
