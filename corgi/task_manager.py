#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
#
# Copyright 2014 Menglong TAN <tanmenglong@gmail.com>
#

from xml.etree import ElementTree as ET
from entity import Flow, Node

import logging

logger = logging.getLogger("corgi")

class TaskManager(object):
    """Task manager"""

    def __init__(self):
        pass

    def load_conf(self, conf_file):
        root = ET.parse(conf_file).getroot()
        self.entry = parse_flow(root)

def parse_flow(root):
    flow = Flow();
    if root.tag != "flow":
        raise RuntimeError("invalid xml tag '%s'" % root.tag)
    try:
        flow.name = root.attrib["name"]
    except KeyError as e:
        raise RuntimeError("invalid flow, missing attribute: %s" % e.message)

    nodes = []
    deps = {}
    for child in root.iter("node"):
        node, node_deps = parse_node(child)
        print node.debug_string()
        nodes.append(node)
        deps[node.name] = node_deps
    for name, dep in deps.iteritems():
        print name, ":", dep
    return flow

def parse_node(root):
    node = Node()
    deps = []
    for child in root:
        if child.tag == "name":
            node.name = child.text
        elif child.tag == "resource":
            node.resource = child.text
        elif child.tag == "dependencies":
            for dep in child.text.split(","):
                deps.append(dep.strip())

    return node, deps
