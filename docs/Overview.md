---
id: overview
title: Overview
description: Casbin Overview
keywords: [overview, supported languages]
authors: [hsluoyz]
---

* Casbin
  * [overview](../README.md)
  * [supported models](SupportedModels.md)
  * 💡how to implement rule-based access control?💡
    * | policy file (.csv),
      * define: subjects + objects + permitted actions
    * | model file (.conf),
      * define: authorization logic—layout + execution flow (== effect) + matchers
    * | Enforcer component,
      * evaluates incoming requests vs your model & policy

## languages / supported -- by -- Casbin

| [![golang](/img/langs/golang.png)](https://github.com/casbin/casbin) | [![java](/img/langs/java.png)](https://github.com/casbin/jcasbin) | [![nodejs](/img/langs/nodejs.png)](https://github.com/casbin/node-casbin) | [![php](/img/langs/php.png)](https://github.com/php-casbin/php-casbin) |
|----------------------------------------------------------------------|-------------------------------------------------------------------|---------------------------------------------------------------------------|------------------------------------------------------------------------|
| [Casbin](https://github.com/casbin/casbin)                           | [jCasbin](https://github.com/casbin/jcasbin)                      | [node-Casbin](https://github.com/casbin/node-casbin)                      | [PHP-Casbin](https://github.com/php-casbin/php-casbin)                 |
| Production-ready                                                     | Production-ready                                                  | Production-ready                                                          | Production-ready                                                       |

| [![python](/img/langs/python.png)](https://github.com/casbin/pycasbin) | [![dotnet](/img/langs/dotnet.png)](https://github.com/casbin/Casbin.NET) | [![c++](/img/langs/cpp.png)](https://github.com/casbin/casbin-cpp) | [![rust](/img/langs/rust.png)](https://github.com/casbin/casbin-rs) |
|------------------------------------------------------------------------|--------------------------------------------------------------------------|--------------------------------------------------------------------|---------------------------------------------------------------------|
| [PyCasbin](https://github.com/casbin/pycasbin)                         | [Casbin.NET](https://github.com/casbin/Casbin.NET)                       | [Casbin-CPP](https://github.com/casbin/casbin-cpp)                 | [Casbin-RS](https://github.com/casbin/casbin-rs)                    |
| Production-ready                                                       | Production-ready                                                         | Production-ready                                                   | Production-ready                                                    |

### features -- by -- language

| Feature                 |   Go   | Java  | Node.js |  PHP  | Python |  C#   | Delphi | Rust  |  C++  |  Lua  | Dart  | Elixir |
|:------------------------|:------:|:-----:|:-------:|:-----:|:------:|:-----:|:------:|:-----:|:-----:|:-----:|:-----:|:------:|
| Enforcement             |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ✅    |   ✅   |   ✅   |   ✅   |   ✅   |   ✅    |
| RBAC                    |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ✅    |   ✅   |   ✅   |   ✅   |   ✅   |   ✅    |
| ABAC                    |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ✅    |   ✅   |   ✅   |   ✅   |   ✅   |   ✅    |
| Scaling ABAC (`eval()`) |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ❌    |   ✅   |   ✅   |   ✅   |   ✅   |   ✅    |
| Adapter                 |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ✅    |   ✅   |   ✅   |   ✅   |   ✅   |   ❌    |
| Management API          |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ✅    |   ✅   |   ✅   |   ✅   |   ✅   |   ✅    |
| RBAC API                |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ✅    |   ✅   |   ✅   |   ✅   |   ✅   |   ✅    |
| Batch API               |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ❌    |   ✅   |   ✅   |   ✅   |   ❌   |   ❌    |
| Filtered Adapter        |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ❌    |   ✅   |   ✅   |   ✅   |   ❌   |   ❌    |
| Watcher                 |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ✅    |   ✅   |   ✅   |   ✅   |   ❌   |   ❌    |
| Role Manager            |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ✅   |   ❌    |   ✅   |   ✅   |   ✅   |   ✅   |   ❌    |
| Multi-Threading         |   ✅    |   ✅   |    ✅    |   ❌   |   ✅    |   ❌   |   ❌    |   ✅   |   ❌   |   ❌   |   ❌   |   ❌    |
| 'in' of matcher         |   ✅    |   ✅   |    ✅    |   ✅   |   ✅    |   ❌   |   ✅    |   ❌   |   ❌   |   ❌   |   ✅   |   ✅    |

* ✅ | Watcher OR Role Manager,
  * == interface exists | core library
  * ❌!= implementation is AVAILABLE | that language❌
