package data_structures

import "fmt"
type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {

	//a := vals[:]
	for index,element := range vals {
		vals[index] = element *element
	}

	return vals
}

func filterInts(op filterOperation, vals []int32) []int32 {
	
	res := []int32{0}
	var i int32

	a := vals[:]
	for _,element := range a {
		
		if(element>2) {
			i++
		}
	}

	b := vals[2:3]
	for _,element := range b {
		if(element>2) {
			i++
		}
	}

	c := vals[1:3]
	for _,element := range c {
		
		if(element>2) {
			i++
		}
	}
	res[0] = i
	return res
}

func concatenate(dest []string, newValues ...string) []string {
	newSplit := dest[:]

	for _,element := range newValues {
		newSplit = append(newSplit,element)
	}

	return newSplit
}

func equals(list1 []string, list2 []string) bool {

	if len(list1) != len(list2) {
        return false
    }

    for i := range list1 {
        if list1[i] != list2[i] {
            return false
        }
    }

    return true
}

func partialReverse(src []int, from, to int) []int {

	reverse := []int{}
	partialReverseSlice := src[from:to+1]
	
	fmt.Println("%d",len(partialReverseSlice))

	
	for i:=len(partialReverseSlice);i>0;i-- {
		fmt.Println("%d",partialReverseSlice[i-1])
		reverse = append(reverse,partialReverseSlice[i-1])
		
	}
	return reverse
}
