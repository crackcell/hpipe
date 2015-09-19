/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file hive.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Tue Aug 25 18:28:05 2015
 *
 **/

package exec

import (
	"fmt"
	"github.com/crackcell/hpipe/config"
	"github.com/crackcell/hpipe/dag"
	"github.com/crackcell/hpipe/log"
	"github.com/crackcell/hpipe/util"
	"github.com/crackcell/hpipe/util/file"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type HiveExec struct {
}

func NewHiveExec() *HiveExec {
	return &HiveExec{}
}

func (this *HiveExec) Setup() error {
	return nil
}

func (this *HiveExec) Run(job *dag.Job) error {
	if !checkJobAttr(job, []string{"output"}) ||
		(!checkJobAttr(job, []string{"script"}) && !checkJobAttr(job, []string{"hql"})) {
		msg := "invalid job"
		log.Error(msg)
		return fmt.Errorf(msg)
	}

	// !!!VERY IMPORTANT!!!
	// Many other operations relay on this TrimRight.
	job.Attrs["output"] = strings.TrimRight(job.Attrs["output"], "/")

	args, err := this.genCmdArgs(job)
	if err != nil {
		log.Error(args)
		return err
	}
	log.Debugf("CMD: hive %s", strings.Join(args, " "))
	retcode, err := util.ExecCmd(job.Name, "hive", args...)
	if err != nil {
		job.Status = dag.Failed
		return err
	}
	if retcode != 0 {
		job.Status = dag.Failed
		return fmt.Errorf("script failed: %d", retcode)
	}
	job.Status = dag.Finished
	return nil
}

//===================================================================
// Private
//===================================================================

func (this *HiveExec) genCmdArgs(job *dag.Job) ([]string, error) {
	args := []string{}

	hql := ""

	// Process resource option, Add different types of files to hive
	if v, ok := job.Attrs["resource"]; ok {
		for _, res := range strings.Split(v, ";") {
			kv := strings.Split(res, ":")
			if len(kv) != 2 {
				return nil, fmt.Errorf("invalid resource: %s", res)
			}
			switch strings.Trim(kv[0], " ") {
			case "jar":
				hql += fmt.Sprintf("add jar %s;",
					config.WorkPath+"/"+strings.Trim(kv[1], " "))
			default:
				return nil, fmt.Errorf("invalid resource type: %s", kv[0])
			}
		}
	}

	// Create UDFs
	if v, ok := job.Attrs["udf"]; ok {
		for _, fun := range strings.Split(v, ";") {
			kv := strings.Split(fun, ":")
			if len(kv) != 2 {
				return nil, fmt.Errorf("invalid udf: %s", fun)
			}
			hql += fmt.Sprintf("create temporary function %s as '%s';",
				strings.Trim(kv[0], " "), strings.Trim(kv[1], " "))
		}
	}

	// Process hive options
	if v, ok := job.Attrs["option"]; ok {
		for _, option := range strings.Split(v, ";") {
			kv := strings.Split(option, "=")
			if len(kv) != 2 {
				return nil, fmt.Errorf("invalid option: %s", option)
			}
			hql += fmt.Sprintf("set %s=%s;",
				strings.Trim(kv[0], " "), strings.Trim(kv[1], " "))
		}
	}

	if v, ok := job.Attrs["hql"]; ok {
		hql += strings.Trim(v, ";") + ";"
	} else if v, ok := job.Attrs["script"]; ok {
		lines, err := file.ReadFileToLines(config.WorkPath + "/" + v)
		if err != nil {
			return nil, err
		}
		for _, line := range lines {
			hql += strings.Trim(line, " ") + " "
		}
	} else {
		return nil, fmt.Errorf("not hql or script for hive job: %s", job.Name)
	}
	if strings.HasSuffix(hql, ";") {
		hql += ";"
	}

	nhql, err := dag.ApplyVarToString(hql, job.Vars)
	if err != nil {
		return nil, err
	}

	args = append(args, "-e")
	args = append(args, "\""+nhql+"\"")

	return args, nil
}
