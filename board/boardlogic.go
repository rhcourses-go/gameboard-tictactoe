package board

// BoardRowContainsOnly prüft, ob die angegebene Zeile nur aus dem übergebenen String besteht.
func BoardRowContainsOnly(b Board, row int, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion GetRow.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	// tag::solution[]
	return RowContainsOnly(GetRow(b, row), s)
	// end::solution[]
	// iftask: return false
}

// BoardColumnContainsOnly prüft, ob die angegebene Spalte nur aus dem übergebenen String besteht.
func BoardColumnContainsOnly(b Board, column int, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion GetColumn.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	// tag::solution[]
	return RowContainsOnly(GetColumn(b, column), s)
	// end::solution[]
	// iftask: return false
}

// BoardDiagRightContainsOnly prüft, ob die Diagonale von links
// oben nach rechts unten nur aus dem übergebenen String besteht.
func BoardDiagRightContainsOnly(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion GetDiagRight.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	// tag::solution[]
	return RowContainsOnly(GetDiagRight(b), s)
	// end::solution[]
	// iftask: return false
}

// BoardDiagLeftContainsOnly prüft, ob die Diagonale von rechts
// oben nach links unten nur aus dem übergebenen String besteht.
func BoardDiagLeftContainsOnly(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie die Funktion GetDiagLeft.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	// tag::solution[]
	return RowContainsOnly(GetDiagLeft(b), s)
	// end::solution[]
	// iftask: return false
}

// BoardContainsOnly prüft, ob das Spielbrett nur aus dem übergebenen String besteht.
func BoardContainsOnly(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Verwenden Sie die Funktion RowContainsOnly.
	*/
	// tag::solution[]
	for row := range b {
		if !BoardRowContainsOnly(b, row, s) {
			return false
		}
	}
	// end::solution[]
	return true
}

// BoardDoesNotContain prüft, ob das Spielbrett den übergebenen String nicht enthält.
// D.h. es darf kein Feld auf dem Spielbrett den übergebenen String enthalten.
func BoardDoesNotContain(b Board, s string) bool {
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die Zeilen zu durchlaufen.
	   Verwenden Sie eine innere for-Schleife, um die Spalten zu durchlaufen.
	*/
	// tag::solution[]
	for _, row := range b {
		for _, field := range row {
			if field == s {
				return false
			}
		}
	}
	// end::solution[]
	return true
}
