> # 🏋️‍♂️ lift
>
> Up your service locally.

[![Build Status][icon_build]][page_build]

## The concept

```bash
$ eval $(lift up)
```

## Motivation

In [Avito](https://tech.avito.ru) we have an excellent [PaaS](https://en.wikipedia.org/wiki/Platform_as_a_service)
which helps us to run our services in [Kubernetes](https://kubernetes.io) clusters with just a few commands.
But I want to run it so quickly and frequently as possible to debug during development.
For that reason, I need a possibility to up services written on [Go](https://golang.org) locally from IDE like
[GoLand](https://www.jetbrains.com/go/) without losing the benefits that
[minikube](https://github.com/kubernetes/minikube) provides.

## How to

0. Describe your dependencies and storage in `app.toml`.

1. Define everything you need in `env_vars` and `envs.local.env_vars`. See an [example](testdata/app.toml).

2. Dump environment variables into `.env` file

```bash
$ lift env > bin/.env
```

3. Use it in your IDE with [EnvFile](https://plugins.jetbrains.com/plugin/7861-envfile) plugin

![GoLand integration](.github/goland_integration.png)

4. Forward required ports using `kubectl port-forward`, or use my other tool named
[forward](https://github.com/kamilsk/forward) for that

```bash
$ lift forward
forward -- demo-local-postgresql- 5432 demo-local-redis- 6379

$ eval $(lift forward)
```

5. or run all together

```bash
$ lift up
export SERVICE_A_URL="http://service-a.cluster/";
export SERVICE_C_URL="http://service-c.cluster/";
export SERVICE_B_URL="http://service-b.cluster/";
export PGHOST="localhost";
export PGPORT="5432";
export PGDATABASE="master";
export PGUSER="postgres";
export PGPASSWORD="";
export REDIS_PORT="6379";
export CUSTOM="variable";
forward -- demo-local-postgresql- 5432 demo-local-redis- 6379 &;
go run cmd/service/main.go;
ps | grep '[f]orward --' | awk '{print $1}' | xargs kill -SIGKILL

$ eval $(lift up)
```

### Good to have

```makefile
.PHONY: lift
lift:
	@lift env > bin/.env

.PHONY: up
up:
	@avito start
	@avito service dev --no-watch

.PHONY: down
down:
	@avito service deletelocal
```

## Installation

### Homebrew

```bash
$ brew install kamilsk/tap/lift
```

### Binary

```bash
$ REQ_VER=0.0.3  # all available versions are on https://github.com/kamilsk/lift/releases/
$ REQ_OS=Linux   # macOS is also available
$ REQ_ARCH=64bit # 32bit is also available
# wget -q -O lift.tar.gz
$ curl -sL -o lift.tar.gz \
       https://github.com/kamilsk/lift/releases/download/"${REQ_VER}/lift_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf lift.tar.gz -C "${GOPATH}"/bin/ && rm lift.tar.gz
```

---

made with ❤️ for everyone

[icon_build]:      https://travis-ci.org/kamilsk/lift.svg?branch=master

[page_build]:      https://travis-ci.org/kamilsk/lift
[page_promo]:      https://github.com/kamilsk/lift
