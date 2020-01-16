// This source file is part of the gratify open source project
//
// Copyright (c) 2017 - 2020 polar software foundation
// Licensed under Apache License v2.0 with Runtime Library Exception
//
// See https://polar.foundation/LICENSE.txt for license information
// See https://polar.foundation/CONTRIBUTORS.txt for the list of polarphp project authors
//
// Created by polarboy on 2020/01/16.

package concurrent

import (
   "reflect"
   "sync/atomic"
)

var goRuntineId int64 = 0

type GoRuntine struct {
   id int64
   targetFunc *reflect.Value
   invokeArgs *[]reflect.Value
}

func NewGoRuntine(targetFunc interface{}, args ...interface{}) *GoRuntine {
   funcValue := reflect.ValueOf(targetFunc)
   argValues := make([]reflect.Value, len(args))
   for i, a := range args {
      argValues[i] = reflect.ValueOf(a)
   }
   return &GoRuntine{
      id: atomic.AddInt64(&goRuntineId, 1),
      targetFunc: &funcValue,
      invokeArgs: &argValues,
   }
}

func (self *GoRuntine) Run() {
   go func() {
      self.targetFunc.Call(*self.invokeArgs)
   }()
}

func (self *GoRuntine) GetId() int64 {
   return self.id
}