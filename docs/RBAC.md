---
id: rbac
title: RBAC
description: Casbin RBAC usage
keywords: [RBAC, role definition, role hierarchy]
authors: [hsluoyz]
---

## `[role_definition]`

* `g = <USER>, <ROLE>`
  * recommendations
    * ❌NOT reuse ❌
      * SAME name -- for -- user & role
        * PROPOSAL: use prefix
* == RBAC role relationships
  * inheritance
  * transitive
    * == if A has role B & B has role C -> A has role C
  * unbounded
* MANDATORY
* ALLOWED
  * 👀| same time, > 1 RBAC systems👀 
    * EACH one are independent
    * _Example of uses:_ 
      * 1 -- for -- users (with role inheritance) 
        * _Example:_ [here](https://github.com/casbin/casbin/blob/master/examples/rbac_model.conf)
      * 1 -- for -- resources
        * _Example:_ [here](https://github.com/casbin/casbin/blob/master/examples/rbac_with_resource_roles_model.conf)
* allows
  * if the request subject has the role | policy -> enable the request
* allows
  * store policy
  * evaluates user–role (resource–role) mappings
* ❌NOT validate❌
  * users OR roles exist
    * Reason:🧠that's part -- of the -- authentication🧠

* Casbin 
  * treats ALL names -- as -- strings (user, resource, or role)

TODO: 

:::info Token name convention
The subject in the policy is usually named `sub` and listed first
In Go Casbin you can use other names and order; 
then you must call `e.SetFieldIndex("p", constant.SubjectIndex, index)` after creating the enforcer 
so that APIs like `DeleteUser` use the correct column.

```ini
# `subject` here is for sub
[policy_definition]
p =  obj, act, subject
```

```go
e.SetFieldIndex("p", constant.SubjectIndex, 2) // index starts from 0
ok, err := e.DeleteUser("alice") // without SetFieldIndex, it will raise an error
```

## Role hierarchy

* == [RBAC1-style](CasbinRBACAndRBAC96.md) role hierarchy
  * _Example:_ if `alice` has `role1` & `role1` has `role2` -> `alice` has `role2` + its permissions

* **hierarchy depth** 
  * == levels of inheritance / you allow
  * _Example:_ [role manager](#role-manager)'s default hierarchy depth = 10

## Distinguishing users -- from -- roles

* users == string & roles == string

* | flat RBAC
  * == [NO role hierarchy](#role-hierarchy)
  * `GetAllSubjects()`
    * 's return: `g` rules' left side
  * `GetAllRoles()`
    * 's return: `g` rules' right side

* | role hierarchy, 
  * SAME name can appear | user & role; -> recommended steps
    * 👀use a naming convention 👀
      * _Example:_ prefix `role::`
    * | interpret results,
      * check it 

## Implicit roles and permissions

Roles and permissions inherited through the hierarchy (not only direct assignments) are called **implicit**
* Use `GetImplicitRolesForUser()` and `GetImplicitPermissionsForUser()` to include them; `GetRolesForUser()` and `GetPermissionsForUser()` return only direct assignments
* See [GitHub #137](https://github.com/casbin/casbin/issues/137).

## Pattern matching

* [here](RBACWithPattern.md)

## Role manager

* [here](RoleManagers.md)
