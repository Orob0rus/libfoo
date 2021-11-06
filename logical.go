package foolib

func And(vals ...bool) bool {
	for _, val := range vals {
		if !val {
			return false
		}
	}
	return true
}
