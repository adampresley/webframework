package webpage

/*
NewGoLayoutFromFile creates a new Mustache-based layout from a file
*/
func NewGoLayoutFromFile(fileName string) (ILayout, error) {
	result := &GoLayout{}
	err := result.LoadLayoutFile(fileName)

	return result, err
}

/*
NewGoLayoutFromString creates a new Mustache-based layout from a byte array
*/
func NewGoLayoutFromString(layout []byte) (ILayout, error) {
	result := &GoLayout{}
	err := result.LoadLayoutString(layout)

	return result, err
}

/*
NewMustacheLayoutFromFile creates a new Mustache-based layout from a file
*/
func NewMustacheLayoutFromFile(fileName string) (ILayout, error) {
	result := &MustacheLayout{}
	err := result.LoadLayoutFile(fileName)

	return result, err
}

/*
NewMustacheLayoutFromString creates a new Mustache-based layout from a byte array
*/
func NewMustacheLayoutFromString(layout []byte) (ILayout, error) {
	result := &MustacheLayout{}
	err := result.LoadLayoutString(layout)

	return result, err
}
