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
        self.flow_children = []
        self.task_children = []

    def debug_string(self):
        print "Flow:\n" + \
            "  name: " + self.name;
        for task in self.task_children:
            print task.debug_string()
