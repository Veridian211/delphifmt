# Formatter — Implementierung
 
## Scope
 
- Der Formatter verändert ausschließlich Whitespace — keine strukturellen Änderungen am Code.
- Idempotenz: zweifaches Formatieren erzeugt identisches Ergebnis.
---
 
## Traversierung
 
- Zentrale `visit(node Node)`-Methode mit Dispatch per Typ-Switch.
- Kein formales Visitor-Interface — der Typ-Switch ist der pragmatische Go-Ansatz.
- Typspezifische Logik steht einmal im Typ-Switch, nicht verstreut über viele Methoden.
```go
type Formatter struct {
    buf   strings.Builder
    depth int
    opts  Options
}
 
func (f *Formatter) visit(node Node) {
    switch n := node.(type) {
    case *IfStmt:
        f.visitIfStmt(n)
    case *AssignStmt:
        f.visitAssignStmt(n)
    case *CallExpr:
        f.visitCallExpr(n)
    }
}
 
func (f *Formatter) visitIfStmt(n *IfStmt) {
    f.write(f.keyword("if") + " ")
    f.visit(n.Cond)
    f.write(" " + f.keyword("then"))
    f.depth++
    f.visit(n.Then)
    f.depth--
    if n.Else != nil {
        f.write(f.keyword("else"))
        f.depth++
        f.visit(n.Else)
        f.depth--
    }
}
 
func (f *Formatter) visitAssignStmt(n *AssignStmt) {
    f.visit(n.Left)
    f.write(" := ")
    f.visit(n.Right)
    f.write(";")
}
```
 
---
 
## Zeilenumbruch
 
- Jeder Ausdruck wird zuerst einzeilig versucht (temporärer Formatter ohne Seiteneffekte).
- Überschreitet er die Zeilenlänge, wird der äußerste Ausdruck umgebrochen.
- Innere Ausdrücke werden rekursiv nach demselben Prinzip behandelt.
---
 
## Formatierungsoptionen
 
### Nicht konfigurierbar (immer angewendet)
 
- Trailing Whitespace entfernen
- Leerzeile am Dateiende sicherstellen
- Leerzeichen um Zuweisungsoperatoren (`:=`, `=` in Deklarationen)
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
 
#### `case of` Beispiel
 
```delphi
case Color of
  clRed: begin
    Foo;
    Bar;
  end;
  clBlue: begin
    Baz;
  end;
  else begin
    Default;
  end;
end;
```
 
### Bewusst nicht unterstützt
 
- Identifier-Case-Normalisierung — zu riskant, Semantik bleibt erhalten
- Automatische Sortierung von `uses`-Klauseln
- Alignment von `:=` in Blöcken — führt zu riesigen Diffs
 
