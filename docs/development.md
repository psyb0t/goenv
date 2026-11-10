---
title: "Development"
description: "Build, lint and test goenv with its Make targets in Docker or on the host, the 90 percent coverage gate, and how CI, badges and releases run."
nav_order: 4
---

# Development

Hacking on goenv itself. Yes, people do that. You need Go 1.26 or newer and Make, or just Docker. The module has no third-party dependencies, so there is no `go.sum` and nothing to install.

## Make targets

Run `make help` for the list. The [`Makefile`](https://github.com/psyb0t/goenv/blob/main/Makefile) has:

| Target | What it does |
| --- | --- |
| `make all` | Runs `dep`, `lint` and `test-coverage`, in that order. |
| `make dep` | `go mod tidy`. |
| `make lint` | Fails if `go fix -diff ./...` has suggestions, then runs `go vet ./...`. |
| `make lint-fix` | Applies `go fix ./...`, then runs `go vet ./...`. |
| `make test` | `go test -race ./...`. |
| `make test-coverage` | Runs the tests with the race detector and a coverage profile, writes the total to `coverage-percent.txt`, and fails below 90 percent. Coverage is 100 percent. |
| `make docs` | Builds this site into `.testing/docs-site` with GitHub Pages' own Jekyll image. Needs Docker. |

[`pre-commit.sh`](https://github.com/psyb0t/goenv/blob/main/pre-commit.sh) runs `make lint && make test-coverage`, the same pair CI runs.

### Without Go on the host

The Go targets run unchanged in the official image:

```bash
docker run --rm \
  --user "$(id -u):$(id -g)" \
  -e HOME=/tmp \
  -e GOCACHE=/tmp/gocache \
  -e GOPATH=/tmp/go \
  -v "$PWD:/src" \
  -w /src \
  golang:1.26 \
  make dep lint test-coverage
```

## Layout

| Path | Holds |
| --- | --- |
| [`goenv.go`](https://github.com/psyb0t/goenv/blob/main/goenv.go) | The whole library. |
| [`goenv_test.go`](https://github.com/psyb0t/goenv/blob/main/goenv_test.go) | Table tests for `Get`, `IsProd` and `IsDev`, using `t.Setenv`. |
| [`.agents/`](https://github.com/psyb0t/goenv/tree/main/.agents) | The [agent skill](agent-integrations.md), the plugin manifests and the contributor rules. |
| [`AGENTS.md`](https://github.com/psyb0t/goenv/blob/main/AGENTS.md) | Instructions for coding agents working on the repository. |
| [`docs/`](https://github.com/psyb0t/goenv/tree/main/docs) | This site, built by GitHub Pages from the `main` branch. |

## CI

[`.github/workflows/pipeline.yml`](https://github.com/psyb0t/goenv/blob/main/.github/workflows/pipeline.yml) runs on every push, branch or tag, once a week on a schedule, and by hand.

- On every branch push, the schedule and manual runs, the shared Go workflow runs `make lint`, `make test-coverage` and govulncheck with Go 1.26.
- When that run is on `main`, the badges job then publishes the coverage, version, license and imported-by badges to the `badges` branch.
- On a tag push, the pipeline does not rerun the checks. It waits for the `main` run of the same commit to pass, cuts the GitHub release with the tag message as its notes, publishes the agent skill to ClawHub, and refreshes the badges with the coverage from that `main` run. The `main` run keeps its coverage artifact for one day, so a tag pushed later than that still gets its release, but the badges job fails.

The weekly run exists for the imported-by badge, which pkg.go.dev updates on its own schedule.

## Releases

Versions follow [semver](https://semver.org). The Git tag is the version; there is no version file. The one place that repeats it is the `version` in the Codex plugin manifest, which the release commit sets to the new version. A release is a commit on `main` that has passed CI, then an annotated `vX.Y.Z` tag on it whose message is that version's [CHANGELOG](https://github.com/psyb0t/goenv/blob/main/CHANGELOG.md) section.
