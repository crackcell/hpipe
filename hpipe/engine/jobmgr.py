#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

import os
import re
import logging

from xml.etree import ElementTree as ET
from datetime import datetime, timedelta

from hpipe.util import setup_logger
from hpipe.engine.entity import Node, Job
from hpipe.engine.launcher import Launcher
from hpipe.engine.grammer import ReferenceResolver

logger = logging.getLogger(__name__)
setup_logger(logger)

class JobManager(object):
    """Job Manager"""

    def __init__(self):
        self.root = Node()
        self.launcher = Launcher()
        self.pattern = re.compile("\$\{(.+?)([-+]*)(\d*)(d*)\}")

    def load_conf(self, conf_file):
        self.__parse_file(self.root, conf_file, {})
        logger.debug("node tree:")
        self.__log_tree(self.root)

    def launch(self):
        logger.info("launching jobs")
        self.launcher.launch(self.root)

    def __parse_file(self, node, conf_file, local_space):
        xmlroot = ET.parse(conf_file).getroot()

        if xmlroot.tag == "flow":
            self.__parse_flow(node, xmlroot, local_space)
        elif xmlroot.tag == "job":
            self.__parse_job(node, xmlroot, local_space)
        else:
            raise RuntimeError("invalid xml tag '%s'" % xmlroot.tag)

    def __parse_flow(self, node, xmlroot, local_space):

        try:
            node.name = xmlroot.attrib["name"]
        except KeyError, e:
            raise RuntimeError("invalid flow, missing attribute: %s" %
                               e.message)

        deps = {}

        for child in xmlroot:
            if child.tag != "node":
                continue
            node_name, node_res, node_deps, local_space = self.__parse_node(child)
            # format:
            #  {node_name : [node_object, deps_dict, depended_by_others, local_space]}
            deps[node_name] = [Node(node_name, node_res), node_deps, False, local_space]

        for name, dep in deps.iteritems():
            self.__parse_file(dep[0], dep[0].resource, dep[3])
            # parse dependence
            for dep_name in dep[1]:
                if deps[dep_name] == None:
                    raise RuntimeError("invalid dependence, "
                                       "missing node: %s" %
                                       dep_name)
                dep[0].depends.append(deps[dep_name][0])
                deps[dep_name][2] = True
                logger.debug("%s depends %s",
                             dep[0].name, deps[dep_name][0].name)

        for n, dep in deps.iteritems():
            if not dep[2]:
                node.depends.append(dep[0])

        return node

    def __parse_node(self, xmlroot):
        name = None
        res = None
        deps = {}
        local_space = {}
        for child in xmlroot:
            if child.tag == "name":
                name = child.text
            elif child.tag == "resource":
                res = child.text
            elif child.tag == "depend":
                deps[child.text.strip()] = 0
            elif child.tag == "property":
                pname, pvalue = self.__parse_property(child)
                local_space[pname] = pvalue
            else:
                raise RuntimeError("invalid property: %s" % child.tag)

        if name == None or res == None:
            raise RuntimeError("invalid node, missing resource")

        return name, res, deps, local_space

    def __parse_job(self, node, xmlroot, local_space):
        job = Job()
        resolved = {}
        for child in xmlroot:
            if child.tag != "property":
                continue
            name, value = self.__parse_property(child)
            if name == "hpipe.file":
                job.files.append(value)
                continue
            job.properties[name] = value
            resolved[name] = len(re.findall("\$\{(.+?)\}", value)) == 0

        # process grammer
        ref_resolver = ReferenceResolver()
        ref_resolver.add_space(os.environ)
        ref_resolver.add_space(local_space)
        ref_resolver.add_space(job.properties)
        ref_resolver.resolve()

        for k, v in job.properties.iteritems():
            orig = job.properties[k]
            job.properties[k] = ref_resolver.space[k]
            logger.debug("resolved $%s: %s -> %s", k, orig, job.properties[k])

        try:
            job.validate()
        except RuntimeError, e:
            raise RuntimeError("invalid job, missing property '%s'" %
                               e.message)

        node.job = job

    def __parse_property(self, xmlroot):
        name = None
        value = None
        for child in xmlroot:
            if child.tag == "name":
                name = child.text
            elif child.tag == "value":
                value = child.text
        if name == None or value == None:
            raise RuntimeError("invalid property, missing name or value")
        return name, value

    def __log_tree(self, node, level=1):
        str = " " * level * 2 + node.__repr__()
        logger.debug(str)
        for n in node.depends:
            self.__log_tree(n, level + 1)
