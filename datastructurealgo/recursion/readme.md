# Recursion

## 1. Introduction
Recursion is a programming technique where a function calls itself to solve a problem in smaller parts. It involves breaking down a complex problem into simpler subproblems and then combining their solutions to get the final result.

## 2. Recursive Methods
Below are some recursive methods along with a brief description:

- **`geometricSum(k int) float64`**: Calculates the sum of the geometric series 1 + 1/2 + 1/4 + ... + 1/2^k.

- **`countZeros(n int) int`**: Counts the number of zeros in the decimal representation of a non-negative integer.

- **`multiplication(m int, n int) int`**: Performs multiplication of two integers without using the `*` operator.

- **`sumOfDigit(i int) int`**: Calculates the sum of digits in a given non-negative integer.

- **`numberOfDigit(i int) int`**: Counts the number of digits in a given non-negative integer.

- **`printNum(i int)`**: Prints numbers from 1 to `i` in increasing order using recursion.

- **`factorial(n int) int`**: Calculates the factorial of a non-negative integer `n`.

- **`fibonacci(n int) int`**: Calculates the `n`th Fibonacci number.

- **`power(x int, n int) int`**: Calculates `x` raised to the power of `n`.

## 3. Types of Recursion
Recursion can be categorized into two main types:

- **Direct Recursion**: When a function directly calls itself to solve a subproblem.

- **Indirect Recursion**: When a function calls another function(s), and eventually, one of the functions in the chain calls back the original function.

## 4. Advantages of Recursion
- Recursion can lead to elegant and concise code when dealing with complex problems.
- It can simplify the implementation of algorithms based on divide-and-conquer strategies.
- Some problems are naturally suited for recursive solutions, making the code more intuitive.

## 5. Disadvantages of Recursion
- Recursive solutions may consume more memory compared to iterative solutions due to function call overhead and maintaining the call stack.
- If not implemented carefully, recursion can lead to infinite loops and stack overflow errors.
- Recursive solutions may be less efficient than iterative solutions for some problems, especially when tail recursion is not used.

## 6. Tail Recursion
Tail recursion is a special form of recursion where the recursive call is the last operation performed in the function. Tail-recursive functions can be optimized by the compiler to use constant stack space, effectively converting the recursion into iteration.

## 7. Conclusion
Recursion is a powerful technique that can simplify problem-solving and lead to elegant code. However, it should be used judiciously, considering potential memory and performance implications. Tail recursion can be a useful optimization for recursive functions, eliminating the risk of stack overflow.
