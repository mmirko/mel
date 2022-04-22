package mel3program

import (
	//"fmt"
	"strings"
)

func MatchSignature(s1 string, s2 string, mode uint8) bool {
	//fmt.Println(s1, s2)

	s1part := strings.Split(s1, ")(")
	s2part := strings.Split(s2, ")(")

	dec1 := strings.Split(s1part[0], "(")
	dec2 := strings.Split(s2part[0], "(")

	//sec1 := strings.Split(s1part[1], ")")
	//sec2 := strings.Split(s2part[1], ")")

	name1 := dec1[0]
	name2 := dec2[0]

	sig1p1 := strings.Split(dec1[1], ",")
	sig2p1 := strings.Split(dec2[1], ",")

	//sig1p2 := sec1[0]
	//sig2p2 := sec2[0]

	if name1 != name2 {
		return false
	}

	shorter := sig1p1
	longer := sig2p1

	if len(sig1p1) > len(sig2p1) {
		shorter = sig2p1
		longer = sig1p1
	}

	var oldarg string
	startvariadic := false
	for i, arg1 := range longer {
		if i >= len(shorter) {
			if !startvariadic {
				if arg1 == "..." {
					break
				} else if len(longer) == i+2 && longer[i+1] == "..." {
					break
				} else {
					return false
				}
			} else {
				if arg1 != oldarg {
					return false
				}
			}
		} else {
			arg2 := shorter[i]
			if arg2 == "..." {
				startvariadic = true
				if arg1 != oldarg {
					return false
				}
			} else {
				if arg1 != arg2 {
					if arg1 == "..." {
						if arg2 != oldarg {
							return false
						}
					} else {
						return false
					}
				}
				oldarg = arg2
			}
		}
	}
	return true
}
