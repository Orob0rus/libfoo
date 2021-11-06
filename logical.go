package foolib

func And(vals ...bool) bool {
	for _, val := range vals {
		if !val {
			return false
		}
	}
	return true
}

func Any(function func(interface{}) (bool, error)) func([]interface{}) (bool, error) {
	return func(vals []interface{}) (bool, error) {
		for _, val := range vals {
			if out, err := function(val); err != nil {
				return false, err
			} else if out {
				return true, nil
			}
		}
		return false, nil
	}
}
