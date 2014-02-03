#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

class Node(object):
    """Node"""

    def __init__(self, name="", resource=""):
        self.name = name
        self.resource = resource
        self.depends = []
        self.job = None
        self.state = "RUNNABLE"
        self.returncode = 0

    def __repr__(self):
        str = self.name
        if self.job != None:
            str += " mapper:" + self.job.properties["hpipe.mapper"] + \
            " reducer:" + self.job.properties["hpipe.reducer"]
        return str

class Job(object):
    """Job"""

    def __init__(self):
        self.properties = {}
        self.files = []

    def __repr__(self):
        str = "Job:{"
        for k, v in self.properties.iteritems():
            str += " " * 10  + k + ":" + v + ","
        str += "}"
        return str

    def validate(self):
        checklist = ["hpipe.job.name", "hpipe.input.dir",
                     "hpipe.output.dir", "hpipe.mapper", "hpipe.reducer"]
        for p in checklist:
            if not p in self.properties.keys():
                raise RuntimeError(p)
