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
    data = {}
    if len(tokens) == 3:
        click_id = tokens[0]
        data["pv_id"] = tokens[1]
        data["shown_ad_list"] = tokens[2]
        if not click_id in pvclick:
            pvclick[click_id] = data
        else:
            pvclick[click_id].update(data)
    elif len(tokens) == 2:
        click_id = tokens[0]
        data["clicked_ad_list"] = tokens[1]
        if not click_id in pvclick:
            pvclick[click_id] = data
        else:
            pvclick[click_id].update(data)

for click_id, data in pvclick.items():
    print "%s\t%s\t%s" % (data["pv_id"], data["shown_ad_list"],
                          data["clicked_ad_list"])
