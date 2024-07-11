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

## How to run the REPL:
```shell
    docker pull symelbak/kamelang
    docker run -it symelbak/kamelang
```

## Future Features:
- [ ] Arrays
- [ ] Strings
- [ ] Convert the interpreter to a compiler\
\
And much more...