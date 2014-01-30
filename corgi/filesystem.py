#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

class Filesystem(object):
    """Filesystem interface"""

    def __init__(self):
        pass

    def touch(self, path):
        pass

    def mkdir(self, path):
        pass

    def rm(self, path):
        pass

class LocalFilesystem(Filesystem):
    pass

class HDFSFilesystem(Filesystem):
    """HDFS"""

    def __init__(self):
        # Load hadoop config from system envrionment variables
        self.hadoop_home = os.environ["hadoop_home"]
        self.hadoop_exec = os.environ["hadoop_exec"]
        self.hadoop_exec_conf = os.environ["hadoop_exec_conf"]
        self.hadoop_streaming_jar = os.environ["hadoop_streaming_jar"]
