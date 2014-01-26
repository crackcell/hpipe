#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

class Task(object):
    """Leave node in dependence tree"""

    def __init__(self, name="", mapper="", reducer=""):
        self.name = name
        self.mapper = mapper
        self.reducer = reducer
        self.files = []

    def add_file(self, filename):
        self.files.append(filename)

    def debug_string(self):
        print "Task:\n" + \
            "  name: " + self.task_name + \
            "  mapper: " + self.mapper + \
            "  reducer: " + self.reducer
        for filename in self.files:
            print "  file: " + filename

class Flow(object):
    """Root and middle node in dependence tree"""

    def __init__(self, name=""):
        self.name = name
        self.nodes = []

    def debug_string(self):
        str = "Flow:\n" + "  name: " + self.name;
        for node in self.nodes:
            str = str + task.debug_string()
        return str

class Node(object):
    """Node"""

    def __init__(self, name="", resource=""):
        self.name = name
        self.resource = resource

    def debug_string(self):
        return "Node:\n" + \
            "  name: " + self.name + "\n" + \
            "  resource: " + self.resource
