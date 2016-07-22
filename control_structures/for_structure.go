package control_structures

func collatzChainLength(n int) int {

	if n%2 != 0 {
		n = 3*n + 1 
	}

	i := 1;
	for n>1 {
		

		if n%2 != 0 {
		n = 3*n + 1 
		}else {
			n = n/2;
		}

		i++;

	}
	return i
}
