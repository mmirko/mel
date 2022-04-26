//go:build DEBUG
// +build DEBUG

package mel3program

import (
	"fmt"
	mel "github.com/mmirko/mel"
)

// The Mel3 registration program every new data struct has to do it (This is not an entry point !! Not MelInit !! )
func (obj *Mel3Object) Mel3Init(implementation map[uint16]*Mel3Implementation, ep *mel.EvolutionParameters) {
	obj.Implementation = implementation

	for _, impl := range implementation {

		impl.Signatures = make(map[uint16]string)

		// Compute signatures
		for programid, pname := range impl.ProgramNames {
			signature := pname + "("

			atleastone := false
			for i, arg := range impl.NonVariadicArgs[programid] {
				atleastone = true
				if i != 0 {
					signature += ","
				}
				signature += implementation[arg.LibraryID].ImplName + "." + implementation[arg.LibraryID].TypeNames[arg.TypeID]
			}

			if impl.IsVariadic[programid] {
				if atleastone {
					signature += ","
				}
				// TODO check
				signature += implementation[impl.VariadicType[programid].LibraryID].ImplName + "." + impl.TypeNames[impl.VariadicType[programid].TypeID] + ",..."
			}

			signature += ")("
			for i, arg := range impl.ProgramTypes[programid] {
				if i != 0 {
					signature += ","
				}
				signature += implementation[arg.LibraryID].ImplName + "." + implementation[arg.LibraryID].TypeNames[arg.TypeID]
			}

			signature += ")"

			impl.Signatures[programid] = signature

			fmt.Println("Signatures:", signature)
		}
	}
}
