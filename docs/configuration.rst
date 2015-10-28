========
配置文件
========

工作流
======

示例
----

配置文件::

   digraph wordcount {
       dummy -> step1
       dummy -> step2
       step1 [
             name="step1"
             type="hadoop"
             vars="$date=${date:YYYYMMDD}"
             input="/tmp/hpipe/examples/wordcount/input/part-*"
             output="/tmp/hpipe/examples/wordcount/output/${bizdate}/step1"
             mapper="cat"
             reducer="wc -l"
             mapred.reduce.tasks=1
             test.custom.val="today is ${date}, yestoday is ${bizdate}, and yestoday is ${bizdate}"
             ]
       step2 [
             name="step2"
             type="hadoop"
             vars="$date=$bizdate"
             input="/tmp/hpipe/examples/wordcount/input/part-*"
             output="/tmp/hpipe/examples/wordcount/output/${bizdate}/step2"
             mapper="cat"
             reducer="wc -l"
             ]
       dummy [
             name="dummy"
             type="dummy"
             vars="$bizdate=${date:YYYYMMDD}+2*$day"
             ]
   }

任务节点
========

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

定义变量
""""""""

语法如下：

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

使用变量
""""""""

在其它配置项里面使用变量，如 ``output`` ， ``script`` 等里面。

例如：

* ``output="/tmp/${bizdate}"``
* ``script="test.sh ${bizdate}"``

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

Hive选项，例如： ``hive.exec.dynamic.partition.mode=nonstrict`` 。

script
^^^^^^

interpreter
"""""""""""

解释器，例如： ``perl`` ， ``bash`` 。可以只写出在系统变量 ``$PATH`` 内命令，也可以写完整的路径。

script
""""""

Script文件。

内置变量
--------

日期时间
^^^^^^^^

gmtdate
"""""""

脚本运行当前的日期。

bizdate
"""""""

脚本运行前一天的日期。可以在启动hpipe的时候用 ``--bizdate`` 指定日期，这个选项会同时更新 ``gmtdate`` 。

job_report
""""""""""

所有任务运行报告。

节点间依赖
==========

大多数情况下，只使用``->``即可定义依赖。但某些特殊功能需要单独定义。

nonstrict
---------

非强制依赖。默认依赖，若上游节点失败，不会继续执行，非强制依赖表示可以无视上游节点状态，继续执行。

具体例子可以参考 ``examples/email`` ::

   digraph email {
       dummy -> script
       script -> mailitsh [nonstrict="true"]
       mailitsh -> mailitpl
       (略)
   }
