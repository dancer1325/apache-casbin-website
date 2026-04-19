TODO:

# | ANY model CONF file, 
## `[request_definition], [policy_definition], [policy_effect], [matchers]`
* [basicSectionsPerEachConfFile.conf](basicSectionsPerEachConfFile.conf)
### if you use **RBAC** -> add `[role_definition]` section
* [rbac_model.conf](rbac_model.conf)
### if you use **RBAC with constraints** -> add `[constraint_definition]` section
* [rbac_with_constraints_model.conf](rbac_with_constraints_model.conf)

## `[request_definition]`
* [modelWithRequestDefinition.conf](modelWithRequestDefinition.conf) 
### parameters / passed -- to -- `e.Enforce(...)`
TODO: 

## `[policy_definition]`
* [modelWithPolicyDefinition.conf](modelWithPolicyDefinition.conf)
### | policy rules
#### 1 policy rule / EACH policy file's line
* [policyUsingPolicyDefinition.csv](policyUsingPolicyDefinition.csv)
  * `p, alice, data1, read`
    * (alice, data1, read) -> (p.sub, p.obj, p.act)
  * `p2, bob, write-all-objects`
    * (bob, write-all-objects) -> (p2.sub, p2.act)
#### 's first token == policy type / MUST match -- with a -- policy definition
* [policyUsingPolicyDefinition.csv](policyUsingPolicyDefinition.csv)
  * `p, alice, data1, read`
    * `p` == policy type
  * `p2, bob, write-all-objects`
    * `p2` == policy type
#### ' ALL elements are treated -- as -- strings
TODO: 

## `[policy_effect]`
### if MULTIPLE policies match (e.g. one allows, another denies) -> how to combine results
TODO:
### ALLOWED built-in policy effects
#### ❌CUSTOM effect expressions, are NOT supported❌
TODO:
#### `some(where (p.eft == allow))`
##### if a matched policy has effect `allow` -> the result is allow
TODO:
#### `!some(where (p.eft == deny))`
##### if NO matched policy has effect `deny` -> the result is allow
TODO:
#### `some(where (p.eft == allow)) && !some(where (p.eft == deny))`
##### requirements: >=1 `allow`& NO `deny`
TODO:
##### if BOTH match -> `deny` overrides
TODO:
#### `priority(p.eft) &#124;&#124; deny`
TODO:
#### `subjectPriority(p.eft)`
##### priority -- based on -- role
TODO:

## `[constraint_definition]`
TODO:

### 

## `[matchers]`
TODO:

### expression order | matchers
#### recommendations: place FIRST, cheaper matchers
* expensive FIRST,
  * | [expressionOrderInMatchers.go](expressionOrderInMatchers.go), use `rbac_models`
  * `go test -run ^TestManyRoles$ github.com/casbin/casbin/v3 -v`
    * take 6"
* cheap FIRST,
  * | [expressionOrderInMatchers.go](expressionOrderInMatchers.go), use `rbac_models_cheap_first`
  * `go test -run ^TestManyRoles$ github.com/casbin/casbin/v3 -v`
    * take << 1"

## TODO: