package partitions

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	mel "github.com/mmirko/mel"
)

func TestPartitionGenerator(t *testing.T) {

	fmt.Println("---- Test: Partition generator ----")

	// Random seed based on seconds since epoch
	rand.Seed(int64(time.Now().Second()))

	ep := new(mel.EvolutionParameters)

	ep.SetValue("setlength", "100")

	for i := 0; i < 20; i++ {
		cTest := Generate(ep)
		fmt.Println("Generated: ")
		fmt.Println("[", cTest, "]")
	}
	fmt.Println("---- End test: Partition generator ----")

}
