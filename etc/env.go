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
 * @file env.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Dec  3 20:28:25 2014
 *
 **/

package etc

import (
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

var Env map[string]string = make(map[string]string)

//===================================================================
// Private
//===================================================================
