---
title: "Home"
description: "goenv is a zero-dependency Go library that reads the ENV variable and tells you if you're in prod or dev, falling back to prod for anything else."
nav_order: 1
permalink: /
seo:
  type: WebSite
---

# goenv

The most over-engineered environment variable reader in the history of Go programming.

goenv reads one environment variable, `ENV`, and tells you whether you're in `prod` or `dev`. `ENV=dev` gets you dev. Anything else gets you prod: unset, empty, `staging`, `DEV`, a typo. That is on purpose, because if you didn't explicitly ask for dev I'm not handing it to you.

## Install

```bash
go get github.com/psyb0t/goenv
```

The module declares `go 1.26` in its [`go.mod`](https://github.com/psyb0t/goenv/blob/main/go.mod), so you need Go 1.26 or newer. The only import is the standard library `os`.

## Quick look

```go
package main

import (
    "fmt"

    "github.com/psyb0t/goenv"
)

func main() {
    fmt.Println("running in", goenv.Get()) // "prod" or "dev"

    if goenv.IsDev() {
        fmt.Println("break whatever you want")
    }

    if goenv.IsProd() {
        fmt.Println("don't fuck this up")
    }
}
```

```bash
ENV=dev go run .   # running in dev
go run .           # running in prod
```

## Where to go next

| You want to | Read |
| --- | --- |
| Look up every function and constant, and see which `ENV` values count as dev | [API reference](api.md) |
| Install the goenv agent skill in Claude Code, Codex or OpenClaw | [Agent integrations](agent-integrations.md) |
| Build, lint and test the library itself, or see how CI and releases run | [Development](development.md) |

The Go package docs are on [pkg.go.dev](https://pkg.go.dev/github.com/psyb0t/goenv), and the source is on [GitHub](https://github.com/psyb0t/goenv).
