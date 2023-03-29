//go:build !DEBUG
// +build !DEBUG

package mel3program

import (
	"errors"

	"github.com/mmirko/mel/pkg/mel"
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

func LibsCheckAndRequirements(libs []string) ([]string, error) {

	ExistingLibs := make(map[string]struct{})
	ExistingLibs["m3uint"] = struct{}{}
	ExistingLibs["m3uintcmp"] = struct{}{}
	ExistingLibs["m3number"] = struct{}{}
	ExistingLibs["m3bool"] = struct{}{}
	ExistingLibs["m3boolcmp"] = struct{}{}
	ExistingLibs["m3statements"] = struct{}{}
	ExistingLibs["m3dates"] = struct{}{}
	ExistingLibs["envfloat"] = struct{}{}

	CheckedLibs := make(map[string]struct{})

	for _, lib := range libs {
		if _, ok := ExistingLibs[lib]; !ok {
			return nil, errors.New("Unknown library: " + lib)
		} else {
			CheckedLibs[lib] = struct{}{}
			switch lib {
			case "m3uint":
			case "m3uintcmp":
				CheckedLibs["m3uint"] = struct{}{}
			case "m3number":
			case "m3bool":
			case "m3boolcmp":
				CheckedLibs["m3bool"] = struct{}{}
			case "m3statements":
			case "m3dates":
				CheckedLibs["m3uint"] = struct{}{}
			case "envfloat":
				CheckedLibs["m3uint"] = struct{}{}
				CheckedLibs["m3uintcmp"] = struct{}{}
				CheckedLibs["m3number"] = struct{}{}
				CheckedLibs["m3bool"] = struct{}{}
				CheckedLibs["m3boolcmp"] = struct{}{}
				CheckedLibs["m3statements"] = struct{}{}
			}
		}
	}

	// Convert map to slice
	libS := make([]string, 0, len(CheckedLibs))
	for lib := range CheckedLibs {
		libS = append(libS, lib)
	}

	return libS, nil
}
