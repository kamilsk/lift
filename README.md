> # 🏋️‍♂️ lift
>
> Up your service locally.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## 💡 Idea

```bash
$ eval $(lift up)
```

Full description of the idea is available [here][design.page].

## 🏆 Motivation

In [Avito](https://tech.avito.ru) we have an excellent [PaaS](https://en.wikipedia.org/wiki/Platform_as_a_service)
which helps us to run our services in [Kubernetes](https://kubernetes.io) clusters with just a few commands.
But I want to run it so quickly and frequently as possible to debug during development.
For that reason, I need a possibility to up services written on [Go](https://golang.org) locally from IDE like
[GoLand](https://www.jetbrains.com/go/) without losing the benefits that
[minikube](https://github.com/kubernetes/minikube) provides.

## 🤼‍♂️ How to

0. Describe your dependencies and storage in `app.toml`.

1. Define everything you need in `env_vars` and `envs.local.env_vars`. See an [example](testdata/app.toml).

2. Dump environment variables into `.env` file

```bash
$ lift env > bin/.env
```

3. Use it in your IDE with [EnvFile](https://plugins.jetbrains.com/plugin/7861-envfile) plugin

![GoLand integration](.github/goland_integration.png)

  - 🔦 tip: use `shift + cmd + .` to see hidden dot files

4. Forward required ports using `kubectl port-forward`, or use my other tool named
[forward](https://github.com/kamilsk/forward) for that

```bash
$ eval $(lift forward)
```

5. or run all together

```bash
$ eval $(lift up -- go run cmd/service/main.go); eval $(lift down)

$ eval $(lift forward -d)
$ lift call -- go run cmd/service/main.go
$ eval $(lift down)
```

### Good to have

```makefile
.PHONY: format
format:
	@lift call -- goimports -local '$$GOMODULE' -ungroup -w .

.PHONY: lift
lift:
	@lift env > bin/.env

.PHONY: up
up:
	@avito start
	@avito service dev --no-watch
	@nohup $$(lift forward) &

.PHONY: status
status:
	@minikube status
	@kubectl get pod

.PHONY: forward
forward:
	@eval $$(lift forward)

.PHONY: down
down:
	@-avito service deletelocal
	@-eval $$(lift down)
```

![goimports integration](.github/goimports_integration.png)

## 🧩 Installation

### Homebrew

```bash
$ brew install kamilsk/tap/lift
```

### Binary

```bash
$ curl -sSL https://bit.ly/install-lift | sh
# or
$ wget -qO- https://bit.ly/install-lift | sh
```

### Source

```bash
# use standard go tools
$ go get -u github.com/kamilsk/lift
# or use egg tool
$ egg tools add github.com/kamilsk/lift
```

> [egg][egg.page]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

### Bash and Zsh completions

```bash
$ lift completion bash > /path/to/bash_completion.d/lift.sh
$ lift completion zsh  > /path/to/zsh-completions/_lift.zsh
```

## 🤲 Outcomes

### Patches

- [github.com/pelletier/go-toml](https://github.com/pelletier/go-toml)
  - [pull request](https://github.com/pelletier/go-toml/pull/281)
  - [release](https://github.com/pelletier/go-toml/releases/tag/v1.5.0)

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.icon]:       https://travis-ci.org/kamilsk/lift.svg?branch=master
[build.page]:       https://travis-ci.org/kamilsk/lift

[design.page]:      https://www.notion.so/octolab/lift-9078cdbe27c842498f0561b6acd88a4d?r=0b753cbf767346f5a6fd51194829a2f3

[promo.page]:       https://github.com/kamilsk/lift

[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg.page]:         https://github.com/kamilsk/egg
