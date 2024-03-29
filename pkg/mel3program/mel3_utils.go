package mel3program

// Get the program id from its name whitin the implementation
func ids_from_name(implementation map[uint16]*Mel3Implementation, input_name string) ([]uint16, []uint16, bool) {
	if implementation == nil {
		return []uint16{}, []uint16{}, false
	}

	libraryids := make([]uint16, 0)
	programids := make([]uint16, 0)

	exists := false

	for libid, impl := range implementation {
		for key, value := range impl.ProgramNames {
			if value == input_name {
				libraryids = append(libraryids, libid)
				programids = append(programids, key)
				exists = true
			}
		}
	}

	if exists {
		return libraryids, programids, true
	}

	return []uint16{}, []uint16{}, false
}

// Get the program id from its name whitin the implementation
func id_from_name(implementation *Mel3Implementation, input_name string) (uint16, bool) {
	if implementation == nil {
		return 0, false
	}

	for key, value := range implementation.ProgramNames {
		if value == input_name {
			return key, true
		}
	}

	return 0, false
}

func count_program(program *Mel3Program) int {
	result := 1

	for i := range program.NextPrograms {
		result = result + count_program(program.NextPrograms[i])
	}

	return result
}

func node_find(inprogr *Mel3Program, starting_node int, searched_node int) *Mel3Program {
	var result *Mel3Program
	if starting_node == searched_node {
		return inprogr
	} else {
		current_node := starting_node + 1
		old_current_node := starting_node + 1

		for i := range inprogr.NextPrograms {
			current_node = current_node + count_program(inprogr.NextPrograms[i])

			if searched_node < current_node {
				result = node_find(inprogr.NextPrograms[i], old_current_node, searched_node)
				break
			}
			old_current_node = current_node
		}
	}
	return result
}

func copy_program(inprogr *Mel3Program) *Mel3Program {

	outprogr := new(Mel3Program)

	if inprogr != nil {

		outprogr.ProgramID = inprogr.ProgramID
		outprogr.ProgramValue = inprogr.ProgramValue
		outprogr.NextPrograms = make([]*Mel3Program, len(inprogr.NextPrograms))

		for i, prog := range inprogr.NextPrograms {
			outprogr.NextPrograms[i] = copy_program(prog)
		}
	}
	return outprogr
}
