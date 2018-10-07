package util

// ParseContinuousColorHexString ...
func ParseContinuousColorHexString(colorStr string) []string {
	if colorStr == "" {
		return nil
	}
	slice := []byte(colorStr)
	const colorlen = 6
	var colors []string
	size := len(slice)

	for i := 0; i < size; i += colorlen {
		if i%colorlen == 0 {
			x := (i + colorlen)
			if x > size {
				break
			}
			colors = append(colors, "#"+string(slice[i:x]))
		}
	}

	return colors
}
