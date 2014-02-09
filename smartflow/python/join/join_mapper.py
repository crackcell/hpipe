#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import os
import sys

class JoinMapper(object):

    def __init__(self):
        self.sub_mappers = {}
        for input_name in os.environ["hpipe_flow_join_inputs"].split(","):
            input_name = input_name.strip()
            input_path = os.environ["hpipe_flow_join_%s_input_dir" % input_name]
            mapper_file = os.environ["hpipe_flow_join_%s_mapper" % input_name]
            self.sub_mappers[input_path] = mapper_file

    def map(self, line):
        input_path = os.environ["map_input_file"]
        if not input_path in self.sub_mappers:
            raise RuntimeError("invalid input path: %s" % input_path)
        mapper_file = self.sub_mappers[input_path]

if __name__ == "__main__":
    mapper = JoinMapper()
    for line in sys.stdin:
        # remove leading and trailing whitespace
        line = line.strip()
