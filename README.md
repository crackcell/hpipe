     _______         __
    |   |   |.-----.|__|.-----.-----.
    |       ||  _  ||  ||  _  |  -__|
    |___|___||   __||__||   __|_____|
             |__|       |__|

# Hpipe v2

hpipe is a workflow engine for various computing systems (Hadoop, Aliyun ODPS etc.). The main purpose of this project is to simplify muilti-job tasks like click-through-rate prediction in online advertising systems.

Features include:

* Workflow definition language based on XML.
* Built-in failover machanism: auto-clean & auto-retry failed jobs.
* Join/Aggregation/Filter support.

## Quick Start

    $ make
    $ ./output

## Examples

    ./output/bin/hpipe-run -f wordcount.xml -w PATH-OF-HPIPE/examples/hadoop/wordcount -m ./ -v --rerun
