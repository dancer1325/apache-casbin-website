package modelFileWithDefaultFileAdapter

import "github.com/casbin/casbin/v3"

e, err := casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")


// check permissions
sub := "alice" 	// user / wants to access a resource
obj := "data1" // resource / is going to be accessed
act := "read" // operation / user performs | resource

ok, err := e.Enforce(sub, obj, act)

if err != nil {
// handle err
}

if ok == true {
// permit alice to read data1
} else {
// deny the request, show an error
}

// You could use BatchEnforce() to enforce some requests in batches.
// This method returns a bool slice, and this slice's index corresponds to the row index of the two-dimensional array.
// e.g. results[0] is the result of {"alice", "data1", "read"}
results, err := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})