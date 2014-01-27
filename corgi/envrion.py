#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

from os import envrion

class Environ(object):
    """Envrion"""

    def __init__(self):
        self.properties = {}

    def setenv(self):
        print envrion
