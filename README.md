     _______         __
    |   |   |.-----.|__|.-----.-----.
    |       ||  _  ||  ||  _  |  -__|
    |___|___||   __||__||   __|_____|
             |__|       |__|

# Hpipe v2

hpipe is a workflow engine for various computing systems (Hadoop, Aliyun ODPS etc.). At first, the main purpose of this project is to simplify muilti-job tasks like click-through-rate prediction in online advertising systems. But by now, it's used widely in our daily data mining tasks besides CTR prediction.

Features include:

* Support for various job types (Hadoop, Aliyun ODPS and shell jobs).
* Simple workflow definition language based on XML.
* Tracking job status in persistent storage, thus we can save the flow progress.
* Built-in failover machanism: auto-clean & auto-retry failed jobs.
* [TODO] Web-based UI
* [TODO] Join/Aggregation/Filter support.

## Quick Start

    $ make
    $ ./output

## Examples

    ./output/bin/hpipe-run -f wordcount.xml -w PATH-OF-HPIPE/examples/hadoop/wordcount -m ./ -v --rerun
