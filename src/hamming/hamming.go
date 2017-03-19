package hamming

import "errors"

func Distance(strand1, strand2 string) (int, error) {
	if len(strand1) != len(strand2) {
		return -1, errors.New("Starnds do not have the same size")
	}
	distance := 0
	for i := 0; i < len(strand1); i++ {
		if strand1[i] != strand2[i] {
			distance++
		}
	}
	return distance, nil
}
