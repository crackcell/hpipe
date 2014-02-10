Welcome to hpipe
================
The hpipe is a workflow engine for Hadoop streaming jobs. The main purpose of
this project is to simplify muliti-job tasks like click-through-rate prediction
in online advertising systems.

Features include:

* Workflow definition language based on XML.
* Built-in failover machanism: auto-clean & auto-retry failed jobs.
* Join/Aggregation/Filter support.

Quick start
-----------

* wordcount_: Count words.
* join_: Join PV log and Click log together.


.. _wordcount: demos/wordcount.html
.. _join: demos/join.html
