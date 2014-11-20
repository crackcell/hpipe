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

package flow

import (
	log "../../levellog"
	"../calc"
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	XMLName xml.Name `xml:"job"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Var     []string `xml:"var"`
}

type xmlParser struct{}

func (this *xmlParser) ParseStepFromFile(entry string,
	workdir string) (*Step, error) {
	return parseStepFromFile(entry, workdir, nil)
}

func (this *xmlParser) ParseJobFromFile(entry string,
	workdir string) (Job, error) {
	return parseJobFromFile(entry, workdir, nil)
}

func parseStepFromFile(entry string, workdir string,
	preDefinedVars map[string]string) (*Step, error) {

	entry = workdir + "/" + entry

	//log.Debug("open:", entry)
	data, err := ioutil.ReadFile(entry)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	s := XMLStep{}
	if err := xml.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	//log.Debug(s)

	step := NewStep()
	step.Name = s.Name

	step.Var = arrayToMap(s.Var, "=")

	step.Var, err = evalMap(preDefinedVars, step.Var)
	if err != nil {
		return nil, err
	}

	for _, do := range s.Do {
		localVar := arrayToMap(do.Arg, "=")
		//log.Debugf("%s ============\n", entry)
		//log.Debugf("predef: %v\n", preDefinedVars)
		//log.Debugf("step.Var: %v\n", step.Var)
		//log.Debugf("localVar: %v\n", localVar)
		localVar, err := evalMap(preDefinedVars, step.Var, localVar)
		if err != nil {
			return nil, err
		}
		//log.Debugf("output: %v\n", localVar)
		j, err := parseJobFromFile(do.Res, workdir, localVar)
		if err != nil {
			return nil, err
		}
		step.Do = append(step.Do, j)
	}

	for _, dep := range s.Dep {
		localVar := arrayToMap(dep.Var, "=")
		//log.Debugf("%s ============\n", entry)
		//log.Debugf("predef: %v\n", preDefinedVars)
		//log.Debugf("step.Var: %v\n", step.Var)
		//log.Debugf("localVar: %v\n", localVar)
		localVar, err := evalMap(preDefinedVars, step.Var, localVar)
		if err != nil {
			panic(err)
		}
		//log.Debugf("output: %v\n", localVar)
		j, err := parseStepFromFile(dep.Res, workdir, localVar)
		if err != nil {
			return nil, err
		}
		step.Dep = append(step.Dep, j)
	}

	return step, nil
}

func parseJobFromFile(entry string, workdir string,
	preDefinedVars map[string]string) (Job, error) {

	entry = workdir + "/" + entry

	//log.Debug("open:", entry)
	data, err := ioutil.ReadFile(entry)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	j := XMLJob{}
	if err := xml.Unmarshal(data, &j); err != nil {
		return nil, err
	}

	//log.Debug(j)

	var job Job

	switch j.Type {
	case "odps":
		job = NewODPSJob()
	case "hadoop":
		job = NewHadoopJob()
	default:
		return nil, fmt.Errorf("unknown job type")
	}
	job.SetName(j.Name)

	localVar := arrayToMap(j.Var, "=")
	//log.Debugf("%s ============\n", entry)
	//log.Debugf("predef: %v\n", preDefinedVars)
	//log.Debugf("evalMap: %v\n", localVar)
	localVar, err = evalMap(preDefinedVars, localVar)
	if err != nil {
		return nil, err
	}
	//log.Debugf("output: %v\n", localVar)
	job.SetVar(localVar)
	if !job.IsValid() {
		return nil, ErrInvalidJob
	}

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
	c := calc.NewCalc()
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
