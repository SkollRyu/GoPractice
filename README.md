# GoPractice
I followed [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests) to learn GO in a TDD approach

## Nice hacks

In our function signature we have made a named return value (prefix string).
This will create a variable called prefix in your function.
This will also display in the Go Doc, so it makes things easier
```
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
```

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
    - `func (t Type) methodName(parameter list)`
    - E.g. `func (c Circle) Area() float64
7. Interface is implicit, no keyword `implement`
    - `type <name> interface`
    - once other struct implements all the methods from interface, it is already using that interface
    - Declaring interfaces so you can define functions that can be used by different types (Parametric polymorphism)
8. Table-driven test
    - Each table entry is a complete test case with inputs and expected results
    - create anonymous struct: `[]struct{ interface InterfaceBase ...}{{test case input, output}, {test case 2 input, output},}`
    - loop over test case: `for _, testCase := range testCases {...}`
9. New type from existing one
    - except `struct`, we can create a less overkilled type
    - syntax: `type <MyName> <OriginalType>`
    - E.g: `type Bitcoin int`
    - we can also declare methods for the type
10. Pointer
    - Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
    - pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools)
    - `func Add(x int)` which take an integer as the parameter.
    - `func AddPtr(x *int)` which takes a pointer integer as the parameter.
    - `AddPtr(&abc)` which pass the memory address
11. nil
    - Pointers can be nil
    - When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
    - Useful for when you want to describe a value that could be missing
12. Errors
    - Errors are the way to signify failure when calling a function/method.
    - 