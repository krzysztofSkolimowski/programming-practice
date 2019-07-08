package diving_board

//recursive brute force
func getAllLengths(numberOfBoards, shorter, longer int) map[int]struct{} {
	lengthsSet := make(map[int]struct{})

	total := 0
	addBoards(total, numberOfBoards, shorter, longer, lengthsSet)

	return lengthsSet
}

func addBoards(total, numberOfBoards, shorter, longer int, lengthSet map[int]struct{}) {
	if numberOfBoards != 0 {
		_, ok := lengthSet[total+shorter]
		if !ok {
			lengthSet[total+shorter] = struct{}{}
		}

		addBoards(total+shorter, numberOfBoards-1, shorter, longer, lengthSet)

		_, ok = lengthSet[total+longer]
		if !ok {
			lengthSet[total+longer] = struct{}{}
		}

		addBoards(total+longer, numberOfBoards-1, shorter, longer, lengthSet)
	}
}

//sum approach
func allLengthsSums(numberOfBoards, shorter, longer int) map[int]struct{} {
	lengthsSet := make(map[int]struct{})
	for i := 0; i <= numberOfBoards; i++ {
		sum := i*shorter + (numberOfBoards-i)*longer
		_, ok := lengthsSet[sum]
		if !ok {
			lengthsSet[sum] = struct{}{}
		}
	}

	return lengthsSet
}
