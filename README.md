# monkey - by Thorsten Ball

*Reference: [interpreterbook.com](https://interpreterbook.com) · [The Lost Chapter](https://interpreterbook.com/lost/)*

## C-like Syntax

```monkey
let x = 10;
let y = 20;

if (x < y) {
    return true;
} else {
    return false;
}
```

---

## Variable Bindings

```monkey
let name = "Monkey";
let version = 1;
let active = true;

name;    // "Monkey"
version; // 1
active;  // true
```

---

## Integers and Booleans

```monkey
let age = 42;
let isReady = true;
let isDone = false;

age;     // 42
isReady; // true
isDone;  // false
```

---

## Arithmetic Expressions

```monkey
let a = 10;
let b = 3;

a + b;  // 13
a - b;  // 7
a * b;  // 30
a / b;  // 3
a == b; // false
a != b; // true
a > b;  // true
a < b;  // false
```

---

## Built-in Functions

| Function | Description |
|---|---|
| `len(x)` | Returns the length of a string or array |
| `first(arr)` | Returns the first element of an array |
| `last(arr)` | Returns the last element of an array |
| `rest(arr)` | Returns a new array without the first element |
| `push(arr, x)` | Returns a new array with `x` appended |
| `puts(x)` | Prints a value to stdout |

```monkey
let arr = [10, 20, 30];

len(arr);        // 3
first(arr);      // 10
last(arr);       // 30
rest(arr);       // [20, 30]
push(arr, 40);   // [10, 20, 30, 40]
puts("Hello!");  // Hello!
```

---

## First-class and Higher-order Functions

```monkey
// Functions assigned to variables
let double = fn(x) { x * 2 };
let add = fn(a, b) { a + b };

double(5);    // 10
add(3, 4);    // 7

// Passing a function as an argument (higher-order)
let apply = fn(f, value) { f(value) };
apply(double, 6); // 12

// Returning a function from a function
let makeAdder = fn(n) {
    fn(x) { x + n }
};
let addFive = makeAdder(5);
addFive(10); // 15
```

---

## Closures

```monkey
let newCounter = fn() {
    let count = 0;
    fn() {
        count + 1
    }
};

let counter = newCounter();
counter(); // 1
counter(); // 1 (each call is independent — pure closure)

// A more expressive closure example
let makeMultiplier = fn(factor) {
    fn(x) { x * factor }
};

let triple = makeMultiplier(3);
triple(7);  // 21
triple(10); // 30
```

---

## String Data Structure

```monkey
let greeting = "Hello";
let name = "World";

let message = greeting + ", " + name + "!";
message;         // "Hello, World!"
len(message);    // 13

puts(message);   // Hello, World!
```
---

## Array Data Structure

```monkey
let numbers = [1, 2, 3, 4, 5];
let mixed  = [1, "two", true, [3, 4]];

first(numbers);      // 1
last(numbers);       // 5
rest(numbers);       // [2, 3, 4, 5]
push(numbers, 6);    // [1, 2, 3, 4, 5, 6]
len(numbers);        // 5

// Accessing by index
numbers[0];  // 1
numbers[2];  // 3
mixed[1];    // "two"
```

---

## Hash Data Structure

```monkey
let person = {
    "name": "Alice",
    "age": 30,
    "active": true
};

person["name"];   // "Alice"
person["age"];    // 30
person["active"]; // true

// Integer and boolean keys
let scores = { 1: "one", 2: "two", true: "yes" };
scores[1];    // "one"
scores[true]; // "yes"
```

---

### `reduce` — Fold an Array into a Single Value

```monkey
let reduce = fn(arr, initial, f) {
    let iter = fn(arr, result) {
        if (len(arr) == 0) {
            result
        } else {
            iter(rest(arr), f(result, first(arr)));
        }
    };
    iter(arr, initial);
};

// Build sum using reduce
let sum = fn(arr) {
    reduce(arr, 0, fn(initial, el) { initial + el });
};

sum([1, 2, 3, 4, 5]);
// 15

// Build product using reduce
let product = fn(arr) {
    reduce(arr, 1, fn(acc, el) { acc * el });
};

product([1, 2, 3, 4, 5]);
// 120
```

---

### Composing `map` and `reduce`


```monkey
let numbers = [1, 2, 3, 4, 5];

// Sum of squares
let square = fn(x) { x * x };

let sumOfSquares = fn(arr) {
    reduce(map(arr, square), 0, fn(acc, el) { acc + el })
};

sumOfSquares(numbers);
// 55  (1 + 4 + 9 + 16 + 25)
```

---

## Quick Reference

| Feature | Syntax Example |
|---|---|
| Variable binding | `let x = 5;` |
| Integer | `42` |
| Boolean | `true`, `false` |
| String | `"hello"` |
| Array | `[1, 2, 3]` |
| Hash | `{"key": "value"}` |
| Function | `fn(x) { x + 1 }` |
| Call | `add(1, 2)` |
| If/else | `if (x > 0) { x } else { -x }` |
| Return | `return x;` |
| Arithmetic | `+ - * /` |
| Comparison | `== != < >` |

---
