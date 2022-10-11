package main

import (
	"sort"
	"strings"
)

//Meeting on codewars
//John has invited some friends. His list is:
//s = "Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill";
//So the result of function meeting(s) will be:
//"(CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)"
type Person struct {
	LastName  string
	FirstName string
}

func Meeting(s string) string {
	persons := []Person{}

	names := strings.Split(s, ";")
	for _, name := range names {
		full_name := strings.Split(name, ":")

		p := Person{strings.ToUpper(full_name[1]), strings.ToUpper(full_name[0])}
		persons = append(persons, p)
	}

	sort.Slice(persons, func(i, j int) bool {
		if strings.EqualFold(persons[i].LastName, persons[j].LastName) {
			if strings.Compare(persons[i].FirstName, persons[j].FirstName) >= 1 {
				return false
			} else {
				return true
			}
		}
		if strings.Compare(persons[i].LastName, persons[j].LastName) >= 1 {
			return false
		} else {
			return true
		}

	})
	result := ""
	for _, person := range persons {
		result += "(" + person.LastName + ", " + person.FirstName + ")"
	}
	return result
}

func Meeting1(s string) string {
	sw := strings.Split(strings.ToUpper(s), ";")
	res := []string{}
	for _, pn := range sw {
		a := strings.Split(pn, ":")
		s := "(" + a[1] + ", " + a[0] + ")"
		res = append(res, s)
	}
	sort.Strings(res)
	return strings.Join(res, "")
}
