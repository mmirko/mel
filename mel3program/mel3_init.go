//go:build !DEBUG
// +build !DEBUG

package mel3program

import (
	mel "github.com/mmirko/mel"
)

// The Mel3 registration program every new data struct has to do it (This is not an entry point !! Not MelInit !! )
func (obj *Mel3Object) Mel3Init(c *mel.MelConfig, implementation map[uint16]*Mel3Implementation, creators map[uint16]Mel3VisitorCreator, ep *mel.EvolutionParameters) {
	obj.Config = c
	obj.Implementation = implementation
	obj.VisitorCreator = creators

	for _, impl := range implementation {

		if impl.Signatures == nil {

			impl.Signatures = make(map[uint16]string)

			// Compute signatures
			for programId, pname := range impl.ProgramNames {
				signature := pname + "("

				atLeastOne := false
				for i, arg := range impl.NonVariadicArgs[programId] {
					atLeastOne = true
					if i != 0 {
						signature += ","
					}
					signature += implementation[arg.LibraryID].ImplName + "." + implementation[arg.LibraryID].TypeNames[arg.TypeID]
				}

				if impl.IsVariadic[programId] {
					if atLeastOne {
						signature += ","
					}
					// TODO check
					signature += implementation[impl.VariadicType[programId].LibraryID].ImplName + "." + impl.TypeNames[impl.VariadicType[programId].TypeID] + ",..."
				}

				signature += ")("
				for i, arg := range impl.ProgramTypes[programId] {
					if i != 0 {
						signature += ","
					}
					signature += implementation[arg.LibraryID].ImplName + "." + implementation[arg.LibraryID].TypeNames[arg.TypeID]
				}

				signature += ")"

				impl.Signatures[programId] = signature
			}
		}
	}
}
