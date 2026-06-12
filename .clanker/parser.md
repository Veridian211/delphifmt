# Parser — Implementierung
 
## Sprachumfang
 
- Generics, Inline-Variablen, Anonyme Methoden und Attributes werden vollständig geparst.
- `asm`-Blöcke werden als opaker Block behandelt (analog zu Compiler-Direktiven): Inhalt wird ungeparst durchgereicht, der Formatter übernimmt nur die Einrückung.
---
 
## Fehlerverhalten
 
- Best-Effort — Fehlerliste sammeln, nicht sofort abbrechen. Der Parser versucht weiterzumachen um möglichst viele Fehler in einem Durchlauf zu melden.
### Fehlersynchronisation
 
Bei einem Syntaxfehler überspringt der Parser Token bis zu einem bekannten Sync-Token und macht dort weiter:
 
- `;` — Ende eines Statements
- `begin`, `end` — Blockgrenzen
- `procedure`, `function`, `type`, `var`, `const` — Deklarationsgrenzen
```go
func (p *Parser) synchronize() {
    for p.current().Type != TokenEOF {
        switch p.current().Type {
        case TokenSemicolon,
             TokenBegin, TokenEnd,
             TokenProcedure, TokenFunction,
             TokenType, TokenVar, TokenConst:
            return
        }
        p.advance()
    }
}
```
 
---
 
## AST
 
### Node
 
```go
// Gemeinsamer Typ-Anker für alle AST-Knoten.
type Node interface{}
 
// Marker-Interface für alle Statement-Knoten.
type Stmt interface{ Node }
 
// Marker-Interface für alle Ausdruck-Knoten.
type Expr interface{ Node }
 
// Position wird direkt in konkreten Structs gehalten wo sie gebraucht wird.
type Position struct {
    Line   int
    Column int
}
```
 
### Beispiel-Knoten
 
```go
type IfStmt struct {
    LeadingComments []Comment
    Cond            Expr
    Then            Stmt
    Else            Stmt
    TrailingComment *Comment
    Pos             Position
}
 
type BinaryExpr struct {
    Left  Expr
    Op    string
    Right Expr
    Pos   Position
}
 
type CallExpr struct {
    Name string
    Args []Expr
    Pos  Position
}
 
type Ident struct {
    Name string
    Pos  Position
}
 
type Literal struct {
    Value string
    Pos   Position
}
```
 
---
 
## `{$IFDEF}`-Branches
 
- Werden als **alternative Subtrees** im AST modelliert, nicht ignoriert.
- Voraussetzung für Idempotenz: der Formatter muss alle Branches kennen.
---
 
## Trivia (Kommentare)
 
Kommentare werden als Trivia an Knoten gehängt:
 
- `LeadingComments []Comment` — Kommentare vor dem Knoten
- `TrailingComment *Comment` — Kommentar in derselben Zeile nach dem Knoten
Die Zuordnung erfolgt über Positionsinformation (Zeilennummer) des Tokens.
 
 
