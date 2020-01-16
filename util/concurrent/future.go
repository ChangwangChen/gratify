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
   "time"
)

type Value interface {}

type Future interface {
   Cancel(mayInterruptIfRunning bool) bool
   IsCancelled() bool
   IsDone() bool
   Get() (Value, error)
   GetWithTimeout(timeout time.Duration) (Value, error)
   GetWithDeadline(deadline time.Time) (Value, error)
   GetWithContext(context context.Context) (Value, error)
   // Returns true if and only if the I/O operation was completed
   // successfully.
   IsSuccess() bool
   // returns true if and only if the operation can be cancelled via Cancel(boolean).
   IsCancellable() bool
   Cause() []byte
   AddListener(listener GenericFutureListener) (Future, error)
   AddListeners(listeners ...GenericFutureListener) (Future, error)
   RemoveListener(listener GenericFutureListener) (Future, error)
   RemoveListeners(listeners ...GenericFutureListener) (Future, error)
   Sync() (Future, error)
   SyncUninterruptibly() (Future, error)
   Await() (Future, error)
   AwaitWithTimeout(timeout time.Duration) (Future, error)
   AwaitWithDeadline(deadline time.Time) (Future, error)
   AwaitWithContext(context context.Context) (Future, error)

   AwaitUninterruptibly() (Future, error)
   AwaitUninterruptiblyWithTimeout(timeout time.Duration) (Future, error)
   AwaitUninterruptiblyWithDeadline(deadline time.Time) (Future, error)
   AwaitUninterruptiblyWithContext(context context.Context) (Future, error)

   GetNow() (Value, error)
   //Executor() (EventExecutor, error)
}

type ScheduledFuture interface {
   Future
   getDelay(uint time.Duration) int64
}

type AbstractFuture struct {}

//func (future *AbstractFuture) Get() (Value, error) {
//
//}
//
//func (future *AbstractFuture) GetWithTimeout(timeout time.Duration) (Value, error) {
//
//}
//
//func (future *AbstractFuture) GetWithDeadline(deadline time.Time) (Value, error) {
//
//}
//
//func (future *AbstractFuture) GetWithContext(context context.Context) (Value, error) {
//
//}