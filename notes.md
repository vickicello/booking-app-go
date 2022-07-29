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
