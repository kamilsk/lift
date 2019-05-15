> # 🏋️‍♂️ lift
>
> Up your service locally.

[![Patreon][icon_patreon]](https://www.patreon.com/octolab)
[![License][icon_license]](LICENSE)

## The concept

```bash
$ lift env > .env

$ eval $(lift up)
```

## Installation

### Homebrew

```bash
$ brew install kamilsk/tap/lift
```

### Binary

```bash
$ REQ_VER=0.0.2  # all available versions are on https://github.com/kamilsk/lift/releases/
$ REQ_OS=Linux   # macOS is also available
$ REQ_ARCH=64bit # 32bit is also available
# wget -q -O lift.tar.gz
$ curl -sL -o lift.tar.gz \
       https://github.com/kamilsk/lift/releases/download/"${REQ_VER}/lift_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf lift.tar.gz -C "${GOPATH}"/bin/ && rm lift.tar.gz
```

---

made with ❤️ for everyone

[icon_license]: https://img.shields.io/badge/license-MIT-blue.svg
[icon_patreon]: https://img.shields.io/badge/patreon-donate-orange.svg
