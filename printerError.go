package main

import "strconv"

func main() {

}

//In a factory a printer prints labels for boxes. For one kind of boxes the printer has to use colors which, for the sake
//of simplicity, are named with letters from a to m. The colors used by the printer are recorded in a control string.
//For example a "good" control string would be aaabbbbhaijjjm meaning that the printer used three times color a, four times color b,
//one time color h then one time color a...
//Sometimes there are problems: lack of colors, technical malfunction and a "bad" control string is produced
//e.g. aaaxbbbbyyhwawiwjjjwwm with letters not from a to m.
//The string has a length greater or equal to one and contains only letters from ato z.
//E.G s="aaabbbbhaijjjm"
//printer_error(s) => "0/14"
//s="aaaxbbbbyyhwawiwjjjwwm"
//printer_error(s) => "8/22"

func PrinterError(s string) string {

	totalBadColors := 0
	totalColors := 0

	for _, v := range s {
		if v < 'a' || v > 'm' {
			totalBadColors++
		}
		totalColors++
	}
	return strconv.Itoa(totalBadColors) + "/" + strconv.Itoa(totalColors)
}
