#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    reducer
    ~~~~~~~

    desc

    :copyright: (c) 2014 Menglong TAN.
"""

import sys

pvclick = {}

for line in sys.stdin:
    tokens = line.strip().split("\t")
    if len(tokens) != 3:
        continue
    ad_id = tokens[0]
    if not ad_id in pvclick:
        pvclick[ad_id] = []
        pvclick[ad_id].append(int(tokens[1]))
        pvclick[ad_id].append(int(tokens[2]))
    else:
        pvclick[ad_id][0] += int(tokens[1])
        pvclick[ad_id][1] += int(tokens[2])

for ad_id in pvclick.keys():
    if pvclick[ad_id][0] == 0:
        ctr = 0
    else:
        ctr = pvclick[ad_id][1] / float(pvclick[ad_id][0])
    print "%s\t%f" % (ad_id, ctr)
