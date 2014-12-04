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
    if len(tokens) != 3:
        return
    pv_id = tokens[0]
    click_id = tokens[1]
    ad_list = tokens[2]
    print "%s\t%s\t%s" % (click_id, pv_id, ad_list)
