package algorithms

func IsPalindrome(wordOne string) bool {
	size := len(wordOne)

	for i := 0; i < size; i++ {
		if wordOne[i] != wordOne[size-i-1] {
			return false
		}
	}

	return true
}
