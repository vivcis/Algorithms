package main

//Complementary DNA
//codewars
func DNAStrand(dna string) string {
	if dna == "" {
		return ""
	}
	result := ""
	for _, t := range dna {
		s := string(t)

		if s == "A" {
			result += "T"
		} else if s == "T" {
			result += "A"
		} else if s == "C" {
			result += "G"
		} else if s == "G" {
			result += "C"
		} else {
			result += s
		}
	}

	return result
}
