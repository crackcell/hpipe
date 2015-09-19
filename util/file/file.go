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
 * @file file.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Sep 19 17:31:36 2015
 *
 **/

package file

import (
	"bufio"
	"os"
)

//===================================================================
// Public APIs
//===================================================================

func ReadFileToLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

//===================================================================
// Private
//===================================================================
