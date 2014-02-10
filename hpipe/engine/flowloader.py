#!/usr/bin/env python
# -*- encoding: utf-8; indent-tabs-mode: nil -*-
"""
    flowloader
    ~~~~~~~~~~

    Load flow tree from config file.

    :copyright: (c) 2014 Menglong TAN.
"""

import os
import re
import logging

from xml.etree import ElementTree as ET

from hpipe.util import setup_logger
from hpipe.engine.entity import Node, Job
from hpipe.engine.grammer import ReferenceResolver

logger = logging.getLogger(__name__)
setup_logger(logger)

class FlowLoader(object):
    """Flow Loader"""

    def __init__(self):
        self.input_pattern = re.compile("hpipe\..*?input.dir")

    def parse_file(self, node, conf_file,
                   global_space=os.environ, scope_space={}):
        logger.debug("load file: %s" % conf_file)
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
            refresolv.resolve()
            child_scope = refresolv.eval_prop(child_scope)

            self.parse_file(dep[0], dep[0].resource,
                            global_space, child_scope)
            # parse dependence
            for dep_name in dep[1]:
                if not dep_name in deps:
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
                job.files = value
                continue
            job.properties[name] = value

        refresolv = ReferenceResolver()
        refresolv.add_space(global_space)
        refresolv.add_space(scope_space)
        refresolv.add_space(job.properties)
        refresolv.resolve()
        job.properties = refresolv.eval_prop(job.properties)

        # separate input paths into array
        for k, v in job.properties.items():
            if self.input_pattern.match(k):
                for i in v.split(","):
                    i = i.strip()
                    if len(i) > 0:
                        job.inputs.append(i)

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
                value = self.__parse_property_value(child)
        if name == None or value == None:
            raise RuntimeError("invalid property, missing name or value")
        return name, value

    def __parse_property_value(self, xmlroot):
        for child in xmlroot:
            if child.tag == "array":
                return self.__parse_array(child)
            else:
                raise RuntimeError("invalid property value type: %s" %
                                   child.tag)
        return xmlroot.text

    def __parse_array(self, xmlroot):
        values = []
        if xmlroot.tag != "array":
            raise RuntimeError("invalid array type: %s" % xmlroot.tag)
        for child in xmlroot:
            if child.tag != "item":
                raise RuntimeError("invalid array tag: %s" % child.tag)
            values.append(child.text)
        return values
