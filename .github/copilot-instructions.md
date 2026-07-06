# Copilot Instructions for goark/koyomi

## Project Purpose

`koyomi` is a Go library for Japanese calendar information.
It provides:

- calendar event retrieval from NAOJ calendar data via Google Calendar iCal feeds,
- date utilities (`value` package),
- Julian Day Number helpers (`jdn` package),
- zodiac helpers (`zodiac` package).

## Design Principles

- Keep exported API compatibility unless a breaking change is explicitly requested.
- Keep behavior deterministic for date/time conversion and event ordering.
- Prefer minimal, focused edits over broad refactors.
- Avoid unnecessary new dependencies.

## Behavior That Should Stay Stable

- Calendar ID mapping in `calendarid.go` should remain stable unless source calendars change.
- `CalendarID.URL()` should keep returning the public iCal URL form based on `url.PathEscape`.
- `Source.Get()` / `getFrom()` should continue returning wrapped errors with useful context keys (`cid`, `start`, `end`).
- JST-based handling in `value` package should remain explicit and consistent (`value.JST`).
- Era boundary behavior in `value/era.go` must not change without tests and documentation updates.

## Error Handling

- Use `github.com/goark/errs` for internal wrapping and context.
- Prefer `errs.Wrap` and `errs.Join` over losing context.
- Keep existing sentinel error behavior stable (`ErrNoData`, `ErrNullPointer`, `ErrInvalidRecord`).

## Coding Style

- Write idiomatic Go with straightforward control flow.
- Keep source code comments concise and in English.
- Follow existing package split and naming conventions.

## Testing and Validation

Run these after relevant changes:

```bash
go test ./...
go vet ./...
```

For vulnerability checks:

```bash
go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

If Go files changed, keep formatting clean:

```bash
gofmt -w <changed-go-files>
```

Task command notes:

- `task -f` is the force execution flag.
- Use `task -t Taskfile.yml` when explicitly specifying a Taskfile path.

## Documentation

- Keep `README.md` examples runnable and aligned with current APIs.
- Update docs when changing behavior, supported data, or outputs.
- Keep explanations concise and practical.

## GitHub Actions Guidance

Current workflows are:

- `ci.yml`
- `codeql.yml`

When editing workflows:

- Prefer stable major action versions (for example `@v4`, `@v5`, `@v6`) over `@main`.
- Keep permissions explicit and minimal.
- Avoid duplicate checks across workflows unless intentional.
- Preserve branch trigger intent (`master` push/PR) unless requested otherwise.

## Pull Request Expectations

- Summarize what changed and why.
- Mention validation commands that were run.
- Call out residual risks or follow-up tasks.
