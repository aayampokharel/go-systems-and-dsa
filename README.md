# DSA in Go

My practice space for DSA (NeetCode, LeetCode, etc.) and Go interview snippets.
Goal: solve it myself first, test it, then learn the better way and write down why.

## Folder Structure

```
dsa/
├── arrays/
│   └── concatenation_of_array/
│       ├── question.md          # question copied as given
│       ├── main.go              # my own solution
│       ├── main_test.go         # my tests
│       ├── better_solution.md   # my approach, complexity, better solution, Go concepts learned
│       └── discussion.md        # things I didn't understand (optional)
├── strings/
├── ...
└── snippets/                    # small Go interview programs (HTTP, URL remover, etc.)
```

Each topic gets its own folder. Each problem gets its own subfolder.

## Steps for Every Problem

1. Copy the question into `question.md`.
2. Write my own solution in `main.go` and run it.
3. Write tests in `main_test.go`.
4. In `better-solution.md`, write:
   - my approach
   - my time and space complexity, and why
   - the better solution and why it is better
   - Go concepts I learned
5. If something confuses me, write it in `discussion.md`.

## Run Tests

```bash
go test ./...        # everything
go test -v           # inside one problem folder
go test -race ./...  # for concurrency snippets
```

## Snippets

Small Go programs I practice for interviews, each in its own folder under `snippets/` with tests.

- [ ] URL remover
- [ ] Worker pool
- [ ] TTL cache
- [ ] Rate limiter
- [ ] Simple REST API

## Progress

| Problem | Topic | Status |
|:--|:--|:--|
| Concatenation of Array | Arrays | Done |

Status: Done / Needed hint / Revisit