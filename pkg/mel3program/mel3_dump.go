package mel3program

import (
	"fmt"

	"github.com/mmirko/mel/pkg/mel"
)

// Mel dump, it prints out the object
func (prog *Mel3Object) MelDump(c *mel.DumpConfig) {

	if prog != nil {
		impl := prog.Implementation
		if impl != nil {
			startprog := prog.StartProgram
			if startprog != nil {
				if dumpEngine(c, impl, startprog, 0) != nil {
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
func dumpEngine(c *mel.DumpConfig, implementation map[uint16]*Mel3Implementation, program *Mel3Program, level int) error {

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
		if c != nil && c.Numeric {
			fmt.Printf("(%d %d)", libraryid, programid)
			if c.Types {
				fmt.Printf("->%s", fmt.Sprint(implementation[libraryid].ProgramTypes[programid]))
			}
			fmt.Println()
		} else {
			fmt.Printf("%s.%s\n", implementation[libraryid].ImplName, implementation[libraryid].ProgramNames[programid])
		}
		for i := range program.NextPrograms {
			if dumpEngine(c, implementation, program.NextPrograms[i], level+1) != nil {
				return fmt.Errorf("dump Engine Failed")
			}
		}
	} else {
		for i := 0; i < level; i++ {
			fmt.Printf("\t")
		}
		if c != nil && c.Numeric {
			fmt.Printf("(%d %d)", libraryid, programid)
			if c.Types {
				fmt.Printf("->%s", fmt.Sprint(implementation[libraryid].ProgramTypes[programid]))
			}
			fmt.Println()
		} else {
			fmt.Printf("%s.%s -> %s\n", implementation[libraryid].ImplName, implementation[libraryid].ProgramNames[programid], program.ProgramValue)
		}
	}

	return nil
}
