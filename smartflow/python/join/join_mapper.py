#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    join_mapper
    ~~~~~~~~~~~

    Helper code to simplify join task.

    :copyright: (c) 2014 Menglong TAN.
"""

import os
import sys
import fnmatch
import re

class JoinMapper(object):

    def __init__(self):
        self.sub_mappers = {}
        for input_name in os.environ["hpipe_flow_join_inputs"].split(","):
            input_name = input_name.strip()
            input_path = os.environ["hpipe_flow_join_%s_input_dir" % input_name]
            mapper_file = os.environ["hpipe_flow_join_%s_mapper" % input_name]
            input_path_pattern = re.compile(".*?" + fnmatch.translate(input_path))
            self.sub_mappers[input_path] = (mapper_file, input_path_pattern)

    def map(self, line):
        input_path = os.environ["map_input_file"]
        mapper_file = self.__find_mapper(input_path)
        print input_path, "\t", mapper_file

    def __find_mapper(self, input_path):
        for i in self.sub_mappers.keys():
            mapper_file, pattern = self.sub_mappers[i]
            if pattern.match(input_path):
                return mapper_file
        raise RuntimeError("invalid input path: %s. available paths %s" %
                           (input_path, self.sub_mappers.keys()))

mapper = JoinMapper()
for line in sys.stdin:
    # remove leading and trailing whitespace
    mapper.map(line.strip())
