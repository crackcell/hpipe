#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import os
import sys
import getopt
import logging

from hpipe import consts

from hpipe.util import setup_logger
from hpipe.engine import JobManager

logger = logging.getLogger(__name__)
setup_logger(logger)

def main():
    # require python >= 2.6
    assert_runtime_version()
    assert_env_var()

    # parse commandline arguments
    try:
        opts, args = getopt.getopt(sys.argv[1:], "hv", ["help", "version"])
    except getopt.GetoptError, err:
        usage()
        sys.exit(1)
    for opt, arg in opts:
        if opt in ("-v", "--version"):
            version()
            sys.exit()
        elif opt in ("-h", "--help"):
            usage()
            sys.exit()
    if len(sys.argv) != len(opts) + 2:
        usage()
        sys.exit()
    flow = sys.argv[len(opts) + 1]
    logger.info("workroot: %s flow: %s", os.environ[consts.HPIPE_ENV_WORKROOT],
                flow)

    # launch job
    jobmgr = JobManager()
    jobmgr.load_conf(flow)
    jobmgr.launch()

def usage():
    print "Usage: run.py [-d --workroot] [-h --help] [-v --version] FLOW"

def version():
    print "Hpipe " + hpipe.VERSION

def assert_runtime_version():
    # Check python runtime version
    if sys.version_info[0] * 10 + sys.version_info[1] < 26:
        raise RuntimeError("invalid python version %s.%s.%s (< 2.6), "
                           "please upgrade it" %
                           (sys.version_info[0], sys.version_info[1],
                            sys.version_info[2]))

def assert_env_var():
    var_checklist = [consts.HPIPE_ENV_WORKROOT, consts.HPIPE_ENV_LOG_LEVEL,
                     consts.HPIPE_ENV_SCHEDULER, consts.HPIPE_ENV_CLIENT,
                     consts.HPIPE_ENV_FILESYSTEM]
    for v in var_checklist:
        if os.environ[v] == None:
            raise RuntimeError("missing environment variable $%s" % v)

    path_checklist = ["hadoop_home", "hadoop_exec", "hadoop_streaming_jar",
                      "hadoop_conf"]
    for v in path_checklist:
        if not os.path.exists(os.environ[v]):
            raise RuntimeError("invalid envrionment variable $%s,"
                               " path does not exist: %s" % (v, os.environ[v]))
