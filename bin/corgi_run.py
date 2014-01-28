#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import os
import sys
import getopt
import logging

from corgi import JobManager

logger = logging.getLogger("corgi")

def main():

    # Check python runtime version
    if sys.version_info[0] * 10 + sys.version_info[1] < 26:
        raise RuntimeError("invalid python version %s.%s.%s (< 2.6), "
                           "please upgrade it" %
                           (sys.version_info[0], sys.version_info[1],
                            sys.version_info[2]))

    logger.setLevel(int(os.environ["corgi_log_level"]))
    try:
        opts, args = getopt.getopt(sys.argv[1:], "hv", ["help", "version"])
    except getopt.GetoptError, err:
        usage()
        sys.exit(1)
    for o, a in opts:
        if o in ("-v", "--version"):
            version()
            sys.exit()
        elif o in ("-h", "--help"):
            usage()
            sys.exit()
    if len(sys.argv) != len(opts) + 2:
        usage()
        sys.exit()
    flow = sys.argv[len(opts) + 1]
    logger.info("flow: %s", flow)

    job_manager = JobManager()
    job_manager.load_conf(flow)
    job_manager.launch()

def usage():
    print "Usage: run.py [-h --help] [-v --version] FLOW"

def version():
    print "Corgi " + corgi.VERSION

if __name__ == "__main__":
    main()
