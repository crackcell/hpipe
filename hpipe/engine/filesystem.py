#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    filesystem
    ~~~~~~~~~~

    Filesystem interface and implementations used by the launcher.

    :copyright: (c) 2014 Menglong TAN.
"""

import os
import logging

from subprocess import call
from hpipe.util import setup_logger

logger = logging.getLogger(__name__)
setup_logger(logger)

class Filesystem(object):
    """Filesystem interface"""

    def __init__(self):
        pass

    def touch(self, path, args=""):
        pass

    def test(self, path, args=""):
        pass

    def mkdir(self, path):
        pass

    def rm(self, path, args=""):
        pass

    def rmr(self, path, args=""):
        pass

class LocalFilesystem(Filesystem):
    pass

class HDFSFilesystem(Filesystem):
    """HDFS"""

    def __init__(self):
        # Load hadoop config from system envrionment variables
        self.hadoop_home = os.environ["hadoop_home"]
        self.hadoop_exec = os.environ["hadoop_exec"]
        self.hadoop_conf = os.environ["hadoop_conf"]
        self.hadoop_streaming_jar = os.environ["hadoop_streaming_jar"]

    def touch(self, path, args=""):
        """touchz"""
        command = self.hadoop_exec + \
                  " --config " + self.hadoop_conf + \
                  " fs -touchz " + args + " " + path
        return call(command, shell=True)

    def test(self, path, args=""):
        """test -e"""
        command = self.hadoop_exec + \
                  " --config " + self.hadoop_conf + \
                  " fs -test " + args + " " + path
        return call(command, shell=True)

    def rm(self, path, args=""):
        """rm"""
        command = self.hadoop_exec + \
                  " --config " + self.hadoop_conf + \
                  " fs -rm " + args + " " + path
        return call(command, shell=True)

    def rmr(self, path, args=""):
        """rmr"""
        command = self.hadoop_exec + \
                  " --config " + self.hadoop_conf + \
                  " fs -rmr " + args + " " + path
        return call(command, shell=True)
