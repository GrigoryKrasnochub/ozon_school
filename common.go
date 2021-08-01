package ozon

// Т.к. строка неизменяемая, получить O(1) по памяти работая со строкой невозможно,
// в принципе, можно примерно тогоже добиться работая с byte, но не все символы кодируются 1 байтом
func cutSymbolDuplication(symToCut rune, sourceString []rune) []rune {
	newString := sourceString[:0]
	var prevSym rune
	for i, x := range sourceString {
		if !((i == 0 || prevSym == x) && x == symToCut) {
			newString = append(newString, x)
		}
		prevSym = x
	}
	return newString
}
