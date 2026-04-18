---
id: enforcers
title: Enforcers
description: The Enforcer is the main structure in Casbin that acts as an interface for users to perform operations on policy rules and models.
keywords: [enforcer, supported enforcers]
authors: [Abingcbc]
---

* **Enforcer**
  * == Casbin’s main interface
  * uses
    * enforce policy
    * read OR update 
      * policies
      * models

## AVAILABLE enforcers

### | Go


| Enforcer                                                                                       | Author | Description                                                                                                                                                                                                                                                                                                                              |
|------------------------------------------------------------------------------------------------|--------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Enforcer](https://github.com/casbin/casbin/blob/master/enforcer.go)                           | Casbin | == base `Enforcer` <br/> provides <br/>  &ensp; primary interface -- for -- policy & model operations <br/> [Management API documentation](ManagementAPI.md)                                                                                                                                                                             |
| [CachedEnforcer](https://github.com/casbin/casbin/blob/master/enforcer_cached.go)              | Casbin | extends the `Enforcer` -- with -- in-memory caching OF enforcement results <br/> supports: configurable cache expiration <br/>  includes -- , via read-write locks, -- thread-safe access <br/> if you want to toggle caching (by default, enabled) -> use `EnableCache`  <br/>  remaining API methods match -- `Enforcer`'s API methods |
| [DistributedEnforcer](https://github.com/casbin/casbin/blob/master/enforcer_distributed.go)    | Casbin | TODO: Designed for distributed deployments, this enforcer wraps `SyncedEnforcer` to work with dispatchers. Additional information is available in the [dispatcher documentation](/docs/dispatchers#distributedenforcer).                                                                                                                 |
| [SyncedEnforcer](https://github.com/casbin/casbin/blob/master/enforcer_synced.go)              | Casbin | Extends `Enforcer` with synchronized access for thread-safe operations.                                                                                                                                                                                                                                                                  |
| [SyncedCachedEnforcer](https://github.com/casbin/casbin/blob/master/enforcer_cached_synced.go) | Casbin | Combines the caching features of `CachedEnforcer` with the synchronization of `SyncedEnforcer`.                                                                                                                                                                                                                                          |
| [ContextEnforcer](https://github.com/casbin/casbin/blob/master/enforcer_context.go)            | Casbin | Implements the `IEnforcerContext` interface to provide context-aware operations. Currently supports context for adapter operations like `LoadPolicyCtx()`, `SavePolicyCtx()`, `AddPolicyCtx()`, and `RemovePolicyCtx()`. This interface is designed for future expansion of context support across additional Casbin operations.         |

### Python

| Enforcer                                                                                    | Author | Description                                                                                                                                                                                                                                                                                                                                                                        |
|---------------------------------------------------------------------------------------------|--------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Enforcer](https://github.com/casbin/pycasbin/blob/master/casbin/enforcer.py)                        | Casbin | The base `Enforcer` provides the primary interface for policy and model operations. See the [Management API documentation](/docs/management-api) for detailed API information.                                                                                                                                                                                                             |
| [DistributedEnforcer](https://github.com/casbin/pycasbin/blob/master/casbin/distributed_enforcer.py) | Casbin | Designed for distributed deployments, this enforcer wraps `SyncedEnforcer` to work with dispatchers. Additional information is available in the [dispatcher documentation](/docs/dispatchers#distributedenforcer).                                                                                                                                                                        |
| [SyncedEnforcer](https://github.com/casbin/pycasbin/blob/master/casbin/synced_enforcer.py)           | Casbin | Extends `Enforcer` with synchronized access for thread-safe operations.                                                                                                                                                                                                                                                                                                           |
| [AsyncEnforcer](https://github.com/casbin/pycasbin/blob/master/casbin/async_enforcer.py)| Casbin | Provides asynchronous API methods for non-blocking enforcement operations.                                                                                                                                                                                                                                                                                                 |
| [FastEnforcer](https://github.com/casbin/pycasbin/blob/master/casbin/fast_enforcer.py)| Casbin | Uses an optimized model architecture offering up to 50× performance improvements over the standard enforcer. More details are available at [fastbin's repository](https://github.com/wakemaster39/fastbin).                                                                                                                                                                                                                                                                                                 |

