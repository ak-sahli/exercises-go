package raindrops

import "strconv"

var drops = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(number int) interface{} {
	var raindrop string
	for factor, sound := range drops {
		if hasFactor(number, factor) {
			raindrop += sound
		}
	}
	if len(raindrop) == 0 {
		raindrop = strconv.Itoa(number)
	}
	return raindrop
}

func hasFactor(number, factor int) bool {
	return number%factor == 0
}
