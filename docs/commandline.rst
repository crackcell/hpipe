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
       --start            Start from a specific job
       --end              End at a specific job
       --rerun            Rerun finished jobs, default: false
       --force            Run job even if already marked started, default: false

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

任务配置
--------

* ``-p`` 或者 ``--path`` ：工作路径，配置项中的文件路径会被加上此前缀
* ``-f`` 或者 ``--flow`` ：任务流入口文件

* ``--max-retry`` ：任务失败会被重新执行，此参数指定任务最大执行次数
* ``--start`` ：从某个任务节点开始执行
* ``--end`` ：执行到某个任务节点结束
* ``--bizdate`` ：指定内置变量 ``$bizdate`` ，设定此值会同时更新 ``$gmtdate``
* ``--rerun`` : 重跑已被标记为finished的任务
* ``--force`` : 强制重跑标记为started的任务

任务状态记录
------------

目前Hpipe支持2种存放任务状态：

#. SQLite（默认方式）
#. HDFS

在启动hpipe的时候通过命令行参数指定：``--status-saver``。

若使用SQLite，需要指定数据库文件路径，例如：``--sqlite /var/run/hpipe/status.db``

若使用HDFS，需要指定HDFS的Namenode服务器地址：``--namenode master:8020``

任务类型开关
------------

Hpipe支持多种任务类型的计算任务，每种任务有特定的开关。

Hadoop Streaming
^^^^^^^^^^^^^^^^

* ``--hadoop`` ：开启Hadoop支持
* ``--jar`` ：指定Hadoop streaming的jar包路径

Hive
^^^^

* ``--hive`` ：开启Hive支持

Aliyun ODPS
^^^^^^^^^^^

* ``--odps`` ：开启ODPS支持
* ``--odps-endpoing`` ：指定ODPS endpoint
* ``--odps-project`` ：指定ODPS project
* ``--odps-access-id`` ：指定ODPS access id
* ``--odps-access-key`` ：指定ODPS access key
