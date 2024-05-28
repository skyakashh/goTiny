package hashing

var s string = "tata"
var s1 string = "novam"

func Hash() string {
	rotateleft()
	rotaterigth()
	res := s + s1
	return res

}

func rotateleft() {
	runes := []rune(s)
	n := len(s)
	runes = append(runes[n-1:], runes...)
	s = string(runes[0:n])
}

func rotaterigth() {
	runes := []rune(s1)
	runes = append(runes, runes[0:1]...)
	s1 = string(runes[1:])
}
