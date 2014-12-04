#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    sort_mapper
    ~~~~~~~~~~~

    desc

    :copyright: (c) 2014 Menglong TAN.
"""

import os
import sys

class SortMapper(object):

    def __init__(self):
        self.splits = int(
            os.environ["hpipe_smartflow_eval_auc_splits"])
        if not 100 % int(self.splits) == 0:
            raise RuntimeError("invalid hpipe.smartflow.eval.auc.splits: %s"
                               % self.splits)

    def map(self, line):
        print self.splits

mapper = SortMapper
for line in sys.stdin:
    # remove leading and trailing whitespace
    mapper.map(line.strip())
