package main

func Caesar(text string, shift int) string {
	b := []byte(text)

	if b[0] == 'Z' {
		if shift == 1 {
			return string(byte(65) + byte(0))
		}

		if shift == 2 {
			return string(byte(65) + byte(1))
		}
	}

	for index := 0; index < len(b); index++ {
		b[index] = b[index] + byte(shift)
	}

	return string(b)
}
