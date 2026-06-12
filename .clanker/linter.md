# Linter — Implementierung
 
## Konzept
 
- Separate Traversierung über den AST, unabhängig vom Formatter.
- Gibt Diagnostics aus (Zeile, Spalte, Meldung), verändert keinen Code.
- Mit `--fix` können bestimmte Regeln automatisch korrigiert werden.
---
 
## Output
 
- Standard: menschenlesbares Format (Zeile, Spalte, Meldung)
- Mit `--format=rdjson` oder `--format=rdjsonl`: reviewdog-kompatibles Format für CI-Pipelines
---
 
## Geplante Regeln
 
**Strukturell (mit `--fix` korrigierbar):**
- Fehlende `begin`/`end` bei einzeiligen `if`/`else`/`while`-Zweigen (dort ist `begin`/`end` optional)
- Fehlende oder überflüssige Klammern bei Methodenaufrufen ohne Parameter (`Foo` vs. `Foo()`)
**Namenskonventionen:**
- Typen ohne `T`-Prefix (`TMyClass`)
- Parameter ohne `A`-Prefix (`AValue`)
- Felder ohne `F`-Prefix (`FName`)
**Stilregeln:**
- Zu tiefe Verschachtelung (konfigurierbare Schwelle)
- Zu lange Funktionen (konfigurierbare Zeilenzahl)
- Magische Zahlen (Literale außerhalb von Konstanten)
- Leere `except`-Blöcke
---
 
## Abgrenzung zum Formatter
 
- Der Formatter korrigiert stumm und verändert nur Whitespace.
- Der Linter meldet Probleme und lässt den Entwickler entscheiden.
- Mit `--fix` überschneiden sich beide — strukturelle Korrekturen wie `begin`/`end` sind Linter-Aufgabe, nicht Formatter-Aufgabe.
 
