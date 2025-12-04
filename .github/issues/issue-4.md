---
id: 4
database_id: 443745523
node_id: MDU6SXNzdWU0NDM3NDU1MjM=
status: closed
title: "how to"
labels: ["type:docs"]
url: https://github.com/kamilsk/lift/issues/4
created_at: 2019-05-14T07:08:23Z
updated_at: 2019-05-16T08:08:38Z
---

# how to

documentation
- [x] motivation
  Avito has an excellent PaaS which helps developers to run their services in Kubernetes clusters...
- [x] requirements
  - forward tool
  - avito tool
  - paas-service
  - EnvFile extension https://plugins.jetbrains.com/plugin/7861-envfile
- [x] good to have
```make
.PHONY: up
up:
	@avito service dev --no-watch

.PHONY: down
down:
	@avito service deletelocal

.PHONY: lift
lift:
	@lift env > bin/.env
```
- [x] step by step
  1. define all you need in `envs.local.env_vars`, storages and dependencies
  2. run `make up`
  3. run `eval $(lift up)`, or use IDE with `.env file` (`make lift`)
