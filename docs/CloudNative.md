---
id: cloud-native
title: Cloud Native
description: Cloud native authorization middlewares
keywords: [cloud native, authorization middleware, middlewares]
authors: [ErikQQY]
---

* Casbin
  * | cloud-native stacks (Envoy, Istio, Traefik, KubeSphere, etc.),
    * can run -- as an -- authorization layer

## Projects / provide adapters OR sidecars -- by -- language

### Go

| Project                                                                 | Author    | Description                                                                                                                                                                                                                                                |
|-------------------------------------------------------------------------|-----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [envoy-authz](https://github.com/casbin/envoy-authz)                    | Casbin    | Authorization middleware -- for -- [Istio](https://istio.io/) & [Envoy](https://envoyproxy.io/)                                                                                                                                                            |
| [kubesphere-authz](https://github.com/casbin/kubesphere-authz)          | Casbin    | Authorization middleware -- for -- [KubeSphere](https://kubesphere.io/)                                                                                                                                                                                    |
| [casbin-forward-auth](https://github.com/grepplabs/casbin-forward-auth) | grepplabs | Authorization middleware -- for -- [Traefik](https://traefik.io/) & [NGINX](https://nginx.org/) & [HAProxy](https://www.haproxy.com/) & [Envoy Gateway](https://gateway.envoyproxy.io/) & [Istio](https://istio.io/) & [Envoy](https://www.envoyproxy.io/) |

### Node.js

| Project                                       | Author                                        | Description                                                                        |
|-----------------------------------------------|-----------------------------------------------|------------------------------------------------------------------------------------|
| [ODPF Shield](https://github.com/odpf/shield) | [Open Data Platform](https://github.com/odpf) | ODPF Shield is a cloud native role-based authorization-aware reverse-proxy service |
