package foolib

func Always(val interface{}) func() interface{} {
	return func() interface{} {
		return val
	}
}

