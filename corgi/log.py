#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import logging
import sys

logger = logging.getLogger("corgi")
formatter = logging.Formatter('%(asctime)s %(levelname)s %(name)s:  %(message)s',
                              '%y/%m/%d %H:%M:%S')
stream_handler = logging.StreamHandler(sys.stderr)
stream_handler.setFormatter(formatter)
logger.addHandler(stream_handler)
