.. _configuration:

========
配置文件
========


任务节点配置
============

公共配置
--------

name
^^^^

节点名。

type
^^^^

节点类型。

output
^^^^^^

节点输出路径。可以使用变量，例如：``output="/tmp/example/${bizdate}/"``。

vars
^^^^

定义变量。语法如下：

* 分号分割的若干赋值语句

* 左值

  $开头的变量名

* 右值

  * int整数

  * 字符串

    单引号扩起来的字符串：``$name='test'``

  * 其它变量

    $号开头的变量名：``$date=$bizdate``

  * 时间变量扩展

    按照格式获取时间：``$date=${date:YYYYMMDD}``

nonstrict
^^^^^^^^^

不要求前置条件全成功。例如：``nonstrict="true"``。

节点类型
--------

* hadoop
* hive
* script

hadoop
^^^^^^

hive
^^^^

hql
"""

HiveQL命令。

script
""""""

HiveQL文件。

option
""""""

Hive选项，例如：``hive.exec.dynamic.partition.mode=nonstrict``。

script
^^^^^^

interpreter
"""""""""""

解释器，例如：``perl``，``bash``。可以只写出在系统变量``$PATH``内命令，也可以写完整的路径。

script
""""""

Script文件。

内置变量
========

日期时间
--------

gmtdate
^^^^^^^

脚本运行当前的日期。

bizdate
^^^^^^^

脚本运行前一天的日期。可以在启动hpipe的时候用``--bizdate``指定日期，这个选项会同时更新``gmtdate``。

job_report
^^^^^^^^^^

所有任务运行报告。
