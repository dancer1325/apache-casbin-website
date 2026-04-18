---
id: model-storage
title: Model Storage
description: Model storage
keywords: [model storage]
authors: [hsluoyz]
---

* models | Casbin
  * 💡are load-ONLY💡
    * == 
      * define your access control logic &
      * | runtime, treated -- as -- static 
    * ❌NO exist API / save OR update the model | storage❌
    * 💡ways to load a model💡
      * [-- from -- a CONF file](#---from----a-conf-file)
      * -- from -- code
      * -- from -- a string

## -- from -- a CONF file

* MOST COMMON
  * Reason:🧠easy to
    * share models
    * discuss models🧠

* _Example:_ [rbac_model.conf](https://github.com/casbin/casbin/blob/master/examples/rbac_model.conf)

* uses
  * | create the enforcer

    ```go
    e := casbin.NewEnforcer("examples/rbac_model.conf", "examples/rbac_policy.csv")
    ```
