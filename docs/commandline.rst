==========
命令行工具
==========

hpipe
======

概要::

   Execute a hpipe workflow
   Usage:
       hpipe [options]
   Options:
       -h, --help         Print this message
       -v, --verbose      Use verbose output
   
       -p, --path         Working path
       -f, --flow         Entry filename of workflow
       --max-retry        Max retry times of failed jobs, default: 3
       --run-at           Run from a specific job
   
       --status-saver     Method to track job status
                          default: sqlite
                          available: hdfs, sqlite
   
       --namenode         Address of Hadoop NameNode, default: 127.0.0.1:8020
       --sqlite           File path for sqlite database
   
       --hadoop           Enable Hadoop streaming job
       --jar              Path of Hadoop streaming jar file
   
       --odps             Enable ODPS job
       --odps-endpoint    Address of ODPS endpoing
       --odps-project     ODPS project name
       --odps-access-id   ODPS access id
       --odps-access-key  ODPS access key
   
       --hive             Enable Hive job
   
       --bizdate          Set variable $bizdate in YYYYMMDD format
   
       --less-log         Less log output
