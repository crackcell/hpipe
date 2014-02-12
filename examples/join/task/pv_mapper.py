#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    pv_mapper
    ~~~~~~~~~

    desc

    :copyright: (c) 2014 Menglong TAN.
"""

def map(line):
    tokens = line.split("\t")
    for ad_id in tokens[1].split(","):
        # format:
        #     ad_id pv click
        print "%s\t1\t0" % ad_id
