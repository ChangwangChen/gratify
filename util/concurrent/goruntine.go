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
   "context"
   "reflect"
   "sync/atomic"
)

var goRuntineId int64 = 0

type GoRuntineStatus int8

const (
   GORUNTINE_CREATED GoRuntineStatus = iota
   GORUNTINE_RUNNING
   GORUNTINE_FINISHED
   GORUNTINE_TERMINATED
)

type GoRuntine struct {
   id int64
   targetFunc *reflect.Value
   invokeArgs *[]reflect.Value
   done chan interface{}
   status GoRuntineStatus
   ctx context.Context
   cancel context.CancelFunc
}

func NewGoRuntine(targetFunc interface{}, args ...interface{}) *GoRuntine {
   funcValue := reflect.ValueOf(targetFunc)
   argValues := make([]reflect.Value, 0, len(args) + 1)
   ctx, cancel := context.WithCancel(context.Background())
   argValues = append(argValues, reflect.ValueOf(ctx))
   for i, a := range args {
      argValues[i + 1] = reflect.ValueOf(a)
   }
   return &GoRuntine{
      id: atomic.AddInt64(&goRuntineId, 1),
      targetFunc: &funcValue,
      invokeArgs: &argValues,
      done: make(chan interface{}),
      status: GORUNTINE_CREATED,
      ctx: ctx,
      cancel: cancel,
   }
}

func (self *GoRuntine) Run() {
   go func(self *GoRuntine) {
      defer func() {
         close(self.done)
         if self.ctx.Err() != nil {
            self.status = GORUNTINE_TERMINATED
         } else {
            self.status = GORUNTINE_FINISHED
         }
      }()
      self.status = GORUNTINE_RUNNING
      self.targetFunc.Call(*self.invokeArgs)
   }(self)
}

func (self *GoRuntine) Wait() {
   <-self.done
}

func (self *GoRuntine) GetId() int64 {
   return self.id
}

func (self *GoRuntine) GetStatus() GoRuntineStatus {
   return self.status
}

func (self *GoRuntine) TryTerminate() {
   self.cancel()
}