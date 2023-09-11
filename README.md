# GoPractice

## Key takeaway
1. There is a built-in testing package support
    - all test file name with `xxx_test.go`
    - `testing.T` for test state
    - `testing.B` for benchmark test
    - `t.run` to write multiple test cases within a general testing function
    - `go test -cover` run a coverage test

2. `const` for const stuff
3. `var <name> <type>` for variable, or `<name> := <value>` for shorthanded init
4. Exported names start with Capital letter; Internal start with small letter
5. variadic functions: take a variable number of arguements
    - `func test (numbersToSum ...[]int) int`
    - It means I can pass 1, 2, 3+ []int into the test function
6. Function is diff from method
    - method is a fucntion with a receiver
    - `func (receiver receiverBase / struct) <name>() <return type>`
    - E.g. `func (c Circle) Area() float64`