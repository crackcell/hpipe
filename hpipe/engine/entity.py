#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

class Node(object):
    """Node"""

    def __init__(self, name="", resource="", properties={}):
        self.name = name
        self.resource = resource
        self.depends = []
        self.properties = properties
        self.job = None
        self.state = "RUNNABLE"
        self.returncode = 0

    def __repr__(self):
        str = self.name
        if self.job != None:
            str += ":\n" + self.job.__repr__()
        return str

class Job(object):
    """Job"""

    def __init__(self):
        self.properties = {}
        self.inputs = []
        self.files = []

    def __repr__(self):
        str = " " * 4 + "Job:{\n"
        str += " " * 8 + "prop:{\n"
        for k, v in self.properties.iteritems():
            str += " " * 12 + k + ": " + v + "\n"
        str += " " * 8 + "}\n"
        str += " " * 8 + "input:[\n"
        for i in self.inputs:
            str += " " * 12 + i + "\n"
        str += " " * 8 + "]\n"
        str += " " * 4 + "}"
        return str

    def validate(self):
        checklist = ["hpipe.job.name", "hpipe.output.dir", "hpipe.mapper",
                     "hpipe.reducer"]
        for p in checklist:
            if not p in self.properties.keys():
                raise RuntimeError(p)
        if len(self.inputs) == 0:
            raise RuntimeError("input")
