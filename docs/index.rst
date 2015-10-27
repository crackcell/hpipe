=====
Hpipe
=====

Hpipe是一个面向离线计算的工作流引擎。它具有一下功能：

* 支持混合任务流：目前支持Hadoop、Hive 和 自定义脚本。其它的计算平台可以通过插件方式快速支持
* 保存任务状态，实现任务的断点续跑：同样插件式的后端支持，可以保存在HDFS上，也可以保存在本地Sqlite
* 通俗易懂得给予Graphviz语法扩展的工作流定义，支持现有任意Graphviz图形编辑器，可以方便地转换成图片实现可视化
* 自定义变量，实现任务参数动态配置

目录
=====

.. toctree::
   :maxdepth: 2

   quickstart
   commandline
   configuration
   jobtracking

索引
====

* :ref:`genindex`
* :ref:`modindex`
* :ref:`search`

