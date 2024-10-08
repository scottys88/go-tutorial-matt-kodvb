# Control Structures

## If, then statements

If-then statements require braces, e.g.

```

if a == b {
    fmt.Println("a  equals b")
} else {
    fmt.Println("a does not equal b)
}

```

Short declaration in an `if` statement.

```

if err := doSomething(); err != nill {
    return err
}

// Is the equivalent of the below.
// Can declare and assign to err in the same line
// as the if statement

err := doSomething();

if err != nill {
    return err
}

```

## For loops

There are several different ways to perform a for loop.

1. Explicit control with an index variable

```
for i := 0; i < 10; i++ {
    fmt.Printfmt("The value oof i is %v\n", i);
}

```

2. Implicit control through the range operator for arrays and slices

The range operator will output two variables which can be used within the loop.
The first variable will be the index `int` the second will be the value `type`.

```
for i := range myArray {
    fmt.Println("the value of index " + i + " of myArray is " + myArray[i])
}

for i, v := range myArray {
    fmt.Println("the value of index " + i + " of myArray is " + v)
}
```

The range operator can also be used on a map. The difference is that the variables
made available to the loop are not index and value like an array. It is key and value.

```
myMap := make(map[string]int)
myMap["hi"]++;

for k, v := range map {
    fmt.Printfmt("the key of the map is ")
}
```

e.g. https://go.dev/play/p/5Vmw6ruerKQ

## Switch statements in Go

Switch statements basically syntatic sugar on top of a series of 'if, else' statements. 

Can be represented like so:

```
switch a := getVal(); a {
    case 0, 1, 2:
        fmt.Println("Do something")
    case 3, 4, 5, 6:
        // can be a nop (no op)
    default:
        fmt.Println("This is the default statment")
}

// The above is using the short declaration style but is equivalent to:

var b = getVal()
switch b {
    // Same switch statements as above
}
```

In Go a switch statement can also be represented like the below. See example: https://go.dev/play/p/_OOtV5Y3H0Q

```

var a = getVal()

switch {
    case a < 1:
        fmt.Println("a is less than one)

    case a == 2

    default:
        fmt.Println("This is the default)
}


```


