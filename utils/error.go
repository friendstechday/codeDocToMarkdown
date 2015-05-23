package utils


func GeneralCheckError(e error) {
	if e != nil {
		panic(e)
	}
}