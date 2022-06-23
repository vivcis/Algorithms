package main

//function that counts valleys
func countingValleys(steps int32, path string) int32 {
	var valleys int32
	var level int32
	for i := 0; i < len(path); i++ {
		if path[i] == 'U' {
			level++
		} else {
			level--
		}
		if level == 0 && path[i] == 'U' {
			valleys++
		}
	}
	return valleys
}
