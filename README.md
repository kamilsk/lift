> # 🏋️‍♂️ lift
>
> Up your service locally.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]
[![Coverage][coverage.icon]][coverage.page]

## 💡 Idea

```bash
$ lift call && lift up
```

A full description of the idea is available [here][design.page].

## 🏆 Motivation

In [Avito](https://tech.avito.ru) we have an excellent [PaaS](https://en.wikipedia.org/wiki/Platform_as_a_service)
which helps us to run our services in [Kubernetes](https://kubernetes.io) clusters with just a few commands.
But I want to run it so quickly and frequently as possible to debug during development.
For that reason, I need a possibility to up services written on [Go](https://golang.org) locally from IDE like
[GoLand](https://www.jetbrains.com/go/) without losing the benefits that
[minikube](https://github.com/kubernetes/minikube) provides.

## 🤼‍♂️ How to

...rethinking

## 🧩 Installation

### Homebrew

```bash
$ brew install kamilsk/tap/lift
```

### Binary

```bash
$ curl -sSfL https://raw.githubusercontent.com/kamilsk/lift/master/bin/install | sh
# or
$ wget -qO-  https://raw.githubusercontent.com/kamilsk/lift/master/bin/install | sh
```

> Don't forget about [security](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/).

### Source

```bash
# use standard go tools
$ go get github.com/kamilsk/lift@latest
# or use egg tool
$ egg tools add github.com/kamilsk/lift@latest
```

> [egg][] is an `extended go get`.

### Bash and Zsh completions

```bash
$ lift completion bash > /path/to/bash_completion.d/lift.sh
$ lift completion zsh  > /path/to/zsh-completions/_lift.zsh
# or autodetect
$ source <(lift completion)
```

> See `kubectl` [documentation](https://kubernetes.io/docs/tasks/tools/install-kubectl/#enabling-shell-autocompletion).

## 🤲 Outcomes

### Patches

- [github.com/pelletier/go-toml](https://github.com/pelletier/go-toml)
  - [pull request](https://github.com/pelletier/go-toml/pull/281)
  - [release](https://github.com/pelletier/go-toml/releases/tag/v1.5.0)

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.com/kamilsk/lift
[build.icon]:       https://travis-ci.com/kamilsk/lift.svg?branch=master
[coverage.page]:    https://codeclimate.com/github/kamilsk/lift/test_coverage
[coverage.icon]:    https://api.codeclimate.com/v1/badges/58b81affe3a64e409047/test_coverage
[design.page]:      https://www.notion.so/octolab/lift-9078cdbe27c842498f0561b6acd88a4d?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://github.com/kamilsk/lift
[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg]:              https://github.com/kamilsk/egg
