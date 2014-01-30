#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

from subprocess import Popen, STDOUT
from entity import Node, Job

import os

class JobLauncher(object):
    """Launch jobs to hadoop"""

    def __init__(self):
        # Load hadoop config from system envrionment variables
        self.hadoop_home = os.environ["hadoop_home"]
        self.hadoop_exec = os.environ["hadoop_exec"]
        self.hadoop_exec_conf = os.environ["hadoop_exec_conf"]
        self.hadoop_streaming_jar = os.environ["hadoop_streaming_jar"]

    def launch(self, tree):
        command = self.hadoop_exec + " jar " + self.hadoop_streaming_jar
        processes = []

        runnables = []
        self.__get_runnable(tree, runnables)
        while len(runnables) != 0:
            for job in runnables:
                # TODO: check filesystem flags before launching job
                processes.append(Popen(command, shell=True))
                # TODO: touch BUSY flag after launching job
            for p in processes:
                p.wait()
            for job in runnables:
                pass
                # TODO: touch DONE flags

            runnables = []
            self.__get_runnable(tree, runnables)

    def __get_runnable(self, node, runnables):
        pending = []
        for dep in node.depends:
            if dep.job != None and dep.state == "RUNNABLE" \
               and len(dep.depends) == 0:
                # TODO: check DONE flags
                runnables.append(dep)
            else:
                pending.append(dep)
                self.__get_runnable(dep, runnables)
        node.depends = pending
