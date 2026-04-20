---
id: rbac-96
title: RBAC vs. RBAC96
description: The Difference Between Casbin RBAC and RBAC96
keywords: [RBAC96, Casbin RBAC]
authors: [kininaru]
---

* goal
  * Casbin RBAC
  * Casbin [RBAC96](https://profsandhu.com/cs6393_s12/lecture-rbac96.pdf)

| RBAC96    | Casbin support  | Notes                                                                                                                             |
|-----------|-----------------|-----------------------------------------------------------------------------------------------------------------------------------|
| **RBAC0** | ✅ Full          | Users, roles, permissions and their relationships.                                                                                |
| **RBAC1** | ✅ Full          | Role hierarchy (e.g. alice→role1→role2)                                                                                           |
| **RBAC2** | ⚠️ Partial      | Constraint-style handling (_Example:_ deny-override) -- via -- [policy effect](SyntaxForModels.md) <br/> ❌NO quantitative limits❌ |
| **RBAC3** | ⚠️ Partial      | RBAC1 + RBAC2 <br/> Constraint-style handling (_Example:_ deny-override) -- via -- [policy effect](SyntaxForModels.md) <br/>  ❌NO quantitative limits❌                                                          |

## RBAC vs RBAC96

### 1. User vs. role
TODO: 
Casbin does not separate “user” and “role” in the type system—both are strings. So:

   ```csv
   p, admin, book, read
   p, alice, book, read
   g, amber, admin
   ```

   `GetAllSubjects()` returns both users and roles that appear as subjects in policy (`[admin alice]`)
* `GetAllRoles()` returns only the right-hand side of `g` rules (`[admin]`)
* So Casbin infers the distinction from usage; for strict separation use a naming convention (e.g
* `user::alice`, `role::admin`).

### 2. Permissions

RBAC96 fixes seven permission types
* In Casbin, permissions are arbitrary strings (e.g. `read`, `write`, `approve`), so you can match your app.

### 3. Domains

Casbin supports [RBAC with domains](/docs/rbac-with-domains), so roles and permissions can be scoped per tenant or domain—beyond the standard RBAC96 model.
