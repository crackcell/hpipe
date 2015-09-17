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
 * @file webui.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Thu Sep 17 18:00:46 2015
 *
 **/

package webui

import (
	"github.com/astaxie/beego"
	_ "github.com/crackcell/hpipe/webui/routers"
)

//===================================================================
// Public APIs
//===================================================================

func Run(addr string) {
	beego.Run()
}

//===================================================================
// Private
//===================================================================
