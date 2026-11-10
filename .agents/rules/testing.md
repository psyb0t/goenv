# Testing

Prefix: TEST. Commands are in `AGENTS.md`.

- **TEST1:** Tests live in `goenv_test.go`, in `package goenv`, as table tests with `t.Run` subtests. Set the variable with `t.Setenv(EnvVarName, ...)`, never `os.Setenv`, so each case restores the environment.
- **TEST2:** A behavior change adds or updates a table row for each of `Get`, `IsProd` and `IsDev` it affects.
- **TEST3:** `make test-coverage` fails below `MIN_TEST_COVERAGE := 90` in the `Makefile`. Coverage is 100% today; keep it there rather than lowering the threshold.
- **TEST4:** `make lint` fails when `go fix -diff` prints anything. Run `make lint-fix` and commit the result instead of silencing it.
- **TEST5:** CI runs `make lint`, `make test-coverage` and govulncheck through `go-workflow.yml` with Go 1.26. Run `make lint test-coverage` (on the host or in `golang:1.26`) before calling a change done.
