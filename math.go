package foolib

func add2Int(num1, num2 int) int {
	return num1+num2
}

func sub2Int(num1, num2 int) int {
	return num1-num2
}

func Add2Int(num1, num2 int) int {
	return TwoIntOperation(add2Int, num1, num2)
}

func Add2IntCurried(num1 int) func(int) int {
	return TwoIntOperationCurried(add2Int, num1)
}

func Sub2Int(num1, num2 int) int {
	return TwoIntOperation(sub2Int, num1, num2)
}

func Sub2IntCurried(num1 int) func(int) int {
	return TwoIntOperationCurried(sub2Int, num1)
}

func TwoIntOperation(f func(int,int) int, num1, num2 int) int {
	return f(num1, num2)
}

func TwoIntOperationCurried(f func(int,int) int, num1 int) func(int) int {
	return func (num2 int) int {
		return f(num1, num2)
	}	
}