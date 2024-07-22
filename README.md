# Kamelang
## Kamelang is a simple programming language that is interpreted by Go. The goal of this project is to learn how to create a programming language and to understand how a programming language works.

## Features:

- [x] Variable Declaration and Assignment
- [x] Arithmetic Operations
- [x] Logical Operations
- [x] If Statement
- [x] Function Declaration and Call
- [x] Higher Order Function
- [x] Return Statement
- [x] Implicit Return
- [x] Recursion

## Data Types:
- [x] Integer
- [x] Boolean
- [x] String
- [x] Array
- [x] Hash

## Built-in Functions:
- [x] len
- [x] first
- [x] last
- [x] rest
- [x] push
- [x] puts

### len: 
Returns the length of an array.\
Syntax: ``` len(array) ```

### first:
Returns the first element of an array.\
Syntax: ``` first(array) ```

### last:
Returns the last element of an array.\
Syntax: ``` last(array) ```

### rest:
Returns all the elements of an array except the first element.\
Syntax: ``` rest(array) ```

### push:
Adds an element to the end of an array.\
Syntax: ``` push(array, element) ```

### puts:
Prints the value of an expression and a null at the end\
Syntax: ``` puts(expression) ```

### Kamel is tokenized and parsed in a REPL. The REPL is a simple command line interface that reads input, evaluates it, prints the result, and loops until the user exits.

## The Interpreter is divided into 5 main parts:
- Lexer
- Parser
- AST(Abstract Syntax Tree)
- Internal Object system
- Evaluator

## REPL
```shell
  Hello Kamel: Username
  Welcome to Kamelang!
  >> let a = 10;
  >> let b = 20;
  >> a + b;
  30
  >> let add = fn(a, b) { a + b; };
  >> add(1, 2);
  3
  ...
```

## Syntax:

### Variable Declaration and Assignment
```shell
    >> let a = 10;
    >> let a = true; 
    >> let a = "Hello Kamel!";
    >> let a = [1, 2, 3, 4];
    >> let a = {"name": "Kamel", "age": 20};
``` 

### Arithmetic Operations
```shell
    >> let a = 10 + 20;
    >> let b = 300 - 20 / 2 * 3; 
```

### Logical Operations
```shell
    >> true == false;
    >> true != false;
    >> 1 < 2;
    >> 1 > 2;
    >> (1 < 2) == true; 
```

### Indexing Arrays and Hashes
```shell
    >> let a = [1, 2, 3, 4];
    >> a[0];
    1
    >> let b = {"name": "Kamel", "age": 20};
    >> b["name"];
    Kamel
```

### If Statement
```shell
    >> if (1 < 2) { return true; } else { return false; }
    >> if (1 < 2) { 10 } else { 20 }
    >> if (10 > 1) { if (10 > 1) { return 10; } return 1; }
```

### Function Declaration and Call
```shell
    >> let add = fn(a, b) { return a + b; };
    >> add(1, 2); 
```

### Higher Order Function
```shell
    >> let add = fn(a, b) { return a + b; };
    >> let apply = fn(a, b, fn) { return fn(a, b); };
    >> apply(1, 2, add); 
```

### Return Statement
```shell
    >> return 10;
    >> return true; 
```

### Implicit Return
```shell
    >> let add = fn(a, b) { a + b; };
    >> add(1, 2);
```    

### Recursion
```shell
    >> let fibonacci = fn(x) { 
        if (x == 0) { return 0; } 
        else { 
            if (x == 1) { return 1; } 
            else { return fibonacci(x - 1) + fibonacci(x - 2); }} 
        }; 
    >> fibonacci(10); 
```

### String Concatenation
```shell
    >> let a = "Hello" + " " + "Kamel!";
    >> a;
    Hello Kamel!
```

### Built-in Functions
```shell
    >> let a = [1, 2, 3, 4];
    >> len(a);
    4
    >> first(a);
    1
    >> last(a);
    4
    >> rest(a);
    [2, 3, 4]
    >> push(a, 5);
    [1, 2, 3, 4, 5]
    >> puts("Hello from Kamelang!", "Hello World!");
    Hello from Kamelang!
    Hello World!
    null
```

## How to run the REPL:
```shell
    docker pull symelbak/kamelang
    docker run -it symelbak/kamelang
```

## Future Plans:
- [ ] Convert the interpreter to a compiler\