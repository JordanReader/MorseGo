package convert



var englishLetters = [...]string {"a", "b", "c", "d", "e"}
var morseLetters = [...]string {".-", "-...", ".. .", "-..", "."}

func Convert(letter string) string {


	for i := 0; i <= 4; i++ {

		if letter == englishLetters[i] {
			return morseLetters[i]
		} else {

		}
	}



	return letter

}