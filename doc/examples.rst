========
Examples
========

Wordcount: hpipe with perl
==========================

This chapter shows a wordcount example using hpipe and perl. The flow contains
four sub-flows. They have no data dependencies to each other. In order to show
the usage of flow dependency management, I make step3 depend on step2 and step2
depend on step1.0 and step1.1's output.

Files:

* Flow entry: examples/wordcount/wordcount.tml
* Sub tasks: examples/wordcount/task/step.[1-3].xml
* Sub task files (include mapper and reducer scripts): examples/wordcount/task/step.[1-3]/
* Data: examples/wordcount/data/part-00000

Steps:

1. Setup your own Hadoop environment.
2. Upload data file

| $ hadoop fs -mkdir /hpipe/wordcount/input/
| $ hadoop fs -put examples/wordcount/data/part-00000 /hpipe/wordcount/input/
|

3. Run the flow

| $ ./bin/run.sh examples/wordcount/wordcount.xml
|

Data joining
============

This example shows how to use Hpipe to accomplish data joining. There are two
kinds of log data: the PV log and Click log. The goal is to calculate each ad's
click-through-rate by joining two data.

Format of Click log:

``click log id``\\t ``pv log id``\\t ``clicked ad list (separated by ,)``

* Click log id: the primary key.
* PV log id: the foreign key to PV log.
* Clicked ad list: a list of AD id clicked in a single pageview.

Data:

| 0	0	ad3
| 1	2	ad1,ad3
| 2	1	ad0
|

Format of PV log:

``pv log id``\\t ``displayed ad list (separated by ,)``

Data:

| 0	ad0,ad3,ad2
| 1	ad1,ad2,ad3
| 2	ad2,ad1,ad0
|

* PV log id: the primary key
* Displayed AD list: ADs shown in a single result list

