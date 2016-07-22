package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {

	people := []string{}
	for key,value := range data {
		for _,element := range value {
			if element == interest {
				people = append(people,key)
			}
		}
	}
	return people
}

func contains(src []string, elem string) bool {

	for _,element := range src {
		if element == elem {
			return true
		}
	}
	return false
}
