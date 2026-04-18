# access control
## == how users & resources interact
* [acl.conf](acl.conf) & [policyForACL.csv](policyForACL.csv)
## expressed -- as -- CONF files
* [acl.conf](acl.conf)
### -- based on the -- ⭐️ **PERM metamodel (Policy, Effect, Request, Matchers)** ⭐️
* [acl.conf](acl.conf)

# Request -- `r` --
## == access request's parameters
TODO:
## == 💡**subject** (who is requesting) + **object** (the resource) + **action** (the operation)💡
* [acl.conf](acl.conf)'s `[request_definition]`
### ⚠️EXACTLY this order ⚠️
TODO:

# Policy -- `p` --
## == access rules' shape
* [acl.conf](acl.conf)'s `[policy_definition]`
### == field names + order
* [acl.conf](acl.conf)'s `[policy_definition]`
#### if `eft` (policy effect) is omitted -> matching policies -- are treated as -- "allow"
TODO:


# TODO: