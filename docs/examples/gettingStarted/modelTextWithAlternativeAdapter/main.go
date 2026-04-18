package modelTextWithAlternativeAdapter

import (
"log"

"github.com/casbin/casbin/v3"
"github.com/casbin/casbin/v3/model"
xormadapter "github.com/casbin/xorm-adapter/v2"
_ "github.com/go-sql-driver/mysql"
)

// Initialize a Xorm adapter with MySQL database.
a, err := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
if err != nil {
log.Fatalf("error: adapter: %s", err)
}

m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
if err != nil {
log.Fatalf("error: model: %s", err)
}

e, err := casbin.NewEnforcer(m, a)
if err != nil {
log.Fatalf("error: enforcer: %s", err)
}



// check permissions
sub := "alice" // the user that wants to access a resource.
obj := "data1" // the resource that is going to be accessed.
act := "read" // the operation that the user performs on the resource.

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