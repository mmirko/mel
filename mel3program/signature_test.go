package mel3program

import (
	"fmt"
	"testing"
)

func TestSignatures(t *testing.T) {

	fmt.Println("---- Test: Signatures ----")

	istrings1 := []string{"multi(int,float)()", "multi(int,float)()", "multi(int,int,...)()", "multi(int,int)()", "multi(int,int)()", "multi(int,...)()", "multi(int,int,int,int)()", "multi(int,...)()"}
	istrings2 := []string{"multi(int,float)()", "multi(int,int)()", "multi(int,int)()", "multi(int,float,...)()", "multi(int,...)()", "multi(int,int)()", "multi(int,...)()", "multi(int,float)()"}

	for i := 0; i < len(istrings1); i++ {
		fmt.Println("Testing "+istrings1[i]+" and "+istrings2[i]+":", MatchSignature(istrings1[i], istrings2[i], uint8(0)))
	}

	fmt.Println("---- End test ----")

}
