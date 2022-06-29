package exp

import (
	"bytes"
	"fmt"
)

const (
	undefined int = iota
	line1
	indentAfterLine1
	keyWordIfStart
	keyWordIfEnd
	indentAfterIf
	expOfIf
	lBrace
	lineAfterLBrace
	indentAfterLine2
	keyWordReturn
	indentAfterReturn
	expOfReturn
	lineAfteKeyWordReturn
	indentAfterLine3
	rBrace
	line4
)

var returnInBytes = []byte{'r', 'e', 't', 'u', 'r', 'n'}

/*
given:
if err != nil { return xxxx }
output:
if err != nil {	return xxxx }
*/
func ApplyIf(src []byte) []byte {
	// { return xxxx }
	status := undefined
	start := 0
loop:
	for i := 0; i < len(src); {
		v := src[i]
		switch status {
		case undefined:
			start = 0
			if v == '\n' {
				status = line1
			} else {
				status = undefined
			}
		case line1, indentAfterLine1:
			if isIndent(v) {
				status = indentAfterLine1
			} else if src[i] == 'i' && i+1 < len(src) && src[i+1] == 'f' {
				status = keyWordIfStart
			} else {
				status = undefined
			}
		case keyWordIfStart:
			if src[i] == 'f' {
				status = keyWordIfEnd
			} else {
				status = undefined
			}
		case keyWordIfEnd:
			if isIndent(v) {
				status = indentAfterIf
			} else {
				status = undefined
			}
		case indentAfterIf:
			if isIndent(v) {
				status = indentAfterIf
			} else if v == '\n' {
				status = undefined
			} else if v == '{' {
				status = lBrace
			} else {
				status = expOfIf
			}
		case expOfIf:
			// no need to consider below since we only fold braces when there is '\n' following '}'
			// this way we can't fail fast, but this kinds of code are really rare to see
			// if nil == func(i int) int {return i} {...}
			// if struct{}{} == struct{}{} {...}
			if v == '\n' || v == '}' {
				status = undefined
			} else if v == '{' {
				status = lBrace
			} else {
				status = expOfIf
			}
		case lBrace:
			if v == '\n' {
				start = i // record i for trim
				status = lineAfterLBrace
			} else {
				status = undefined
			}
		case lineAfterLBrace, indentAfterLine2:
			if isIndent(v) {
				status = indentAfterLine2
			} else if v == 'r' && i+6 < len(src) && bytes.Compare(returnInBytes, src[i:i+6]) == 0 {
				i = i + 6
				status = keyWordReturn
				continue loop
			} else {
				status = undefined
			}
		case keyWordReturn:
			if v == '\n' {
				status = lineAfteKeyWordReturn
			} else if isIndent(v) {
				status = indentAfterReturn
			} else {
				status = undefined
			}
		case indentAfterReturn, expOfReturn:
			if v == '{' || v == '/' {
				status = undefined
			} else if v == '\n' {
				status = lineAfteKeyWordReturn
			} else if v == '}' {
				status = rBrace
			} else {
				status = expOfReturn
			}
		case lineAfteKeyWordReturn, indentAfterLine3:
			if isIndent(v) {
				status = indentAfterLine3
			} else if v == '}' {
				status = rBrace
			} else {
				status = undefined
			}
		case rBrace:
			// todo handle { return    errors.WithMessage(err,"{sss}"   )   }
			runei := rune(v)
			fmt.Println(runei)
			runei1 := rune(src[i+1])
			fmt.Println(runei1)
			if v == '\n' {
				src[start] = ' '
				for k := start + 1; k < i; k++ {
					if isIndent(src[k]) {
						src[k] = '\u0000'
					} else {
						break
					}
				}
				end := i - 2
				for j := end; j > 0; j-- {
					if isIndent(src[j]) {
						src[j] = '\u0000'
					} else if src[j] == '\n' {
						src[j] = ' '
					} else {
						break
					}
				}
				status = 0
			} else {
				status = undefined
			}
		default:
			status = undefined
		}
		i++
	}
	ret := make([]byte, 0, len(src))
	for _, v := range src {
		if v != '\u0000' {
			ret = append(ret, v)
		}
	}
	return ret
}

func isIndent(v byte) bool {
	if v == '\t' || v == ' ' {
		return true
	} else {
		return false
	}
}
