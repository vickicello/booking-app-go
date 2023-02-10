# Go Notes

## variables 

* var can be reassigned, const cannot
* you can declare variable before providing a value like `var foo int` then provide it later using the pointer `&foo`
* or you can provide the value when you declar it like `var foo int = 50`
* or use the shorthand foo := 50 (Go can infer that 50 is int)
* shorthand := cannot be used for package level variables though
* style: const should be placed above var for package level variables

## arrays and slices

* arrays have a fixed length and can be declared like `var foo [10]string` (an array of strings with 10 elements).  Use len() to get the length and `foo[index] = 'bar'` to add items
* you can't mix datatypes in an array
* slices are an abstraction of an array that allow for flexible sizing 
* can be declared like `var foo []string` and you can use append() to add to the slice


## loops

* Go only uses for loops, no while or until (but you can basically do an until/while by adding a condition after `for` like `for remainingTickets > 0`...or you could make an infinite loop with `for true` which is the same as just using `for` without a condition)
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

## organization

* we can modularize our code by adding sub packages in their own folders, and then import them into the main package
* to export functions and variables, capitalize the first letter of the function or variable name
* run all of the go subpackages and files with `go run .` to run the entire folder

## scope

* 3 levels of scope: local (inside the function/block only), package (inside the package only), global (can be accessed across packages)
* global vars have capitalized variable name and are placed outside of functions


## maps

* to declare a map, you define the datatype of the key and the datatype of the value: `map[string]uint`; create it by using the built in make() function: `make(map[string]uint)`
* add to the map like `mapName["key"] = value`
* just like arrays, you can't mix datatypes in maps - they all have to match the declaration you create
* you can make a slice of maps (variable length maps starting at size n) `var myMap = make([]map[string]int, n)`
