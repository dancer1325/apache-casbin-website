package fileAdapter

import (
	"github.com/casbin/casbin/v3"
	fileadapter "github.com/casbin/casbin/v3/persist/file-adapter"
)

// way1
// 		.NewEnforcer()		==	load the policies AUTOMATICALLY
e := casbin.NewEnforcer("examples/basic_model.conf", "examples/basic_policy.csv")

// way2
a := fileadapter.NewAdapter("examples/basic_policy.csv")
e := casbin.NewEnforcer("examples/basic_model.conf", a)