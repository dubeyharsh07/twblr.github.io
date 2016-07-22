package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {

	a := vals[:]
	for index,element := range a {
		vals[index] = element *element
	}

	return vals
}

func filterInts(op filterOperation, vals []int32) []int32 {
	a := vals[:]
	res := []int32{0}
	var i int32
	for _,element := range a {
		
		if(element>2) {
			i++
		}
	}
	res[0] = i
	return res
}

func concatenate(dest []string, newValues ...string) []string {
	return nil
}

func equals(list1 []string, list2 []string) bool {
	return false
}

func partialReverse(src []int, from, to int) []int {
	return nil
}
