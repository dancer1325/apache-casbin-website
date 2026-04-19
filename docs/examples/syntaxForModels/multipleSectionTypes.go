package syntaxForModels

// 1. set an `EnforceContext`
EnforceContext{"r2","p2","e2","m2"}
type EnforceContext struct {
	RType string
	PType string
	EType string
	MType string
}

// 2. `.enforce(previouslyDefinedEnforceContext)`
// Pass in a suffix as a parameter to NewEnforceContext, such as 2 or 3, and it will create r2, p2, etc.
enforceContext := NewEnforceContext("2")
// You can also specify a certain type individually
enforceContext.EType = "e"
// Don't pass in EnforceContext; the default is r, p, e, m
e.Enforce("alice", "data2", "read")        // true
// Pass in EnforceContext
e.Enforce(enforceContext, struct{ Age int }{Age: 70}, "/data1", "read")        //false
e.Enforce(enforceContext, struct{ Age int }{Age: 30}, "/data1", "read")        //true
