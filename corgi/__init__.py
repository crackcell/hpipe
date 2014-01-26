#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import logging
import sys

logger = logging.getLogger("corgi")
formatter = logging.Formatter('%(name)-12s %(asctime)s %(levelname)-8s %(message)s',
                              '%a, %d %b %Y %H:%M:%S',)
stream_handler = logging.StreamHandler(sys.stderr)
logger.addHandler(stream_handler)
