======================
Quickstart by examples
======================

Wordcount: hpipe with perl
==========================

This chapter shows a wordcount example using hpipe and perl. The flow contains
four sub-flows. They have no data dependencies to each other. In order to show
the usage of flow dependency management, I make step3 depend on step2 and step2
depend on step1.0 and step1.1's output.

Files:

* Flow entry: demo/wordcount/wordcount.tml
* Sub tasks: demo/wordcount/task/step.[1-3].xml
* Sub task files (include mapper and reducer scripts): demo/wordcount/task/step.[1-3]/
* Demo data: demo/wordcount/data/part-00000

Steps:

1. Setup your own Hadoop environment.
2. Upload data file

   $ hadoop fs -mkdir hpipe/wordcount/input/
   $ hadoop fs -put demo/wordcount/data/part-00000 hpipe/wordcount/input/

3. Run the demo

   $ ./bin/run.sh demo/wordcount/wordcount.xml
