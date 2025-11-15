package main

// If you had initialized go mod with go mod init main you will be presented with an error _testmain.go:13:2: cannot import "main".
//  This is because according to common practice, package main will only contain integration of other packages
//  and not unit-testable code and hence Go will not allow you to import a package with name main.
func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	return sum
}
