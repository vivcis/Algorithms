package main

import (
	"sort"
	"strings"
)

func main() {

}

func superReducedString(s string) string {
	x := strings.Split(s, "")
	//var d []string
	d := []string{}
	m := map[string]int{}
	//var m map[string]int
	x = append(x, " ")
	for i := 0; i < len(x); i++ {
		if x[i] == " " {
			break
		}
		if _, ok := m[x[i]]; !ok {
			m[x[i]] = 1
		} else {
			delete(m, x[i])
		}
	}
	for k, _ := range m {
		d = append(d, k)
	}
	sort.Strings(d)
	if len(m) == 0 {
		return "Empty string"
	}
	return strings.Join(d, "")
}

//OR

func superReducedString1(s string) string {
	// Write your code here

	//divide the letters or split the strings
	//iterate through them
	//set counter to 0
	//check for similarities btwn the letters in the string
	//if similar to a previous letter append to an empty slice
	//if none is similar to previous letter return
	//if empty return empty string

	x := strings.Split(s, "")
	//fmt.Println(x[2])
	// [a  a  b  c  d]
	//[a  b  c  d]
	var sim []string
	var notSim []string
	for i := 0; i < len(x); i++ {
		for j := 1; j < len(x)-1; j++ {
			if x[i] == x[j] {
				sim = append(sim, x[j])
			} else if s[i] != s[j] {
				// nosim= [a, b, c]
				notSim = append(notSim, x[i])
			}
		}
	}
	if len(notSim) != 0 {
		return strings.Join(sim, "")
	}
	return "Empty string"
}
