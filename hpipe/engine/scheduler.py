#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

class Scheduler(object):

    def get_runnables(self, node, runnables):
        pass

class SimpleScheduler(Scheduler):

    def get_runnables(self, node, runnables):
        pending = []
        for dep in node.depends:
            if dep.job != None and dep.state == "RUNNABLE" \
               and len(dep.depends) == 0:
                # TODO: check DONE flags
                runnables.append(dep)
            else:
                pending.append(dep)
                self.get_runnables(dep, runnables)
        node.depends = pending
