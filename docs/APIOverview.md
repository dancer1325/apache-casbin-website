---
id: api-overview
title: API Overview
description: Casbin API Usage
keywords: [API overview, API, API usage]
authors: [hsluoyz]
---

* goal
  * Casbin’s main APIs

## Enforce API

* `NewEnforcer("pathToModel","pathToPolicy", yourFileAdapterOrIfYouOmitTheDefaultFileAdapter)`
  * loads the 
    * model
    * policy
  * 's returns
    * enforcer
    * error

* `enforcer.Enforce("<SUBJECT>", "<OBJECT>", "<ACTION>")`
  * 's return
    * `(true, nil)`
      * == allowed, OR
    * `(false, nil)` 
      * == denied

## EnforceEx API

* `EnforceEx()`
  * allows
    * check the policy / allowed a request

* related APIs
  * `ok, err := enforcer.EnforceWithMatcher(matcher, request)`
    * == enforce -- with a -- 1-off matcher expression
  * `EnforceExWithMatcher(matcher, request)` 
    * == `EnforceEx` + custom matcher
  * `BatchEnforce(requests)`
    * == enforce MANY requests at once
    * 's return
      * a slice of booleans

* MORE |
  * [Management API](ManagementAPI.md)
  * [RBAC API](RBACAPI.md)

## Management API

### Get API

* 's return
  * policy data

* `GetAllSubjects()`
  * 's return
    * ALL subjects / appear | policy

* `GetAllNamedSubjects("p")`

* `GetPolicy()` / `GetFilteredPolicy(0, "alice")`
  * 's return
    * ALL policy rules / policy filtered rules by field
* `GetNamedPolicy("p")` / `GetFilteredNamedPolicy("p", 0, "bob")`
  * 's return
    * ALL named policy rules / named filtered policy by field
* `GetGroupingPolicy()` / `GetFilteredGroupingPolicy(0, "alice")`
  * 's return
    * ALL grouping policy rules / grouping filtered policy by field
* `GetNamedGroupingPolicy("g")` / `GetFilteredNamedGroupingPolicy("g", 0, "alice")`
  * 's return
    * ALL named grouping policy rules / named grouping filtered policy by field

### Add, delete, update API

You can change policy at runtime
* Example: add, remove, update, and check:

```go
// load information from files
enforcer, err := casbin.NewEnforcer("./example/model.conf", "./example/policy.csv")
if err != nil {
    fmt.Printf("Error, details: %s\n", err)
}

// add a policy and use HasPolicy() to confirm
enforcer.AddPolicy("added_user", "data1", "read")
hasPolicy := enforcer.HasPolicy("added_user", "data1", "read")
fmt.Println(hasPolicy) // true, the policy was added successfully

// remove a policy and use HasPolicy() to confirm
enforcer.RemovePolicy("alice", "data1", "read")
hasPolicy = enforcer.HasPolicy("alice", "data1", "read")
fmt.Println(hasPolicy) // false, the policy was removed successfully

// update a policy and use HasPolicy() to confirm
enforcer.UpdatePolicy([]string{"added_user", "data1", "read"}, []string{"added_user", "data1", "write"})
hasPolicy = enforcer.HasPolicy("added_user", "data1", "read")
fmt.Println(hasPolicy) // false, the original policy has expired
hasPolicy = enforcer.HasPolicy("added_user", "data1", "write")
fmt.Println(hasPolicy) // true, the new policy is in effect
```

The same add/remove/update/has pattern exists for filtered and named policies and for grouping policies
* For batch updates, use the plural forms (e.g
* `AddPolicies`, `UpdatePolicies`) with slice arguments.

Example:

```go
enforcer.UpdatePolicy([]string{"eve", "data3", "read"}, []string{"eve", "data3", "write"})
```

Change `Policy` to `Policies` and modify parameters for batch operations:

```go
enforcer.UpdatePolicies([][]string{{"eve", "data3", "read"}, {"jack", "data3", "read"}}, [][]string{{"eve", "data3", "write"}, {"jack", "data3", "write"}})
```

Batch operations also apply to `GroupingPolicy` and `NamedGroupingPolicy`.

### AddEx API

The AddEx methods add rules in batch but **skip** rules that already exist instead of failing:

```go
AddPoliciesEx(rules [][]string) (bool, error)
AddNamedPoliciesEx(ptype string, rules [][]string) (bool, error)
AddGroupingPoliciesEx(rules [][]string) (bool, error)
AddNamedGroupingPoliciesEx(ptype string, rules [][]string) (bool, error)
SelfAddPoliciesEx(sec string, ptype string, rules [][]string) (bool, error) 
```

With `AddPolicies`, if one rule already exists the whole call fails and no rules are added
* With `AddPoliciesEx`, existing rules are skipped and the rest are added
* Example:

```go
func TestDemo(t *testing.T) {
    e, err := NewEnforcer("examples/basic_model.conf", "examples/basic_policy.csv")
    if err != nil {
        fmt.Printf("Error, details: %s\n", err)
    }
    e.ClearPolicy()
    e.AddPolicy("user1", "data1", "read")
    fmt.Println(e.GetPolicy())
    testGetPolicy(t, e, [][]string{{"user1", "data1", "read"}})

    // policy {"user1", "data1", "read"} now exists

    // Use AddPolicies to add rules in batches
    ok, _ := e.AddPolicies([][]string{{"user1", "data1", "read"}, {"user2", "data2", "read"}})
    fmt.Println(e.GetPolicy())
    // {"user2", "data2", "read"} failed to add because {"user1", "data1", "read"} already exists
    // AddPolicies returns false and no other policies are checked, even though they may not exist in the existing ruleset
    // ok == false
    fmt.Println(ok)
    testGetPolicy(t, e, [][]string{{"user1", "data1", "read"}})

    // Use AddPoliciesEx to add rules in batches
    ok, _ = e.AddPoliciesEx([][]string{{"user1", "data1", "read"}, {"user2", "data2", "read"}})
    fmt.Println(e.GetPolicy())
    // {"user2", "data2", "read"} is added successfully
    // because AddPoliciesEx automatically filters the existing {"user1", "data1", "read"}
    // ok == true
    fmt.Println(ok)
    testGetPolicy(t, e, [][]string{{"user1", "data1", "read"}, {"user2", "data2", "read"}})
}
```

## RBAC API

The enforcer also exposes RBAC helpers (roles, users, permissions)
* For the model setup, see [RBAC](/docs/rbac).

Load model and policy as before:

```go
enforcer, err := casbin.NewEnforcer("./example/model.conf", "./example/policy.csv")
if err != nil {
    fmt.Printf("Error, details: %s\n", err)
}
```

Example RBAC calls:

```go
roles, err := enforcer.GetRolesForUser("amber")
fmt.Println(roles) // [admin]
users, err := enforcer.GetUsersForRole("admin")
fmt.Println(users) // [amber abc]
```

`GetRolesForUser("amber")` returns `[admin]`
* `GetUsersForRole("admin")` returns all users with that role
* `HasRoleForUser("amber", "admin")` is `true`.

```go
fmt.Println(enforcer.Enforce("bob", "data2", "write")) // true
enforcer.DeletePermission("data2", "write")
fmt.Println(enforcer.Enforce("bob", "data2", "write")) // false
```

`DeletePermission("data2", "write")` removes that permission for everyone
* `DeletePermissionForUser("alice", "data1", "read")` removes it only for Alice.

For the full list of RBAC APIs, see [RBAC API](/docs/rbac-api).
