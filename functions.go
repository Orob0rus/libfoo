package foolib

func Flip(function func(args ...interface{}) (interface{}, error)) func(args ...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		size := len(args)
		if size < 2 {
			return function(args)
		}
		var processedArgs []interface{}
		if size > 2 {
			copy(processedArgs, args[2:])
		}
		processedArgs = append([]interface{}{args[1], args[0]}, processedArgs...)
		if val, err := function(processedArgs); err != nil {
			return nil, err
		} else {
			return val, nil
		}
	}
}
