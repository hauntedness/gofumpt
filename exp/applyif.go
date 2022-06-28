package exp

/*
given:
if err != nil {
	return xxxx
}
output:
if err != nil {	return xxxx }
*/
func ApplyIf(src []byte) []byte {
	// { return xxxx }
	index := 0
	start := 0
	tobetrimed := []int{}
	for i, v := range src {
		switch index {
		case 0:
			if v == '{' {
				index++
			} else {
				index = 0
			}
		case 1:
			if v == '\n' {
				start = i
				index++
			} else {
				index = 0
			}
		case 2:
			if v == 'r' {
				index++
			} else if v == '\t' || v == ' ' {
				tobetrimed = append(tobetrimed, i)
				// do nothing
			} else {
				index = 0
			}
		case 3:
			if v == 'e' {
				index++
			} else {
				index = 0
			}
		case 4:
			if v == 't' {
				index++
			} else {
				index = 0
			}
		case 5:
			if v == 'u' {
				index++
			} else {
				index = 0
			}
		case 6:
			if v == 'r' {
				index++
			} else {
				index = 0
			}
		case 7:
			if v == 'n' {
				index++
			} else {
				index = 0
			}
		case 8:
			if v == '\t' || v == ' ' {
				index++
			} else {
				index = 0
			}
		case 9:
			// todo handle { return    errors.WithMessage(err,"{sss}"   )   }
			if v == '}' {
				src[start] = ' '
				src[i-1] = ' '
				for j := i - 1; j > 0; j-- {
					if src[j] == ' ' || src[j] == '\t' {
						src[j] = '\u0000'
					} else if src[j] == '\n' {
						src[j] = ' '
						break
					} else {
						break
					}
				}
				for _, k := range tobetrimed {
					src[k] = '\u0000'
				}
				index = 0
			} else if v == '{' {
				index = 0
			} else {
				// do nothing
			}
		default:
			index = 0
		}
		if index == 0 {
			tobetrimed = []int{}
		}
	}
	ret := make([]byte, 0, len(src))
	for _, v := range src {
		if v != '\u0000' {
			ret = append(ret, v)
		}
	}
	return ret
}
