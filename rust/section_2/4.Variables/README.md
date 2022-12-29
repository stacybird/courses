Strongly typed language:

can either say `let fox = 2;` or `let fox: i32 = 2;`

If annotation is left out, rust makes a best guess and sticks to that type.

Variables are immutable.  fixed on creation.
Why? immutable for safety concurrency and speed.  

to make a mutable variable: `let mut fox = 4;`

while variables are immutable, there are constants.
convention is caps, snake.  type is required:
`const COOL_VAR: i32 = 31415;`

so why const? 
* scope.  think outside the function level, like a global
* speed.  inline at compile time.
