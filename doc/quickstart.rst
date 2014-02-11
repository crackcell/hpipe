==========
Quickstart
==========

Setup config
============

Edit ``conf/hpipe.env``, specify config according to the underlaying systems

for Hadoop, you **MUST** specify:
- hadoop_home: path to Hadoop client.
- hadoop_streaming_jar: path to streaming jar.
- hadoop_conf: Hadoop client config file. may be located in users' home folder.
