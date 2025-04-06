package hanoi

var m int

func MoveHanoi(n int, source, target, auxiliary string, mainProblem bool) {
	if mainProblem {
		m = n
		println("Problem size:", m)
	}

	if n == 1 {
		println("Move disk 1 from", source, "to", target)
		return
	}
	MoveHanoi(n-1, source, auxiliary, target, false)
	println("Move disk", n, "from", source, "to", target)
	MoveHanoi(n-1, auxiliary, target, source, false)
}
