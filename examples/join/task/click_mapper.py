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
    if len(tokens) != 2:
        return
    print line
