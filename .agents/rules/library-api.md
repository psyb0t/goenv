# Library API

Prefix: API. The exported names in `goenv.go` are the public API of a module that other repositories import (see the imported-by badge in the README).

- **API1:** The exported set is `EnvVarName`, `Type`, `Prod`, `Dev`, `Get`, `IsProd` and `IsDev`. Removing or renaming one, or changing a signature, breaks importers. Do it only when the owner asks for a breaking release.
- **API2:** `Get()` returns `Dev` only for the exact, case-sensitive value `"dev"`, and `Prod` for everything else, unset and empty included. Do not add case folding, trimming, aliases such as `"development"`, or more environments without the owner asking; the fail-safe default to `prod` is the point of the package.
- **API3:** `Type` is an alias (`type Type = string`). Turning it into a defined type changes the API for every caller that mixes it with `string`, so it counts as a breaking change under API1.
- **API4:** `Get()` reads `os.Getenv` on every call. Do not cache the value; the tests rely on `t.Setenv` taking effect between calls.
- **API5:** The library imports only the standard library (today only `os`). Do not add a third-party module; the README and the skill both promise zero dependencies.
