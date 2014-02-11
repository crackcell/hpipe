==========
Quickstart
==========

Setup config
============

Edit ``conf/hpipe.env``, specify config according to the underlaying systems

for Hadoop, you **MUST** specify:

* $hadoop_home: path to Hadoop client.
* $hadoop_streaming_jar: path to streaming jar.
* $hadoop_conf: Hadoop client config file. may be located in users' home folder.

Create flow config
==================

Create flow config
------------------

for example:

.. code-block:: xml

    <?xml version="1.0"?>
    <flow name="wordcount">
      <node>
        <name>step1</name>
        <resource>examples/wordcount/task/step1.xml</resource>
        <property>
          <name>bizdate</name>
          <value>${today-1d}</value>
        </property>
      </node>
    </flow>

Create task config
------------------

for example:

.. code-block:: xml

    <?xml version="1.0"?>
    <job>
      <property>
        <name>hpipe.job.name</name>
        <value>step1</value>
      </property>
      <property>
        <name>mapred.job.name</name>
        <value>${hpipe.job.name}: ${bizdate}</value>
      </property>
      <property>
        <name>hpipe.input.dir</name>
        <value>/hpipe/wordcount/input/part-*</value>
      </property>
      <property>
        <name>hpipe.output.dir</name>
        <value>/hpipe/wordcount/${bizdate}/step1</value>
      </property>
      <property>
        <name>hpipe.mapper</name>
        <value>examples/wordcount/task/step1/mapper.pl</value>
      </property>
      <property>
        <name>hpipe.reducer</name>
        <value>examples/wordcount/task/step1/reducer.pl</value>
      </property>
      <property>
        <name>hpipe.file</name>
        <value>
          <array>examples/wordcount/task/step1/mapper.pl</array>
          <array>examples/wordcount/task/step1/reducer.pl</array>
        </value>
      </property>
    </job>

Use variables
^^^^^^^^^^^^^

Please review config above, there are two varibales used, ``$today`` and
``bizdate``. The former is defined in environment variables, you may find it
in ``bin/run.sh``. The later defined in the flow config. The flow ``wordcount``
defines a task ``step1`` and attach a property setting ``$bizdate``.

Hpipe loads variable from:

* Environment variables
* Flow config
* Task config
