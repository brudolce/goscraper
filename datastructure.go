package main

// golang has no "Set" datatype, for unique stuff. Probably cause it has no generics - implementing a "string" Set
func uniqueStringArray(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
