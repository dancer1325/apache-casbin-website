---
id: syntax-for-models
title: Model Syntax
description: Syntax for Models
keywords: [syntax]
authors: [nodece]
---

* | ANY model CONF file
  * `[request_definition], [policy_definition], [policy_effect], [matchers]`
    * if you use 
      * **RBAC** -> add `[role_definition]` section
      * **RBAC with constraints**
        * -> add `[constraint_definition]`
        * _Example:_ separation of duties 
    * lines / start with `#`
      * == comments

## `[request_definition]`

* == parameters / passed -- to -- `e.Enforce(...)`

## `[policy_definition]`


* | policy rules
  * 1 policy rule / EACH policy file's line
  * ' first token
    * == policy type / 
      * MUST match -- with a -- policy definition
  * ' ALL elements are treated -- as -- strings
    * [here](https://github.com/casbin/casbin/issues/113)

## `[policy_effect]`

* if MULTIPLE policies match (e.g. one allows, another denies) -> how to combine results 

* ALLOWED built-in policy effects
  * | CURRENT implementations,
    * ❌CUSTOM effect expressions, are NOT supported❌ 

| Policy Effect                                                  | Meaning                                                                                                                      | Example                                           |
|----------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------|
| `some(where (p.eft == allow))`                                 | allow-override <br/> &ensp;&ensp; if a matched policy has effect `allow` -> the result is allow                              | [ACL, RBAC, etc.](SupportedModels.md)    |
| `!some(where (p.eft == deny))`                                 | deny-override <br/> &ensp;&ensp; if NO matched policy has effect `deny` -> the result is allow                               | [Deny-override](SupportedModels)  |
| `some(where (p.eft == allow)) && !some(where (p.eft == deny))` | allow-and-deny <br/> &ensp;&ensp; requirements: >=1 `allow`& NO `deny` <br/> &ensp;&ensp;  if BOTH match -> `deny` overrides | [Allow-and-deny](SupportedModels) |
| `priority(p.eft) &#124;&#124; deny`                            | priority                                                                                                                     | [Priority](SupportedModels.md)                    |
| `subjectPriority(p.eft)`                                       | priority -- based on -- role                                                                                                 | [Subject-Priority](SupportedModels.md)            |

## `[constraint_definition]`

* requirements
  * `[role_definition]`
* OPTIONAL
* uses
  * | RBAC, invariants
    * _Example:_ separation of duties
* | role assignments change,
  * constraints are checked

### Constraint Types

* **Separation of Duties -- `sod` --**
  * == user can NOT hold both roles
  * if Alice has `finance_requester`, she cannot be assigned `finance_approver`.


* **Separation of Duties Max -- `sodMax` --**
  * limits the NUMBER roles | set / user can have

* **Role Cardinality -- `roleMax` --**
  * limits the NUMBER of users / can have a role

* **Prerequisite Role -- `rolePre` --** 
  * a role is REQUIRED BEFORE another

### how do constraints work?

* constrains are checked | 
  * group policies change (e.g. `AddGroupingPolicy()`, `RemoveGroupingPolicy()`),
    * == if a change would violate a constraint ->
      * operation fails -- with an -- error
      * policy is NOT updated
  * load the model
    * checked vs current policies

* invalid constraints
  * _Examples:_ bad syntax, missing RBAC setup, or current data / already violates a constraint
  * throw a clear error

## `[matchers]`

* == how requests are evaluated -- vs -- policy rules

* support
  * arithmetic operators
    * `+`, `-`, `*`, `/`
  * logical operators
    * `&&`, `||`, `!` 

### expression order | matchers

* affect
  * performance

* recommendations
  * place FIRST, cheaper matchers

## MULTIPLE section types

* == define MULTIPLE 
  * [request](#request_definition), 
  * [policy](#policy_definition),
  * [effect](#policy_effect),
  * [matcher](#matchers)

* add suffixes
  * _Examples:_
    * `r2`
      * == request definition 2
    * `p2`
      * == policy definition 2
    * `e2`
      * == policy effect 2
    * `m2`
      * == matcher 2
  * are paired -- by -- suffix
    * _Example:_ `r2` is matched -- , via `m2`, with -- `p2` + combined -- with -- `e2`

* how to use them?
  * set an `EnforceContext`
  * `.enforce(previouslyDefinedEnforceContext)`

* _Examples:_
  * [model](https://github.com/casbin/casbin/blob/master/examples/multiple_policy_definitions_model.conf) 
  * [policy](https://github.com/casbin/casbin/blob/master/examples/multiple_policy_definitions_policy.csv)

## `in` operator

* allows
  * checks whether the right-hand `[]someType` contains the left-hand value `someType`
    * | Go,
      * right-hand MUST be `[]interface{}` 

* vs `==`
  * NO type coercion

* _Examples:_
  * [rbac_model_matcher_using_in_op](https://github.com/casbin/casbin/blob/277c1a2b85698272f764d71a94d2595a8d425915/examples/rbac_model_matcher_using_in_op.conf)
  * [keyget2_model](https://github.com/casbin/casbin/blob/277c1a2b85698272f764d71a94d2595a8d425915/examples/keyget2_model.conf)
  * [keyget_model](https://github.com/casbin/casbin/blob/277c1a2b85698272f764d71a94d2595a8d425915/examples/keyget_model.conf)

## Expression evaluators OR expression engine

* provide 
  * a unified PERM syntax
  * extra features / EACH engine

* language-specific

| Implementation | Language | Expression Evaluator                                                                                                                                                                                     |
|----------------|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Casbin         | Golang   | [https://github.com/casbin/govaluate](https://github.com/casbin/govaluate)                                                                                                                               |
| jCasbin        | Java     | [https://github.com/killme2008/aviatorscript](https://github.com/killme2008/aviatorscript)                                                                                                               |
| Node-Casbin    | Node.js  | [https://github.com/donmccurdy/expression-eval](https://github.com/donmccurdy/expression-eval)                                                                                                           |
| PHP-Casbin     | PHP      | [https://github.com/symfony/expression-language](https://github.com/symfony/expression-language)                                                                                                         |
| PyCasbin       | Python   | [https://github.com/danthedeckie/simpleeval](https://github.com/danthedeckie/simpleeval)                                                                                                                 |
| Casbin.NET     | C#       | [https://github.com/davideicardi/DynamicExpresso](https://github.com/davideicardi/DynamicExpresso)                                                                                                       |
| Casbin4D       | Delphi   | [https://github.com/casbin4d/Casbin4D/tree/master/SourceCode/Common/Third%20Party/TExpressionParser](https://github.com/casbin4d/Casbin4D/tree/master/SourceCode/Common/Third%20Party/TExpressionParser) |
| casbin-rs      | Rust     | [https://github.com/jonathandturner/rhai](https://github.com/jonathandturner/rhai)                                                                                                                       |
| casbin-cpp     | C++      | [https://github.com/ArashPartow/exprtk](https://github.com/ArashPartow/exprtk)                                                                                                                           |

* impact
  * enforcement speed

* [Benchmarks](Benchmark.md)
