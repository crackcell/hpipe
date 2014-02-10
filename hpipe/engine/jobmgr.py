#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    jobmgr
    ~~~~~~

    Load workflow definitions and call the launcher to submit jobs.

    :copyright: (c) 2014 Menglong TAN.
"""

import os
import re
import logging

from xml.etree import ElementTree as ET

from hpipe.util import setup_logger
from hpipe.engine.entity import Node, Job
from hpipe.engine.flowloader import FlowLoader
from hpipe.engine.launcher import Launcher
from hpipe.engine.grammer import ReferenceResolver

logger = logging.getLogger(__name__)
setup_logger(logger)

class JobManager(object):
    """Job Manager"""

    def __init__(self):
        self.root = Node()
        self.launcher = Launcher()
        self.flowloader = FlowLoader()

        self.input_pattern = re.compile("hpipe\..*?input.dir")

    def load_conf(self, conf_file):
        self.flowloader.parse_file(self.root, conf_file)
        logger.debug("node tree:")
        self.__log_tree(self.root)

    def launch(self):
        logger.info("launching jobs")
        self.launcher.launch(self.root)

    def __log_tree(self, node, level=0):
        str = node.__repr__()
        lines = str.split("\n")
        for i in range(len(lines)):
            prefix = "| "
            if i == 0:
                prefix = "+ "
            logger.debug("|    " * level + prefix + lines[i])
        for n in node.depends:
            self.__log_tree(n, level + 1)
