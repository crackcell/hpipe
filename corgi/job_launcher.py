#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

from subprocess import call
from entity import Node, Job

import os

class JobLauncher(object):
    """Launch jobs to hadoop"""

    def __init__(self):
        self.run_queue = []
        # Load hadoop config from system envrionment variables
        self.hadoop_home = os.environ["hadoop_home"]
        self.hadoop_exec = os.environ["hadoop_exec"]
        self.hadoop_exec_conf = os.environ["hadoop_exec_conf"]
        self.hadoop_streaming_jar = os.environ["hadoop_streaming_jar"]

    def launch(self, tree):
        self.__fill_run_queue(tree)
        #print call([self.hadoop_exec, "jar", self.hadoop_streaming_jar])

    def __fill_run_queue(self, node):
        for dep in node.depends:
            pass
