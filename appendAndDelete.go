package main

func appendAndDelete(s string, t string, k int32) string {
	b := int32(len(s))
	for i := 0; i < len(s); i++ {
		if s == t {
			return "Yes"
		} else if k >= 0 {
			return "Yes"
		} else if k > b {
			return "No"
		} else if k == 2 {
			return "No"
		}
	}
	return "No"
}
