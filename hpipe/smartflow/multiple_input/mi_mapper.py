#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    mapper
    ~~~~~~

    Helper code to simplify join task.

    :copyright: (c) 2014 Menglong TAN.
"""

import os
import re
import sys
import imp
import fnmatch

filedir = os.path.abspath(os.path.join(os.path.dirname(__file__), '.'))

class MultipleInputMapper(object):

    def __init__(self):
        self.sub_mappers = {}
        for input_name in os.environ["hpipe_smartflow_mi_inputs"].split(","):
            input_name = input_name.strip()
            input_path = os.environ["hpipe_smartflow_mi_%s_input_dir" % input_name]
            path_pattern = re.compile(".*?" + fnmatch.translate(input_path),
                                      re.IGNORECASE)
            mapper_path = os.path.join(
                filedir,
                os.path.basename(os.environ["hpipe_smartflow_mi_%s_mapper" % input_name]))
            mapper = self.__load_mapper(input_name, mapper_path)
            self.sub_mappers[input_path] = (path_pattern, mapper)

    def map(self, line):
        input_path = os.environ["map_input_file"]
        mapper = self.__find_mapper(input_path)
        mapper.map(line)

    def __find_mapper(self, input_path):
        for i in self.sub_mappers.keys():
            input_pattern, mapper = self.sub_mappers[i]
            if input_pattern.match(input_path):
                return mapper
        raise RuntimeError("invalid input path: %s. available paths %s" %
                           (input_path, self.sub_mappers.keys()))

    def __load_mapper(self, name, path):
        if path.endswith(".py"):
            return imp.load_source(name, path)
        elif path.endswith(".pyc"):
            return imp.load_compiled(name, path)


mapper = MultipleInputMapper()
for line in sys.stdin:
    # remove leading and trailing whitespace
    mapper.map(line.strip())
