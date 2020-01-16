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

type GenericFutureListener interface {
   // Invoked when the operation associated with the Future has been completed.
   OperationComplete(future Future) error
}