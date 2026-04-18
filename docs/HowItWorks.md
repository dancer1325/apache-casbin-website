---
id: how-it-works
title: How It Works
description: How Casbin Works
keywords: [PERM, request, policy, matcher, effect]
authors: [nodece]
---

* access control
  * == how users & resources interact
  * expressed -- as -- CONF files /
    * -- based on the -- ⭐️ **PERM metamodel (Policy, Effect, Request, Matchers)** ⭐️

* custom models
  * == combine [supported models](SupportedModels.md)
    * _Example:_ RBAC roles + ABAC attributes 

### Request

* == access request's parameters

* basic request
  * == 💡**subject** (who is requesting) + **object** (the resource) + **action** (the operation)💡
    * ⚠️EXACTLY this order ⚠️
    * _Example:_ `r = sub, obj, act`

### Policy

* == access rules' shape
  * == field names + order
    * if `eft` (policy effect) is omitted -> 
      * the effect field | policy file, is ignored
      * matching policies -- are treated as -- "allow"
        * == by default, "allow"
  * _Examples:_  
    * `p = sub, obj, act`
    * `p = sub, obj, act, eft`

### Matcher

* == how to match a request -- against -- policies
  * _Example:_ `m = r.sub == p.sub && r.act == p.act && r.obj == p.obj`
* if you want to split matchers ACROSS lines -> end EACH line -- with -- `\`

* operators
  * ⚠️-- depend on -- Casbin language implementation⚠️ 
    * | [Go implementation](https://github.com/casbin/govaluate),
      * `in` operator
        * requirements
          * \>=1 element

### Effect

* == ALL matched policies' effects + logical expression
  * _Examples:_ 
    * `e = some(where (p.eft == allow))`
    * `e = some(where (p.eft == allow)) && !some(where (p.eft == deny))`
      * allow the result ONLY if
        * \>=1 policy, allows
        * NO policy deny
      * if allow & deny match -> deny wins

* effect
  * requirements
    * request subject + object + action / match a policy’s fields
  * is stored | `p.eft`
