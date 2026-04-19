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

Casbin implements RBAC1-style role hierarchy: if `alice` has `role1` and `role1` has `role2`, 
then `alice` effectively has `role2` and its permissions.

The **hierarchy depth** is how many levels of inheritance you allow
The default role manager uses a maximum depth of 10 (configurable),
so a user can inherit up to 10 levels of roles.

## Distinguishing users from roles

In Casbin, users and roles are both strings
* In flat RBAC (no role hierarchy), `GetAllSubjects()` and `GetAllRoles()` return the left and right sides of `g` rules (often users and roles)
* With a hierarchy, the same name can appear as both user and role; if your app does not track which is which, use a naming convention (e.g
* prefix `role::`) and check it when interpreting results.

## Implicit roles and permissions

Roles and permissions inherited through the hierarchy (not only direct assignments) are called **implicit**
* Use `GetImplicitRolesForUser()` and `GetImplicitPermissionsForUser()` to include them; `GetRolesForUser()` and `GetPermissionsForUser()` return only direct assignments
* See [GitHub #137](https://github.com/casbin/casbin/issues/137).

## Pattern matching

* [here](RBACWithPattern.md)

## Role manager

* [here](RoleManagers.md)
