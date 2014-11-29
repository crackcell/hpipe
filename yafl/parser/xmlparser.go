/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * XML parser
 *
 * @file xmlparser.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon Nov 10 15:42:26 2014
 *
 **/

package parser

import (
	"../../log"
	"../ast"
	"../tit"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

func NewXMLParser() Parser {
	return new(xmlParser)
}

//===================================================================
// Private
//===================================================================

type XMLStep struct {
	XMLName xml.Name `xml:"step"`
	Name    string   `xml:"name,attr"`
	Var     []string `xml:"var"`
	Dep     []XMLDep `xml:"dep"`
	Do      []XMLDo  `xml:"do"`
}

type XMLDep struct {
	XMLName xml.Name `xml:"dep"`
	Res     string   `xml:"res"`
	Var     []string `xml:"var"`
}

type XMLDo struct {
	XMLName xml.Name `xml:"do"`
	Res     string   `xml:"res"`
	Arg     []string `xml:"arg"`
}

type XMLJob struct {
	XMLName    xml.Name  `xml:"job"`
	Name       string    `xml:"name,attr"`
	InstanceID string    `xml:"instance_id"`
	Type       string    `xml:"type,attr"`
	Var        []string  `xml:"var"`
	Properties []XMLProp `xml:"property"`
}

type XMLProp struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type xmlParser struct{}

func (this *xmlParser) ParseStepFromFile(entry string,
	workdir string) (*ast.Step, error) {
	return parseStepFromFile(entry, workdir, nil)
}

func (this *xmlParser) ParseJobFromFile(entry string,
	workdir string) (*ast.Job, error) {
	return parseJobFromFile(entry, workdir, nil)
}

func parseStepFromFile(entry string, workdir string,
	preDefinedVars map[string]string) (*ast.Step, error) {

	path := workdir + "/" + entry

	//log.Debug("open:", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	s := XMLStep{}
	if err := xml.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	//log.Debug(s)

	step := ast.NewStep()
	step.Name = s.Name

	step.Var = arrayToMap(s.Var, "=")

	step.Var, err = evalMap(preDefinedVars, step.Var)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, do := range s.Do {
		localVar := arrayToMap(do.Arg, "=")
		localVar, err := evalMap(preDefinedVars, step.Var, localVar)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		j, err := parseJobFromFile(do.Res, workdir, localVar)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		step.Do = append(step.Do, j)
	}

	for _, dep := range s.Dep {
		localVar := arrayToMap(dep.Var, "=")
		localVar, err := evalMap(preDefinedVars, step.Var, localVar)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		j, err := parseStepFromFile(dep.Res, workdir, localVar)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		step.Dep = append(step.Dep, j)
	}

	return step, nil
}

func parseJobFromFile(entry string, workdir string,
	preDefinedVars map[string]string) (*ast.Job, error) {

	path := workdir + "/" + entry

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	j := XMLJob{}
	if err := xml.Unmarshal(data, &j); err != nil {
		log.Fatal(err)
		return nil, err
	}

	job := ast.NewJob()
	job.Type = j.Type
	job.Name = j.Name

	localVar := arrayToMap(j.Var, "=")
	localVar, err = evalMap(preDefinedVars, localVar)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, prop := range j.Properties {
		job.Prop[prop.Name] = prop.Value
	}
	newprop, err := applyMap(localVar, job.Prop)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	newinstid, err := applySrc(localVar, j.InstanceID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	job.InstanceID = newinstid
	job.Prop = newprop
	job.Var = localVar
	job.File = entry
	return job, nil
}

func updateMap(dest map[string]string, src map[string]string) {
	for k, v := range src {
		dest[k] = v
	}
}

func arrayToMap(array []string, sep string) map[string]string {
	m := make(map[string]string)
	updateMapWithArray(m, array, sep)
	return m
}

func updateMapWithArray(dest map[string]string, array []string, sep string) {
	for _, v := range array {
		tokens := strings.Split(v, sep)
		if len(tokens) != 2 {
			log.Fatalf("invalid var: %s", v)
			continue
		}
		dest[tokens[0]] = tokens[1]
	}
}

func evalMap(maps ...map[string]string) (map[string]string, error) {
	c := tit.NewTit()
	for _, m := range maps {
		if m != nil {
			c.AddVarMap(m)
		}
	}
	output, err := c.DoEval()
	if err != nil {
		return nil, err
	}
	return output, nil
}

var varPattern *regexp.Regexp = regexp.MustCompile(`\$\{([a-zA-Z0-9_]+)\}`)

func applySrc(vars map[string]string, src string) (string, error) {
	for _, g := range varPattern.FindAllSubmatch([]byte(src), -1) {
		raw_name := string(g[0])
		var_name := string(g[1])
		val, ok := vars[var_name]
		if !ok {
			errmsg := fmt.Sprintf("undefined var: %s", var_name)
			log.Error(errmsg)
			return "", fmt.Errorf(errmsg)
		}
		src = strings.Replace(src, raw_name, val, -1)
	}
	return src, nil
}

func applyMap(vars map[string]string,
	src map[string]string) (map[string]string, error) {

	for k, v := range src {
		v, err := applySrc(vars, v)
		if err != nil {
			return nil, err
		}
		src[k] = v
	}
	return src, nil
}
