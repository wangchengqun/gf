// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.
// 常用数据类型以及对象封装

package g

import "gitee.com/johng/gf/g/container/gvar"

// 框架动态变量，可以用该类型替代interface{}类型
type Var   = gvar.Var

// 常用map数据结构(使用别名)
type Map   = map[string]interface{}

// 常用list数据结构(使用别名)
type List  = []Map

// 常用slice数据结构(使用别名)
type Slice = []interface{}
type Array = Slice
