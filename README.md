# Golang Help

- <a href="https://go.dev/ref/spec" target="_blank">Golang Language Specification</a>
- <a href="https://gobyexample.com/" target="_blank" >Go by Example</a>
- <a href="https://go.dev/play/" target="_blank">Go Playground</a>
- <a href="https://pkg.go.dev/std" target="_blank">Go Packages</a>
- <a href="https://godoc.org">GoDoc</a>

# Content

<a href="#function">Functions</a> | <a href="#structs">Structs</a>

# Variables, Values and Type

# Different Types

# Conditionals

# Composite Data Types

<h1 id="structs">Structs</h1>

### What is a struct?

- A struct is a data structure that **compose together values of different types**.

```
type person struct{}
```

- Empty struct

```
type person struct {
	first string
	last  string
	age int
}
```

- Struct fields must be unique. Cannot have another `first string`.
- Must have an identifier and type.

### Embedded Struct

```
type secretAgent struct {
	person	// another struct
	ltk bool
}
```

- Anonymous field: a field declared with a type but no explicit field name.
- A field from embedded struct is promoted to the outer level.

### Anonymous Struct

```
	p3 := struct {
		first string
		last  string
		age int
	}{
		first: "Tuck Yern",
		last:  "Chan",
		age: 25,
	}
```

<h1 id="function">Functions</h1>

### Syntax

- `func (r receiver) identifier(parameters) (return(s)) {...}`
- **Difference between parameters and arguments**: we define our func with parameters (if any) and we call our func and pass in arguments.

- Example Syntax:

```
func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := true
	return a, b
}
```

### Variadic Parameter

```
func foo(x ...int) int {
```

- Takes in an unlimited number of int as the parameter.

```
z := foo()
```

- **Variadic Parameter**: Passing in 0 or more arguments. Must be the final parameter (...T)
- `func foo(s string, x ...int)` allowed. `z := foo("Leon")`
- `func foo(x ...int, s string)` not allowed.

### Unfurling a Slice

```
	yi := []int{2,3,4,5,6,7,8,9}
	y := foo(yi...)
	fmt.Println(y)
```

- Using the `...` notation, can unpack a slice into the integers.

### defer

- `defer` is a Keyword.
- `defer`: defer the execution of a function until whereever it is being called comes to an end.

```
func main() {
	defer foo()
	bar()
} // foo() is ran right here before the function ends

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
```

- Deferring the function `foo()` in `main()`

### Receiver

```
func (s secretAgent) speak()
```

- Any value of type secretAgent has access to this method.
- Attaches the function `speak()` to the type `secretAgent`

### Interface

```
type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}
```

- Anything that has the method `speak()` is also of type `human`.

```
func bar(h human) {
	switch h.(type) {  // this is an assert; asserting, "h is of this type"
	case person:
		fmt.Println("I was passed into bar!!!", h.(person).first)
	case secretAgent :
		fmt.Println("I was passed into bar!!!", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}
```

- **Assertion** using `switch` statement. Asserting the type.

### Anonymous Function

- Anonymous Function with no parameters

```
func main() {

	// Anonymous Function
	func() {
		fmt.Println("Anonymous function ran")
	}()

}
```

- Anonymous Function with parameters

```
func (x int) {
	fmt.Println("The meaning of life:", x)
} (42)
```

### Func Expression

- Assigning function to variables

```
func main() {
	f := func() {
		fmt.Println("my first func expression")
	}

	f()

	y := func(x int) {
		fmt.Println("The year I was born in is", x)
	}

	y(2003)
}
```

### Returning a Function

- Return a string

```
func main() {
	s1 := foo()
	fmt.Println(s1)
}

func foo() string {
	s := "Hello World!"
	return s
}
```

- Returns a function that returns an int

```
func bar() func() int {
	return func() int {
		return 451
	}
}

func main() {
	x := bar()
	fmt.Println(x())  // calling func() int
	fmt.Printf("%T\n", x())
}
```

- Can also run `bar()` entirely
  `fmt.Println(bar()())`

### Function Callback

- Passing a function as an argument to another function.

```
func even(f func(xi ...int) int, vi ...int) int {
	// Finding the even number
	var yi []int
	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

func main() {
	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)
}
```

### Closure

- Using a block or scope to enclose some variables.

### Recursive

- **Recursion**: When a function calls itself.

```
func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)
}

func factorial (n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}
```

# Additional Code

|         Code         |               Description                |
| :------------------: | :--------------------------------------: |
| `func foo(x ...int)` | Passing in an unlimited number of `int`. |
|        `s...`        |            Unfurling a Slice.            |

# Pointers

### What are pointers?

- Storing values in Golang creates a memory address.

|  Code  |                               Description                               |
| :----: | :---------------------------------------------------------------------: |
|  `&a`  |               Gives you the **address**. `"Pointer to a"`               |
|  `*a`  | Gives you the **value** stored at an address. "Deferencing an address". |
| `*int` |               Pointer to the type `int` stored in memory.               |

### When to use pointers?

- When you have a large chunk of data, and you don't want to pass the huge chunk around, just use the address where that value is stored.

### Method Sets

- Method sets determine what methods attach to a TYPE.

|     Method Sets      |                     Description                      |
| :------------------: | :--------------------------------------------------: |
| Non-Pointer Receiver | Works with values that are POINTERS or NON-POINTERS. |
|   Pointer Receiver   |      ONLY works with values that are POINTERS.       |

| Receivers |  Values   |
| :-------: | :-------: |
|  `(t T)`  | T and \*T |
| `(t *T)`  |    \*T    |

# Application

### JSON Documentation
