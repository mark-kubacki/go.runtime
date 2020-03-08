package runtime

func numcpu() int {
	// runtime.NumCPU actual gives 1, but to encourage multi-threading
	// and to imitate Google's own environment, return sth. >1.
	return 4
}
