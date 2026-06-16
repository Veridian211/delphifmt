# delphifmt — Agent Summary

## What this project is

A CLI formatter and linter for modern Delphi (XE+), implemented in Go. Input is always files, never stdin.

## CLI

```
delphifmt fmt   [--write] [--check] <files>
delphifmt lint  [--fix] [--format=rdjson] <files>
```

- `--write` formats in-place; `--check` exits with code 1 if changes would be made (CI use).
- Glob patterns are supported (`**/*.pas`).

## Architecture

```
Source → Lexer → Token-Stream → Parser → AST → Formatter → Output
                                          ↓
                                        Linter → Diagnostics
```

- **Lexer** — produces a token stream; compiler directives and comments are preserved as opaque tokens; aborts on invalid input. More information in [/.clanker/lexer.md](/.clanker/lexer.md).
- **Parser** — builds an AST; `{$IFDEF}` branches are modelled as alternative subtrees; comments attach as trivia; collects errors best-effort. More information in [/.clanker/parser.md](/.clanker/parser.md).
- **Formatter** — type-switch AST traversal; only modifies whitespace; must be idempotent. More information in [/.clanker/formatter.md](/.clanker/formatter.md).
- **Linter** — independent AST traversal; emits diagnostics (line, col, message); never changes code unless `--fix`. More information in [/.clanker/linter.md](/.clanker/linter.md).

## Formatting rules

Always applied: trailing whitespace removal, trailing newline, spaces around `:=` / `=`.

Configurable via `.delphifmt.yaml` (searched upward from the file, stopping at `.git`):

| Option | Default |
|---|---|
| `indent` | 2 spaces |
| `keyword_case` | `lowercase` |
| `line_length` | 100 |
| `begin_position` | end of line |
| `blank_lines_between_methods` | 1 |
| `uses` style | `one_per_line` |

Not supported by design: identifier-case normalisation, `uses` sorting, `:=` alignment.

## Language scope

Generics, inline variables, anonymous methods, and attributes are fully formatted. `asm` blocks are treated as opaque — content is not parsed, only indentation is applied.

## Encoding

UTF-8 only. BOM is tolerated but never written.

## Testing

Golden-file tests: each case has `input.pas` + `expected.pas`. Idempotency is explicitly verified by running the formatter a second time on `expected.pas`.

## Project layout

```
main.go
token/
lexer/
ast/
parser/
formatter/   (formatter.go, options.go, stmt.go, expr.go, decl.go, comment.go)
```

## Running the project

Run main.go with `make run`.
Run tests with `make test`.