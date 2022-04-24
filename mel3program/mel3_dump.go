package mel3program

import (
	"fmt"
)

// Mel dump, it prints out the object
func (prog *Mel3_object) MelDump() {

	if prog != nil {
		impl := prog.Implementation
		if impl != nil {
			startprog := prog.StartProgram
			if startprog != nil {
				if dump_engine(impl, startprog, 0) != nil {
					fmt.Printf("Dump Object Failed\n")
				}
			} else {
				fmt.Printf("Uninitializated Object Program\n")
			}
		} else {
			fmt.Printf("Uninitializated Object Implementation\n")
		}
	} else {
		fmt.Printf("Uninitializated Object Program\n")
	}

	//	fmt.Println(count_program(prog.Start_program))
}

// Dump engine: it recurse over the program and show it
func dump_engine(implementation map[uint16]*Mel3_implementation, program *Mel3_program, level int) error {

	programid := program.ProgramID
	libraryid := program.LibraryID

	isfunctional := true

	if len(implementation[libraryid].NonVariadicArgs[programid]) == 0 && !implementation[libraryid].IsVariadic[programid] {
		isfunctional = false
	}

	if isfunctional {
		for i := 0; i < level; i++ {
			fmt.Printf("\t")
		}
		fmt.Printf("%s.%s\n", implementation[libraryid].Implname, implementation[libraryid].ProgramNames[programid])
		for i := range program.NextPrograms {
			if dump_engine(implementation, program.NextPrograms[i], level+1) != nil {
			}
		}
	} else {
		for i := 0; i < level; i++ {
			fmt.Printf("\t")
		}
		fmt.Printf("%s.%s -> %s\n", implementation[libraryid].Implname, implementation[libraryid].ProgramNames[programid], program.ProgramValue)
	}

	return nil
}
