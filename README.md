     _______         __
    |   |   |.-----.|__|.-----.-----.
    |       ||  _  ||  ||  _  |  -__|
    |___|___||   __||__||   __|_____|
             |__|       |__|

# Hpipe

Hpipe是一个支持多种计算平台的工作流引擎。它的设计初衷是为了将在线广告系统中，诸如CTR预估线下流程
这样的，由多个零散的任务组成的工作流用一个统一的工具串起来。

目前Hpipe支持Hadoop、阿里云ODPS 和 普通shell 任务，如果想支持别的任务类型，也可以方便扩展。

功能：

* 混合任务类型的支持（Hadoop、阿里云ODPS）
* 基于XML的简单工作流语言
* 自动保存任务状态，支持 自动重试 以及 流程中断重启之后进度回复

## 快速开始

    $ make
    $ ./output

## 示例

### Hadoop wordcount

    ./output/bin/hpipe-run -f wordcount.xml -w PATH-OF-HPIPE/examples/hadoop/wordcount -m ./ -v --rerun

## 其它

* [设计文档](doc/design.md)

