#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import os
import logging

from subprocess import Popen
from hpipe.util import setup_logger

logger = logging.getLogger(__name__)
setup_logger(logger)

class Client(object):
    """Client interface"""

    def launch(self, tree):
        pass

class HadoopClient(Client):

    def __init__(self):
        # Load hadoop config from system envrionment variables
        self.hadoop_home = os.environ["hadoop_home"]
        self.hadoop_exec = os.environ["hadoop_exec"]
        self.hadoop_conf = os.environ["hadoop_conf"]
        self.hadoop_streaming_jar = os.environ["hadoop_streaming_jar"]

    def launch(self, runnables):
        processes = []
        returncode = []
        for node in runnables:
            if node.state != "RUNNABLE":
                continue
            command = self.__assemble_command(node.job)
            logger.debug("submit job: %s", command)
            processes.append(Popen(command, shell=True))
        for p in processes:
            p.wait()
        for i, p in enumerate(processes):
            runnables[i].returncode = p.returncode

    def __assemble_command(self, job):
        command = self.hadoop_exec + \
                  " --config " + self.hadoop_conf + \
                  " jar " + \
                  self.hadoop_streaming_jar

        for k, v in job.properties.iteritems():
            command += " -D " + k + "=\"" + v + "\""

        for i in job.inputs:
            command += " -input " + i

        command += " -output " + job.properties["hpipe.output.dir"] + \
                   " -mapper " + job.properties["hpipe.mapper"] + \
                   " -reducer " + job.properties["hpipe.reducer"]

        for f in job.files:
            command += " -file " + f

        return command
