# Go Notes

## variables 

* var can be reassigned, const cannot
* you can declare variable before providing a value like `var foo int` then provide it later using the pointer `&foo`
* or you can provide the value when you declar it like `var foo int = 50`
* or use the shorthand foo := 50 (Go can infer that 50 is int)

## arrays and slices

* arrays have a fixed length and can be declared like `var foo [10]string` (an array of strings with 10 elements).  Use len() to get the length and `foo[index] = 'bar'` to add items
* you can't mix datatypes in an array
* slices are an abstraction of an array that allow for flexible sizing 
* can be declared like `var foo []string` and you can use append() to add to the slice


## loops

* Go only uses for loops, no while or until (but you can basically do an until/while by adding a condition after `for` like `for remainingTickets > 0`...or you could make an infinite loop with `for true` which is hte same as just using `for` without a condition)
* use `break` to exit the loop, or `continue` skip the rest of this iteration and move to the next

## conditionals

* uses traditional `if/else/else if` syntax
* can also use switch statements

## functions and return values

* functions are declared with `func <funcName>()` 
* parameters for the function need to specify the type
* if a return value is used, you also need to specify the type of the return value like `func foo(mySlice []string) []string {` (indicates this function will return a slice of strings)
* in Go, you can return `n` number of return values in a function, unlike most languages
* just return a, b, c and add their return values in the func definition line (bool, string, int) etc.  And you can also assign multiple vars to the return vals like `a, b, c := func(x, y, z)`
