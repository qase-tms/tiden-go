# tiden-go

Go reporting for [Tiden](https://tiden.ai) — the continuous quality system.

## What's here

| Package | Purpose |
|---|---|
| `pkg/gotest` | `go test -json` (test2json) stream parser + the **stable case-identity rules** shared by every Tiden Go reporting surface |

The primary consumer today is the `tiden` CLI's zero-touch wrapper:

```bash
tiden run exec -- go test ./...
```

It creates a Tiden test run, streams your suite's results into it, and
completes the run with the real verdict — no changes to your test code.

An in-process reporting library (`tiden.Test(t, ...)` with steps, params and
attachments) is the planned second phase and will live in this repository.

## Identity rules (the part you must not break)

- **Leaf subtest = case.** `TestLogin/expired_token` is its own repository
  case; `TestLogin` is only a container when it has subtests.
- **Signature** = `"<package import path>::<full test path>"` — param-free,
  case-sensitive, byte-stable across releases (golden-tested).
- **Suite path** = module-relative package path, one segment per element.
- Dynamic subtest names (`fmt.Sprintf` with counters, timestamps, random
  values) are an anti-pattern for reporting: Tiden's live documentation
  creates a permanent repository case per identity and never deletes on
  absence.

## Development

```bash
go test ./...
go vet ./...
gofmt -l .
```

Golden fixtures under `pkg/gotest/testdata/` are captured from real
`go test -json` runs — regenerate with `pkg/gotest/testdata/gen.sh`.
