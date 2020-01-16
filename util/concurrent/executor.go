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
   "container/list"
   "time"
)

type Runnable interface {
   Run()
}

type Executor interface {
   Execute(command Runnable)
}

type Callable interface {
   call() (interface{}, error)
}

type ExecutorService interface {
   Executor
   Shutdown()
   ShutdownNow() list.List
   IsShutdown() bool
   IsTerminated() bool
   AwaitTermination(timeout int64, duration time.Duration) bool
   SubmitCallable(callable Callable) (Future, error)
   SubmitRunnableWithResult(runnable Runnable, result interface{}) (Future, error)
   SubmitRunnable(runnable Runnable) (Future, error)
   InvokeAll(tasks list.List) (list.List, error)
   InvokeAllWithTimeout(tasks list.List, timeout time.Duration) (list.List, error)
   InvokeAny(tasks list.List,) (interface{}, error)
   InvokeAnyWithTimeout(tasks list.List, timeout time.Duration) (interface{}, error)
}

type ScheduledExecutorService interface {
   ExecutorService
   ScheduleRunnable(command Runnable, delay time.Duration) (ScheduledFuture, error)
   ScheduleCallable(callable Callable, delay time.Duration) (ScheduledFuture, error)
   ScheduleAtFixedRate(command Runnable, initDelay time.Duration, period time.Duration) (ScheduledFuture, error)
   ScheduleWithFixedDelay(command Runnable, initDelay time.Duration, delay time.Duration) (ScheduledFuture, error)
}

type EventExecutorGroup interface {
   ScheduledExecutorService
   IsShuttingDown() bool
   ShutdownGracefully() (Future, error)
   ShutdownGracefullyWithTimeout(quietPeriod int64, timeout time.Duration) (Future, error)
   TerminationFuture() (Future, error)
   Next() (EventExecutor, error)
}

type EventExecutor interface {
   EventExecutorGroup
   InEventLoop() bool

}