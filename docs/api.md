---
title: "API reference"
description: "Every exported name in goenv: Get, IsProd, IsDev, the Prod, Dev and EnvVarName constants, the Type alias, and which ENV values mean dev."
nav_order: 2
---

# API reference

The whole package is one file, [`goenv.go`](https://github.com/psyb0t/goenv/blob/main/goenv.go): three functions, three constants and one type. That's it. There's nothing else to find.

```go
import "github.com/psyb0t/goenv"
```

## Functions

| Function | Returns | What it's basically doing |
| --- | --- | --- |
| `Get() Type` | `Dev` when `ENV` is exactly `"dev"`, otherwise `Prod` | `if os.Getenv("ENV") == "dev" { return "dev" } else { return "prod" }` |
| `IsProd() bool` | `true` unless `ENV` is exactly `"dev"` | `Get() == Prod` |
| `IsDev() bool` | `true` only when `ENV` is exactly `"dev"` | `Get() == Dev` |

All three read the variable with `os.Getenv` on every call. Nothing is cached, so if your process changes `ENV` at runtime, the next call sees the new value. In practice set it once before the process starts and leave it alone.

None of them return an error. There is nothing that can fail.

## Constants

| Constant | Type | Value | Surprise level |
| --- | --- | --- | --- |
| `EnvVarName` | untyped string | `"ENV"` | None |
| `Prod` | `Type` | `"prod"` | None |
| `Dev` | `Type` | `"dev"` | None |

`EnvVarName` is there so you can reference the name, not change it. The variable goenv reads is always `ENV`.

## Type

```go
type Type = string
```

`Type` is a type alias, not a new type. `goenv.Type` is `string`, so any string is assignable to it and the compiler won't stop you comparing `Get()` against `"staging"`. `Prod` and `Dev` are convenient names, not a closed enum.

## Which values mean dev

`Get()` does one exact, case-sensitive `switch` on the value of `ENV`:

| `ENV` value | `Get()` | `IsProd()` | `IsDev()` |
| --- | --- | --- | --- |
| `dev` | `dev` | `false` | `true` |
| `prod` | `prod` | `true` | `false` |
| unset | `prod` | `true` | `false` |
| empty | `prod` | `true` | `false` |
| `DEV`, `Dev` | `prod` | `true` | `false` |
| `development`, `staging`, `test` | `prod` | `true` | `false` |
| anything else | `prod` | `true` | `false` |

Only the literal lowercase `dev` counts as dev. That's on purpose: a missing or misspelled `ENV` fails safe toward production behavior instead of quietly switching on your dev-only shortcuts.

```bash
export ENV=dev   # you're developing
export ENV=prod  # you're in production
export ENV=      # also production, because paranoia is a feature
```

## Limits

- Two environments only. If you need `staging`, `test` or `ci` as real, separate environments, goenv can't express it. Use a proper config library.
- The variable name is fixed at `ENV`.
- No type safety beyond `string`, as explained under [Type](#type).
