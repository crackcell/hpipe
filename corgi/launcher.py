#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import os
import logging

from subprocess import Popen
from entity import *
from scheduler import *
from client import *

logger = logging.getLogger("corgi")

class Launcher(object):
    """Launch jobs to hadoop"""

    def __init__(self):
        self.scheduler = SimpleScheduler()
        self.client = HadoopClient()

    def launch(self, tree):
        processes = []
        runnables = []
        self.scheduler.get_runnables(tree, runnables)
        while len(runnables) != 0:
            for node in runnables:
                # TODO: check flags
                pass
            self.client.launch(runnables)
            for node in runnables:
                # TODO: touch DONE flags
                pass
            runnables = []
            self.scheduler.get_runnables(tree, runnables)
