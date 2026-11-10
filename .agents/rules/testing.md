# Testing

Prefix: TEST. Commands are in `AGENTS.md`.

- **TEST1:** Tests live in `goenv_test.go`, in `package goenv`, and set the variable with `t.Setenv(EnvVarName, ...)`.
- **TEST3:** `make test-coverage` fails below `MIN_TEST_COVERAGE := 90` in the `Makefile`. Coverage is 100% today.
- **TEST4:** `make lint` fails when `go fix -diff` prints anything; `make lint-fix` applies the rewrites.
- **TEST5:** CI runs `make lint`, `make test-coverage` and govulncheck through `go-workflow.yml` with Go 1.26.
