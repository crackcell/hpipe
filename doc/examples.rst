========
Examples
========

Wordcount: using perl
=====================

This chapter shows a wordcount example using hpipe and perl. The flow contains
four sub-flows. They have no data dependencies to each other. In order to show
the usage of flow dependency management, I make step3 depend on step2 and step2
depend on step1.0 and step1.1's output.

Files
-----

* Flow entry: examples/wordcount/wordcount.tml
* Sub tasks: examples/wordcount/task/step.[1-3].xml
* Sub task files (include mapper and reducer scripts): examples/wordcount/task/step.[1-3]/
* Data: examples/wordcount/data/part-00000

Steps
-----

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

Input
-----

Click log:

``click log id``\\t ``clicked ad list``

* Click log id: the primary key.
* Clicked ad list: a list of AD id clicked in a single pageview.

PV log:

``pv log id``\\t ``click log id``\\t ``displayed ad list``

* PV log id: the primary key.
* Click log id: the foreign key to click log.
* Displayed AD list: ADs shown in a single result list.

Output
------

``pv log id``\\t ``displayed ad list``\\t ``clicked ad list``

Files
-----

* Flow entry: examples/join/join.xml
* Task: examples/join/task/join.xml
* Multiple input mapper: hpipe/smartflow/multiple_input/mi_mapper.py
* PV log mapper: examples/join/task/pv_mapper.py
* Click log mapper: examples/join/task/click_mapper.py
* Reducer: examples/join/task/reducer.py

Steps
-----

1. Upload data files

| hadoop fs -mkdir /hpipe/join/pvlog /hpipe/join/clicklog
| hadoop fs -put examples/join/data/pvlog.txt /hpipe/join/pvlog/part-00000
| hadoop fs -put examples/join/data/clicklog.txt /hpipe/join/clicklog/part-00000
|

2. Run the flow

| ./bin/run.sh examples/join/join.xml
|
