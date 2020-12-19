## Welcome to Page
### A Functional Programming Language Based on Lisp.
This is the interpreter for Page, a functional programming langauge based on Lisp. All parts were written from scratch using Go. The interpreter uses a recursive descent approach for parsing. 

### Program Examples
##### 1. Fibonacci

![Lisp Demo](./demos/fibonacci.gif)

##### 2. Factorial

![Lisp Demo](./demos/factorial.gif)

### Documentation

##### 1. Lists

```
> (list 1 2 3) // define list.
(list 1 2 3)

> (car (list 1 2 3)) // get first element
1

> (cdr (list 4 5 6)) // get rest
(list 5 6)

> (cons 1 (list 4 5 6)) // construct new list with elements combined
(list 1 4 5 6)

> (length (list 1 2 3))
3
```

##### 2. Function Definitions

```
> (define (add a b) (+ a b))

> (add 2 5)
7

> (define (isEven n) (= (% n 2) 0))

> (isEven 5)
False

```

##### 3. Strings

```
> (length "hello")
5

> (equals "hello" "world")
False

> "Hello World"
Hello World
```

##### 3. Numbers

```
> 10
5

> (+ 5 5 )
10

> (% 4 2)
0

> (/ 4 2)
2

> (* 2 2)
4
```






