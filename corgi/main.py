#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import sys
import getopt

from corgi import logger

def main():
    try:
        opts, args = getopt.getopt(sys.argv[1:], "hvf:", ["help"])
    except getopt.GetoptError as err:
        logger.error("invalid arguments")
    logger.error("invalid arguments")

if __name__ == "__main__":
    main()
