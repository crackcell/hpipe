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
        self.jobconf = None
        self.depends = []

    def __repr__(self):
        str = self.name
        if self.jobconf != None:
            str += " mapper:" + self.jobconf.properties["mapper"] + \
            " reducer:" + self.jobconf.properties["reducer"]
        return str

class JobConf(object):
    """Job Config"""

    def __init__(self):
        self.properties = {}

    def __repr__(self):
        str = "JobConf:{"
        for k, v in self.properties.iteritems():
            str += " " * 10  + k + ":" + v + ","
        str += "}"
        return str

    def validate(self):
        checklist = ["mapper", "reducer"]
        for p in checklist:
            if not p in self.properties.keys():
                raise RuntimeError(p)
