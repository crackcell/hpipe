/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Evaluate variables in ast
 *
 * @file eval.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Nov 14 16:36:06 2014
 *
 **/

package flow

import (
	_ "fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Evaluator interface {
	DoEval(root *Step) error
}

func NewTopDownEvaluator() Evaluator {
	return new(topDownEvaluator)
}

//===================================================================
// Private
//===================================================================

type topDownEvaluator struct{}

func (this *topDownEvaluator) DoEval(root *Step) error {

	return nil
}
