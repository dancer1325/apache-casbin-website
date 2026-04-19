---
id: policy-subset-loading
title: Loading Policy Subsets
description: Loading filtered policies
keywords: [filtered policies, policy subset, performance optimization]
authors: [hsluoyz]
---

* **filtered policy loading**
  * supported | SOME [adapters](Adapters.md)
    * == -- depend on -- the adapter
  * == ONLY load a subset of policies 
    * -- from -- storage
    * -- according to the -- filter
  * allows
    * reducing memory
    * speed up loading
  * use cases
    * large OR multi-tenant setups
      * Reason:🧠OTHERWISE, loading the full policy set == impractical🧠
  * `SavePolicy()`
    * is disabled
      * -> by mistake, you do NOT overwrite the full policy
  * ways
    * -- via --     
      * `LoadFilteredPolicy(filter)` + supported adapter
        * replace the filtered policies
      * `LoadIncrementalFilteredPolicy(filter)`
        * adds the filtered policies | CURRENT in-memory policy
  * ALLOWED
    * \>1 filtered loads / enforcer
