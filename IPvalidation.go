package main

import (
	"strconv"
	"strings"
)

//IP validation on codewars

//check for empty ip
//split with the dot
//check for length if its not upto 4, should return false
//range through, check if suffix n prefix is 0 should return false
//convert to int
//check if its less than 0 or greater than 256 should return false
//else, if all these other conditions ain't ment, should return true
func Is_valid_ip(ip string) bool {
	if ip == "" {
		return false
	}
	seg := strings.Split(ip, ".")
	if len(seg) != 4 {
		return false
	}

	for _, v := range seg {
		if strings.HasPrefix(v, "0") && !strings.HasSuffix(v, "0") {
			return false
		}
		num, err := strconv.Atoi(v)

		if err != nil {
			return false
		}
		if num < 0 || num > 256 {
			return false
		}
	}
	return true
}
