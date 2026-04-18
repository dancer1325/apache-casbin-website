---
id: supported-models
title: Supported Models
description: Supported models of Casbin
keywords: [supported models]
authors: [nodece]
---

* supported access control models & patterns

1. [**ACL (Access Control List)**](https://en.wikipedia.org/wiki/Access_control_list)
2. **ACL + [superuser](https://en.wikipedia.org/wiki/Superuser)**
3. **ACL WITHOUT users**
   * _Example:_ API-only OR device access
4. **ACL WITHOUT resources**
   * == permissions apply 
     * | resource types
     * NOT | specific instances
       * _Examples:_ "write-article", "read-log"
5. **[RBAC (Role-Based Access Control)](https://en.wikipedia.org/wiki/Role-based_access_control)**
6. **RBAC + resource roles**
   * users & resources can have 
     * roles, OR
     * group memberships
7. **RBAC + domains/tenants**
   * users can have DIFFERENT roles | 
     * DIFFERENT domains OR
     * DIFFERENT tenants
8. **[ABAC (Attribute-Based Access Control)](https://en.wikipedia.org/wiki/Attribute-Based_Access_Control)**
   * uses attributes | rules
     * _Example:_  `resource.Owner`)
9. **[PBAC (Policy-Based Access Control)](https://heimdalsecurity.com/blog/rbac-vs-abac-vs-pbac/#:~:text=ABAC%20has%20many%20benefits%20over,no%20predefined%20roles%20or%20permissions)**
   * == Authorization is driven -- by -- rule-based policies
   * uses
     * dynamic decisions
     * context-aware decisions
10. **[BLP (Bell–LaPadula)](https://en.wikipedia.org/wiki/Bell%E2%80%93LaPadula_model)** 
    * == Formal model + (security labels & clearances)
11. **[Biba](https://en.wikipedia.org/wiki/Biba_Model)** 
    * restricts information flow
      * Reason:🧠prevent unauthorized changes🧠
12. **[LBAC (Lattice-Based Access Control)](/docs/lbac)**
    * == combine confidentiality + integrity | lattice-based framework
13. **[OrBAC (Organisation-Based Access Control)](/docs/orbac)**
    * == RBAC extended /
      * abstraction layers -- for -- MULTI-organization policies
14. **[UCON (Usage Control)](/docs/ucon)** 
    * enable
      * ongoing authorization
      * mutable attributes
      * obligations
      * conditions
15. **RESTful**
    * == Path patterns (e.g. `/res/*`, `/res/:id`) + HTTP methods
16. **IP Match** 
    * == match -- by -- IP address OR CIDR for network-level control
17. **Deny-override**
    * == if you have allow policy + deny policies -> ALWAYS deny policy overrides
18. **Priority**
    * == order the policy rules /
      * first match wins

## Examples

| Model                     | Model file                                                                                                                       | Policy file                                                                                                                      |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------|
| ACL                       | [basic_model.conf](https://github.com/casbin/casbin/blob/master/examples/basic_model.conf)                                       | [basic_policy.csv](https://github.com/casbin/casbin/blob/master/examples/basic_policy.csv)                                       |
| ACL with superuser        | [basic_with_root_model.conf](https://github.com/casbin/casbin/blob/master/examples/basic_with_root_model.conf)                   | [basic_policy.csv](https://github.com/casbin/casbin/blob/master/examples/basic_policy.csv)                                       |
| ACL WITHOUT users         | [basic_WITHOUT_users_model.conf](https://github.com/casbin/casbin/blob/master/examples/basic_without_users_model.conf)           | [basic_WITHOUT_users_policy.csv](https://github.com/casbin/casbin/blob/master/examples/basic_without_users_policy.csv)           |
| ACL WITHOUT resources     | [basic_WITHOUT_resources_model.conf](https://github.com/casbin/casbin/blob/master/examples/basic_without_resources_model.conf)   | [basic_without_resources_policy.csv](https://github.com/casbin/casbin/blob/master/examples/basic_without_resources_policy.csv)   |
| RBAC                      | [rbac_model.conf](https://github.com/casbin/casbin/blob/master/examples/rbac_model.conf)                                         | [rbac_policy.csv](https://github.com/casbin/casbin/blob/master/examples/rbac_policy.csv)                                         |
| RBAC with resource roles  | [rbac_with_resource_roles_model.conf](https://github.com/casbin/casbin/blob/master/examples/rbac_with_resource_roles_model.conf) | [rbac_with_resource_roles_policy.csv](https://github.com/casbin/casbin/blob/master/examples/rbac_with_resource_roles_policy.csv) |
| RBAC with domains/tenants | [rbac_with_domains_model.conf](https://github.com/casbin/casbin/blob/master/examples/rbac_with_domains_model.conf)               | [rbac_with_domains_policy.csv](https://github.com/casbin/casbin/blob/master/examples/rbac_with_domains_policy.csv)               |
| ReBAC                     | [rebac_model.conf](https://github.com/casbin/casbin/blob/master/examples/rebac_model.conf)                                       | [rebac_policy.csv](https://github.com/casbin/casbin/blob/master/examples/rebac_policy.csv)                                       |
| ABAC                      | [abac_model.conf](https://github.com/casbin/casbin/blob/master/examples/abac_model.conf)                                         | N/A                                                                                                                              |
| BLP                       | [blp_model.conf](https://github.com/casbin/casbin/blob/master/examples/blp_model.conf)                                           | N/A                                                                                                                              |
| Biba                      | [biba_model.conf](https://github.com/casbin/casbin/blob/master/examples/biba_model.conf)                                         | N/A                                                                                                                              |
| LBAC                      | [lbac_model.conf](https://github.com/casbin/casbin/blob/master/examples/lbac_model.conf)                                         | N/A                                                                                                                              |
| OrBAC                     | [orbac_model.conf](https://github.com/casbin/casbin/blob/master/examples/orbac_model.conf)                                       | [orbac_policy.csv](https://github.com/casbin/casbin/blob/master/examples/orbac_policy.csv)                                       |
| IP Match                  | [ipmatch_model.conf](https://github.com/casbin/casbin/blob/master/examples/ipmatch_model.conf)                                   | [ipmatch_policy.csv](https://github.com/casbin/casbin/blob/master/examples/ipmatch_policy.csv)                                   |
| RESTful                   | [keymatch_model.conf](https://github.com/casbin/casbin/blob/master/examples/keymatch_model.conf)                                 | [keymatch_policy.csv](https://github.com/casbin/casbin/blob/master/examples/keymatch_policy.csv)                                 |
| Deny-override             | [rbac_with_not_deny_model.conf](https://github.com/casbin/casbin/blob/master/examples/rbac_with_not_deny_model.conf)             | [rbac_with_deny_policy.csv](https://github.com/casbin/casbin/blob/master/examples/rbac_with_deny_policy.csv)                     |
| Allow-and-deny            | [rbac_with_deny_model.conf](https://github.com/casbin/casbin/blob/master/examples/rbac_with_deny_model.conf)                     | [rbac_with_deny_policy.csv](https://github.com/casbin/casbin/blob/master/examples/rbac_with_deny_policy.csv)                     |
| Priority                  | [priority_model.conf](https://github.com/casbin/casbin/blob/master/examples/priority_model.conf)                                 | [priority_policy.csv](https://github.com/casbin/casbin/blob/master/examples/priority_policy.csv)                                 |
| Explicit Priority         | [priority_model_explicit](https://github.com/casbin/casbin/blob/master/examples/priority_model_explicit.conf)                    | [priority_policy_explicit.csv](https://github.com/casbin/casbin/blob/master/examples/priority_policy_explicit.csv)               |
| Subject-Priority          | [subject_priority_model.conf](https://github.com/casbin/casbin/blob/master/examples/subject_priority_model.conf)                 | [subject_priority_policyl.csv](https://github.com/casbin/casbin/blob/master/examples/subject_priority_policy.csv)                |
