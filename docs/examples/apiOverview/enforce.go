package apiOverview

import "github.com/casbin/casbin/v3"

// Load information from files.
enforcer, err := casbin.NewEnforcer("./example/model.conf", "./example/policy.csv")
if err != nil {
log.Fatalf("Error, detail: %s", err)
}
ok, err := enforcer.Enforce("alice", "data1", "read")