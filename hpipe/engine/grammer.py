#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    grammer
    ~~~~~~~

    Variables resovler for workflow definition config.

    :copyright: (c) 2014 Menglong TAN.
"""

import re
import logging

from datetime import datetime, timedelta

from hpipe.util import setup_logger

logger = logging.getLogger(__name__)
setup_logger(logger)

class ReferenceResolver(object):

    def __init__(self):
        self.space = {}
        self.pattern = re.compile("\$\{(.+?)([-+]*)(\d*)(d*)\}")
        self.resolved = False

    def add_space(self, space):
        self.space.update(space)

    def resolve(self):
        changed = True
        while changed:
            changed = self.__resolve_space(self.space)
        self.resolved = True

    def eval_prop(self, prop):
        if not self.resolved:
            raise RuntimeError("call resolve() before setting prop")
        for k in prop.iterkeys():
            if k in self.space:
                prop[k] = self.space[k]
        return prop

    def __calc_value(self, match, newvalue, oldvalue):
        if match[3] == "":     # normal numeric calculation (no unit)
            if match[1] == "-":
                newvalue = str(int(result) - int(match[2]))
            elif match[1] == "+":
                newvalue = str(int(result) + int(match[2]))
        else:                  # datetime caculation
            orig_time = datetime.strptime(newvalue, "%Y%m%d")
            if match[3] == "d":   # day diff
                diff_time = timedelta(days=int(match[2]))
            elif match[3] == "w": # week diff
                diff_time = timedelta(weeks=int(match[2]))
            else:
                raise RuntimeError("invalid datettime unit: %s" %
                                   match[3])
            if match[1] == "-":
                newvalue = (orig_time - diff_time).strftime("%Y%m%d")
            elif match[1] == "+":
                newvalue = (orig_time + diff_time).strftime("%Y%m%d")
        return oldvalue.replace(
            "${" + match[0] + match[1] + match[2] + match[3] + "}", newvalue)

    def __resolve_space(self, space):
        changed = False
        for name, value in space.iteritems():
            # regex: \$\{(.+?)([-+]*)(\d*)(d*)\}
            # format: match[0]   match[1]   match[2]   match[3]
            #         var_name   +/-        number     unit
            for match in self.pattern.findall(value):
                changed = True
                var = match[0]
                if len(match) != 1 and len(match) != 3 and len(match) != 4:
                    raise RuntimeError("invalid property value: %s" %
                                       value)
                result = None

                if var in space:
                    result = space[var]
                else:
                    raise RuntimeError("undefined variable: ${%s}" % var)

                orig = space[name]
                space[name] = self.__calc_value(match, result, space[name])
                logger.debug("resolving $%s: %s -> %s", name, orig,
                             space[name])

        return changed
