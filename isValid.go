package main

func main() {

}

func isValid(s string) bool {
	var braces []string

	for _, v := range s {
		t := string(v)

		if t == "(" || t == "{" || t == "[" {
			braces = append(braces, t)
		} else {
			if len(braces) == 0 {
				return false
			}
			endingBrace := braces[len(braces)-1]
			if t == ")" && endingBrace == "(" {
				braces = braces[:len(braces)-1]
			} else if t == "}" && endingBrace == "{" {
				braces = braces[:len(braces)-1]
			} else if t == "]" && endingBrace == "[" {
				braces = braces[:len(braces)-1]
			}

		}
		if len(braces) == 0 {
			return true
		}
	}
	return false
}
