#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

from xml.etree import ElementTree as ET

from entity import Node, JobConf
from envrion import Envrion

import logging
import re

logger = logging.getLogger("corgi")

class JobManager(object):
    """Job Manager"""

    def __init__(self):
        self.root = Node()

    def load_conf(self, conf_file):
        parse_file(self.root, conf_file)
        logger.debug("node tree:")
        print_tree(self.root)

    def launch(self):
        logger.info("launching jobs")

def parse_file(node, conf_file):
    xmlroot = ET.parse(conf_file).getroot()

    if xmlroot.tag == "flow":
        parse_flow(node, xmlroot)
    elif xmlroot.tag == "jobconf":
        parse_jobconf(node, xmlroot)
    else:
        raise RuntimeError("invalid xml tag '%s'" % xmlroot.tag)

def parse_flow(node, xmlroot):

    try:
        node.name = xmlroot.attrib["name"]
    except KeyError as e:
        raise RuntimeError("invalid flow, missing attribute: %s" %
                           e.message)

    deps = {}
    for child in xmlroot.iter("node"):
        node_name, node_res, node_deps = parse_node_info(child)
        #{node_name : [node_object, deps_dict, depended_by_others]}
        deps[node_name] = [Node(node_name, node_res), node_deps, False]

    for name, dep in deps.iteritems():
        parse_file(dep[0], dep[0].resource)
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

def parse_node_info(xmlroot):
    name = None
    res = None
    deps = {}
    for child in xmlroot:
        if child.tag == "name":
            name = child.text
        elif child.tag == "resource":
            res = child.text
        elif child.tag == "dependencies":
            for dep in child.text.split(","):
                deps[(dep.strip())] = 0

    if name == None or res == None:
        raise RuntimeError("invalid node, missing resource")

    return name, res, deps

def parse_jobconf(node, xmlroot):
    jobconf = JobConf()

    resolved = {}
    for child in xmlroot.iter("property"):
        name, value = parse_property_info(child)
        jobconf.properties[name] = value
        resolved[name] = len(re.findall("\$\{(.+?)\}", value)) == 0

    for name, value in jobconf.properties.iteritems():
        unresolved = re.findall("\$\{(.+?)\}", value)
        for var in re.findall("\$\{(.+?)\}", value):
            if not var in jobconf.properties or not resolved[var]:
                raise RuntimeError("invalid property, "
                                   "undefined or unresolved "
                                   "variable: ${%s}" %
                                   var)
            orig = jobconf.properties[name]
            jobconf.properties[name] = jobconf.properties[name].replace(
                "${" + var + "}",
                jobconf.properties[var])
            logger.debug("resolving: %s -> %s",
                         orig, jobconf.properties[name])
    try:
        jobconf.validate()
    except RuntimeError as e:
        raise RuntimeError("invalid jobconf, missing property '%s'" %
                           e.message)
    node.jobconf = jobconf

def parse_property_info(xmlroot):
    name = None;
    value = None;
    for child in xmlroot:
        if child.tag == "name":
            name = child.text
        elif child.tag == "value":
            value = child.text
    if name == None or value == None:
        raise RuntimeError("invalid property, missing name or value")
    return name, value

def print_tree(node, level=1):
    str = " " * level * 2 + node.__repr__()
    logger.debug(str)
    for n in node.depends:
        print_tree(n, level + 1)
