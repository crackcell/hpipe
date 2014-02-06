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
        self.__parse_file(self.root, conf_file, os.environ, {})
        logger.debug("node tree:")
        self.__log_tree(self.root)

    def launch(self):
        logger.info("launching jobs")
        # TODO
        #self.launcher.launch(self.root)

    def __parse_file(self, node, conf_file, global_space, scope_space):
        xmlroot = ET.parse(conf_file).getroot()
        if xmlroot.tag == "flow":
            self.__parse_flow(node, xmlroot, global_space, scope_space)
        elif xmlroot.tag == "job":
            self.__parse_job(node, xmlroot, global_space, scope_space)
        else:
            raise RuntimeError("invalid xml tag '%s'" % xmlroot.tag)

    def __parse_flow(self, node, xmlroot, global_space, scope_space):
        try:
            node.name = xmlroot.attrib["name"]
        except KeyError, e:
            raise RuntimeError("invalid flow, missing attribute: %s" %
                               e.message)
        deps = {}
        for child in xmlroot:
            if child.tag != "node":
                continue
            node_name, node_res, node_deps, node_prop = self.__parse_node(child)
            # format:
            #  {node_name : [node_object, deps_dict, depended_by_others]}
            deps[node_name] = [Node(node_name, node_res, node_prop),
                               node_deps, False]

        for name, dep in deps.iteritems():
            child_scope = {}
            child_scope.update(scope_space)
            child_scope.update(dep[0].properties)

            refresolv = ReferenceResolver()
            refresolv.add_space(global_space)
            refresolv.add_space(scope_space)
            refresolv.add_space(child_scope)
            child_scope = refresolv.resolve_properties(child_scope)

            self.__parse_file(dep[0], dep[0].resource,
                              global_space, child_scope)
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
        properties = {}
        for child in xmlroot:
            if child.tag == "name":
                name = child.text
            elif child.tag == "resource":
                res = child.text
            elif child.tag == "depend":
                deps[child.text.strip()] = 0
            elif child.tag == "property":
                pname, pvalue = self.__parse_property(child)
                properties[pname] = pvalue
            else:
                raise RuntimeError("invalid property: %s" % child.tag)

        if name == None or res == None:
            raise RuntimeError("invalid node, missing resource")

        return name, res, deps, properties

    def __parse_job(self, node, xmlroot, global_space, scope_space):
        job = Job()
        for child in xmlroot:
            if child.tag != "property":
                continue
            name, value = self.__parse_property(child)
            if name == "hpipe.file":
                job.files.append(value)
                continue
            job.properties[name] = value

        refresolv = ReferenceResolver()
        refresolv.add_space(global_space)
        refresolv.add_space(scope_space)
        refresolv.add_space(job.properties)
        job.properties = refresolv.resolve_properties(job.properties)

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

    def __log_tree(self, node, level=0):
        str = node.__repr__()
        lines = str.split("\n")
        for i in range(len(lines)):
            prefix = "| "
            if i == 0:
                prefix = "+ "
            logger.debug("|    " * level + prefix + lines[i])
        for n in node.depends:
            self.__log_tree(n, level + 1)
