package board

// BoardRow ist eine Zeile des Spielbretts.
// Wir verwenden Strings, um die Spielfiguren zu repräsentieren.
// Daher ist BoardRow ein Slice von Strings.
type BoardRow []string

// MakeBoardRow erzeugt eine Zeile des Spielbretts.
// Die Funktion erwartet die Länge der Zeile und den String,
// mit dem die Zeile gefüllt werden soll.
func MakeBoardRow(length int, fill string) BoardRow {
	row := make(BoardRow, length)
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Slice mit dem gewünschten String zu füllen.
	*/
	// tag::solution[]
	for i := range row {
		row[i] = fill
	}
	// end::solution[]
	return row
}

// MakeEmptyBoardRow erzeugt eine leere Zeile des Spielbretts.
// Die Funktion erwartet die Länge der Zeile
// und füllt die Zeile mit Leerzeichen.
func MakeEmptyBoardRow(length int) BoardRow {
	/* Hinweis:
	   Verwenden Sie die Funktion MakeBoardRow.
	*/
	// tag::solution[]
	return MakeBoardRow(length, " ")
	// end::solution[]
	// iftask: return []string{}
}
