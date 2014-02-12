#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    click_mapper
    ~~~~~~~~~~~~

    desc

    :copyright: (c) 2014 Menglong TAN.
"""

def map(line):
    tokens = line.split("\t")
    for ad_id in tokens[1].split(","):
        # format:
        #     ad_id pv click
        print "%s\t0\t1" % ad_id
