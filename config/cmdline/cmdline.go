/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 *
 *
 * @file cmdline.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sun Nov 23 20:59:30 2014
 *
 **/

package cmdline

import (
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

var (
	FlagHelp      bool
	FlagVerbose   bool
	FlagWorkRoot  string
	FlagEntryFile string
)

//===================================================================
// Private
//===================================================================
