#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    entity
    ~~~~~~

    Workflow entities.

    :copyright: (c) 2014 Menglong TAN.
"""

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
        str += self.__repr_prop("prop", self.properties)
        str += self.__repr_array("inputs", self.inputs)
        str += self.__repr_array("files", self.files)
        str += " " * 4 + "}"
        return str

    def __repr_prop(self, label, prop):
        str = " " * 8 + label + ":{\n"
        for k, v in prop.iteritems():
            str += " " * 12 + k + ": " + v + "\n"
        str += " " * 8 + "}\n"
        return str

    def __repr_array(self, label, array):
        str = " " * 8 + label + ":[\n"
        for i in array:
            str += " " * 12 + i + "\n"
        str += " " * 8 + "]\n"
        return str

    def validate(self):
        checklist = ["hpipe.job.name", "hpipe.output.dir", "hpipe.mapper",
                     "hpipe.reducer"]
        for p in checklist:
            if not p in self.properties.keys():
                raise RuntimeError(p)
        if len(self.inputs) == 0:
            raise RuntimeError("input")
