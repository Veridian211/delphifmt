# Delphi Formatter — Designentscheidungen
 
## Zielsprache & Tooling
 
- **Implementierungssprache:** Go
- **Ziel-Dialekt:** Modernes Delphi (XE+)
- **Nutzung:** CLI-Tool (Datei-Input)
 
## Sprachumfang
 
- Generics, Inline-Variablen, Anonyme Methoden und Attributes werden vollständig formatiert.
- `asm`-Blöcke werden als opaker Block behandelt: Inhalt ungeparst, Formatter übernimmt nur Einrückung.
 
## Encoding
 
- UTF-8 only. BOM wird toleriert aber nicht hinzugefügt.
 
## Formatierungsoptionen
 
### Nicht konfigurierbar (immer angewendet)
 
- Trailing Whitespace entfernen
- Leerzeile am Dateiende sicherstellen
- Leerzeichen um Zuweisungsoperatoren (`:=`, `=` in Deklarationen)
- Idempotenz: zweifaches Formatieren erzeugt identisches Ergebnis
### Konfigurierbar
 
| Option | Default |
|---|---|
| Einrückung | 2 Spaces |
| Keyword-Case | `lowercase` |
| Zeilenlänge | 100 |
| `begin`-Position | Ende der Zeile |
| Leerzeilen zwischen Methoden | 1 |
| `case`-Labels | eingerückt, `else` auf gleicher Ebene wie Labels |
| `uses`-Stil | `one_per_line` (`one_per_line` oder `single_line`) |
 
### Bewusst nicht unterstützt
 
- Identifier-Case-Normalisierung — zu riskant, Semantik bleibt erhalten
- Automatische Sortierung von `uses`-Klauseln
- Alignment von `:=` in Blöcken — führt zu riesigen Diffs
 
## CLI-Interface
 
Es gibt eine einzelne Executable, die Formatter und Linter als getrennte Subkommandos anbietet:
 
```
delphifmt fmt   [--write] [--check] <files>
delphifmt lint  [--fix] [--format=rdjson] <files>
```
 
Nur Datei-Input wird unterstützt, kein Stdin.
 
- `--write` — Formatiert Dateien in-place
- `--check` — Prüft nur, Exit-Code 1 bei Unterschied (für CI)
- Mehrere Dateien und Glob-Patterns werden unterstützt (`**/*.pas`)
### Config-Suche
 
- Der Formatter sucht ausgehend vom Verzeichnis der zu formatierenden Datei aufwärts nach `.delphifmt.yaml`.
- Root ist das Verzeichnis mit `.git` (Verzeichnis, nicht Datei) — die Suche endet dort.
 
### Konfigurationsdatei (`.delphifmt.yaml`)
 
```yaml
indent: 2
line_length: 100
keyword_case: lowercase
begin_position: end_of_line
blank_lines_between_methods: 1
indent_char: space
end_of_line: lf
 
include:
  - "**/*.pas"
  - "**/*.dpr"
  - "**/*.dpk"
 
exclude:
  - "vendor/**"
  - "generated/**"
```
 
## Architektur
 
Pipeline: `Source` → `Lexer` → `Token-Stream` → `Parser` → `AST` → `Formatter` → `Output`
 
Linter traversiert denselben AST unabhängig: `AST` → `Linter` → `Diagnostics`
 
### Lexer
 
Wandelt Quelltext in einen Token-Stream um. Compiler-Direktiven und Kommentare werden als eigene Token-Typen erhalten und opak durchgereicht. Bei ungültigem Input: sofort abbrechen.
 
Details: siehe `LEXER.md`
 
### Parser
 
Baut den AST aus dem Token-Stream. `{$IFDEF}`-Branches werden als alternative Subtrees modelliert. Kommentare werden als Trivia an Knoten gehängt. Bei Fehlern: Best-Effort, Fehlerliste sammeln.
 
Details: siehe `PARSER.md`
 
### Formatter
 
Traversiert den AST per Typ-Switch und schreibt formatierten Output. Verändert ausschließlich Whitespace.
 
Details: siehe `FORMATTER.md`
 
### Linter
 
Separate AST-Traversierung, gibt Diagnostics aus (Zeile, Spalte, Meldung), verändert keinen Code. Mit `--fix` automatische Korrektur bestimmter Regeln.
 
Details: siehe `LINTER.md`
 
## Projektstruktur
 
```
delphi-fmt/
├── main.go
├── go.mod
├── token/
│   └── token.go
├── lexer/
│   ├── lexer.go
│   └── lexer_test.go
├── ast/
│   ├── ast.go
│   └── ast_test.go
├── parser/
│   ├── parser.go
│   └── parser_test.go
└── formatter/
    ├── formatter.go
    ├── options.go
    ├── stmt.go
    ├── expr.go
    ├── decl.go
    ├── comment.go
    ├── formatter_test.go
    └── formatter_internal_test.go
```
 
## Tests
 
### Golden Files
 
- Jeder Testfall besteht aus zwei Dateien: `input.pas` und `expected.pas`.
- Der Formatter wird auf `input.pas` ausgeführt und das Ergebnis mit `expected.pas` verglichen.
- Idempotenz wird explizit getestet: `expected.pas` durch den Formatter gejagt muss identisch bleiben.
```
tests/
  case_of/
    input.pas
    expected.pas
  ifdef_branch/
    input.pas
    expected.pas
```
 