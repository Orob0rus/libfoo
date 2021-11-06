package foolib

import (
	"fmt"
)

var ErrorSliceIndexOutOfBounds error = fmt.Errorf("slice index out of bounds")

func MapperSlice(f func(interface{}) (interface{}, error), slice []interface{}) ([]interface{}, error) {
	output := make([]interface{}, len(slice))
	for idx, val := range slice {
		if value, err := f(val); err != nil {
			return nil, err
		} else {
			output[idx] = value
		}
	}
	return output, nil
}

func AdjustSlice(index int, f func(interface{}) (interface{}, error), slice []interface{}) ([]interface{}, error) {
	if index >= len(slice) || index < 0 {
		return nil, ErrorSliceIndexOutOfBounds
	}
	output := make([]interface{}, len(slice))
	for idx, val := range slice {
		if idx == index {
			if value, err := f(val); err != nil {
				return nil, err
			} else {
				output[idx] = value
			}
			continue
		}
		output[idx] = val

	}
	return output, nil
}

func AllSlice(condition func(interface{}) (bool, error)) func([]interface{}) (bool, error) {
	return AllPassSlice([]func(interface{}) (bool, error){condition})
}

func AllPassSlice(conditions []func(interface{}) (bool, error)) func([]interface{}) (bool, error) {
	return func(slice []interface{}) (bool, error) {
		for _, condition := range conditions {
			for _, val := range slice {
				if pass, err := condition(val); err != nil || !pass {
					return pass, err
				}
			}
		}
		return true, nil
	}
}
