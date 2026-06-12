# Lexer — Implementierung
 
## Token-Typen
 
- String-Literals: `'...'` und `#13`-Escape-Sequenzen
- Numerische Literale: Dezimal, Hex (`$FF`), Binär (`%1010`)
- Keywords, Identifier, Operatoren, Interpunktion
---
 
## Compiler-Direktiven
 
- Werden als `TokenDirective` erfasst.
- Ihr Inhalt wird **ungeparst** durchgereicht (opaker Block).
- Der Formatter übernimmt die Einrückung, verändert den Inhalt nicht.
- `{$IFDEF}`-Branches werden als **alternative Subtrees** im AST modelliert, nicht ignoriert — Voraussetzung für Idempotenz.
**Begründung:** Direktiven können syntaktisch unvollständige Code-Fragmente enthalten (z.B. `{$IFDEF}` mitten in einem Block). Ein Parser, der alle Branches gleichzeitig parsen will, scheitert. Durch opake Behandlung bleibt der Formatter robust gegenüber beliebigen Direktiv-Kombinationen.
 
---
 
## Kommentare
 
- Alle drei Kommentar-Typen werden als Token erhalten: `//`, `{ }`, `(* *)`
- Kommentare bleiben im Token-Stream (werden nicht verworfen).
**Begründung:** Für einen Formatter ist der Erhalt von Kommentaren Pflicht. Verlorene Kommentare wären ein kritischer Fehler.
 
---
 
## Fehlerverhalten
 
- Bei ungültigem Input: sofort abbrechen und Fehler melden.
 
