package lib

//  CheckErr 处理错误
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Fibonacci(index int) int64 {
	arr := []int64{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	if index+1 > len(arr) {
		return arr[len(arr)-1]
	}
	return arr[index]
}
