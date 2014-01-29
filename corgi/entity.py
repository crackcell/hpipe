#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import os

from threading import Thread
from subprocess import call
from multiprocessing import Process, Lock

class Node(object):
    """Node"""

    def __init__(self, name="", resource=""):
        self.name = name
        self.resource = resource
        self.depends = []
        self.job = None
        self.launcher = None

    def __repr__(self):
        str = self.name
        if self.job != None:
            str += " mapper:" + self.job.properties["mapper"] + \
            " reducer:" + self.job.properties["reducer"]
        return str

class Job(object):
    """Job"""

    def __init__(self):
        self.properties = {}

    def __repr__(self):
        str = "Job:{"
        for k, v in self.properties.iteritems():
            str += " " * 10  + k + ":" + v + ","
        str += "}"
        return str

    def validate(self):
        checklist = ["mapper", "reducer"]
        for p in checklist:
            if not p in self.properties.keys():
                raise RuntimeError(p)
