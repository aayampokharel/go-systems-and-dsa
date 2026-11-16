# DSA & Go Interview Preparation

This repository is my personal learning space for revisiting **Data Structures and Algorithms (DSA)** and preparing for Go developer interviews.

I am starting from the basics and gradually working towards solving more complex problems.

I will be practicing problems from [NeetCode](https://neetcode.io/), [LeetCode](https://leetcode.com/), and other resources.

The goal is not just to solve problems but to understand them, improve my problem-solving skills, write clean Go code, and understand the performance of different solutions.

## Goals

- Relearn DSA concepts from the basics.
- Practice problems from NeetCode, LeetCode, and other resources.
- Solve problems independently before looking at other solutions.
- Write clean and readable Go code.
- Understand time and space complexity.
- Compare my solutions with more optimal solutions.
- Write Go unit tests for every problem.
- Learn new concepts, techniques, and Go-specific features along the way.
- Practice small Go programming tasks commonly asked in technical interviews.

## Repository Structure

```text
.
├── dsa/                    # DSA problems and solutions
│   ├── arrays-and-hashing/
│   ├── two-pointers/
│   ├── sliding-window/
│   ├── stack/
│   ├── binary-search/
│   ├── linked-list/
│   ├── trees/
│   ├── graphs/
│   ├── dynamic-programming/
│   └── backtracking/
│
├── go-interview/            # Small Go interview coding exercises
│   ├── http/
│   ├── strings/
│   ├── concurrency/
│   ├── channels/
│   ├── goroutines/
│   └── miscellaneous/
│
├── concepts/                # Important DSA and Go concepts
│
└── new-things-learned/      # New terms and concepts I discover
```

## 1. DSA Problem-Solving Workflow

For every problem, I will follow this process:

1. **Understand the problem:** Read the question carefully and identify the expected input and output.
2. **Think independently:** Try to solve the problem without immediately looking at the solution.
3. **Write my solution:** Implement my approach in Go.
4. **Test my solution:** Write Go unit tests and test different cases, including edge cases.
5. **Analyze complexity:** Calculate the time and space complexity of my solution and explain why.
6. **Find a better approach:** Look at NeetCode, LeetCode discussions, or other explanations to understand alternative solutions.
7. **Compare solutions:** Implement a more optimal solution if one exists and compare both approaches.
8. **Document learning:** Write down new algorithms, patterns, Go features, and mistakes I discovered.

## 2. Problem Documentation

Each problem will have its own directory containing:

- `question.md` – Problem statement, examples, constraints, and my understanding of the problem.
- `solution.go` – My implementation and, where applicable, the optimal implementation.
- `solution_test.go` – Unit tests written using Go's testing package.

For each problem, I will document:

### Problem Information

- Problem name and source link.
- DSA topic or pattern.
- Difficulty level.

### My Approach

- How I understood the problem.
- My initial thought process.
- Why I chose this approach.
- Any difficulties or mistakes I encountered.

### My Solution

- Go implementation.
- Time complexity and why.
- Space complexity and why.

### Optimal Solution

- Better or alternative approach, if available.
- Go implementation.
- Time complexity and why.
- Space complexity and why.

### Comparison

Compare my solution with the optimal solution based on:

- Time complexity.
- Space complexity.
- Readability.
- Simplicity.
- Trade-offs.

The goal is to understand why one solution is better than another, not just memorize the optimal code.

## 3. Go Interview Practice

Apart from DSA, I will maintain a separate collection of small Go programming exercises that are commonly asked in technical interviews.

Examples:

- Remove URL schemes (`http://`, `https://`) from a string.
- Reverse a string.
- Check whether a string is a palindrome.
- Find duplicate elements in a slice.
- Implement a simple HTTP URL parser.
- Work with maps and slices.
- Implement custom sorting.
- Practice goroutines and channels.
- Solve concurrency-related problems.

Each exercise will contain working Go code and tests wherever applicable.

## 4. Concepts and New Things Learned

While solving problems, I will encounter new terms, algorithms, Go features, and programming concepts.

I will document them instead of simply moving on.

Examples:

- Two Pointers
- Sliding Window
- Hash Maps
- Binary Search
- Recursion
- Backtracking
- Big O Notation
- Amortized Time Complexity
- Go Slice Internals
- Variadic Functions
- Mutex and RWMutex
- Goroutine Synchronization

These notes will help me revise concepts without having to revisit every problem.

## 5. Testing

I will use Go's built-in testing package.

```bash
go test ./...
```

I will try to cover:

- Normal test cases.
- Empty inputs.
- Single-element inputs.
- Duplicate elements.
- Boundary conditions.
- Large inputs where relevant.

I will write tests independently rather than relying only on the examples provided in the problem.

## 6. Learning Resources

- [NeetCode](https://neetcode.io/)
- [LeetCode](https://leetcode.com/)
- [Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Big-O Cheat Sheet](https://www.bigocheatsheet.com/)

---

## My Approach

I am not trying to solve hundreds of problems as quickly as possible.

I want to understand the fundamentals, recognize common patterns, write correct and efficient Go code, and become confident explaining my solutions during technical interviews.

**Understand first. Solve independently. Test properly. Analyze complexity. Improve continuously.**