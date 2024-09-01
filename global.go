package lorem

var globalLorem = New()

// Generate a word in a specfied range of letters.
func Word(min, max int) string {
	return globalLorem.Word(min, max)
}

// Generate a sentence with a specified range of words.
func Sentence(min, max int) string {
	return globalLorem.Sentence(min, max)
}

func Paragraph(min, max int) string {
	return globalLorem.Paragraph(min, max)
}

// Generate a random URL
func Url() string {
	return "http://" + Host() + "/" + Word(2, 5)
}

// Host
func Host() string {
	return Word(2, 4) + "." + Word(2, 3)
}

// Email
func Email() string {
	return Word(4, 10) + `@` + Host()
}

// UUID
func UUID() string {
	return globalLorem.UUID()
}
