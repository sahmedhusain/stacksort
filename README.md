# Push Swap

Push-Swap is a Go project that involves sorting a list of integers using two stacks (A and B) and a set of predefined instructions.

The project consists of two programs: `push_swap` and `checker`.

## Project Structure

```sh
PUSH-SWAP
├── CheckerC
│   └── main.go
├── Functions
│   ├── CheckerStacks
│   │   └── stacks.go
│   ├── PushInstructions
│   │   └── instructions.go
│   ├── PushSort
│   │   └── sort.go
│   ├── PushStacks
│   │   └── stacks.go
├── PushSwapC
│   └── main.go
├── .gitignore
├── go.mod
├── LICENSE.md
├── README.md
└── test.sh
```

## Instructions

- **pa**: Push the top element of stack B to stack A.
- **pb**: Push the top element of stack A to stack B.
- **sa**: Swap the first two elements of stack A.
- **sb**: Swap the first two elements of stack B.
- **ss**: Execute `sa` and `sb`.
- **ra**: Rotate stack A (shift up all elements by 1, the first element becomes the last).
- **rb**: Rotate stack B.
- **rr**: Execute `ra` and `rb`.
- **rra**: Reverse rotate A (shift down all elements by 1, the last element becomes the first).
- **rrb**: Reverse rotate B.
- **rrr**: Execute `rra` and `rrb`.

Return `n` size of instructions for sorting `x` number of values:
- If `x = 3`, then `n <= 3`.
- If `x = 5`, then `n <= 12`.
- If `x = 100`, then `n <= 1500`.
- If `x = 500`, then `n <= 11500`.

## Programs

### Push Swap

This program calculates and displays the smallest set of instructions to sort stack `a` in ascending order.

**Usage:**

```sh
$ ./push_swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
```

<br>

### Checker

This program reads instructions from standard input and executes them on the given stack A. It then checks if stack A is sorted and stack B is empty.

**Usage:**

```sh
$ ./checker "3 2 1 0"
sa
rra
pb

KO

$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
```

1. Clone the repository:
```sh
git clone https://github.com/sahmedhusain/push-swap.git
```

2. Navigate to the project directory:
```sh
cd push-swap
```

3. Build the project:
```sh
go build -o push-swap PushSwapC/main.go
go build -o checker CheckerC/main.go
```

### Example

```sh
$ ARG="2 1 3 6 5 8"; ./push-swap "$ARG" | wc -l
8

$ ARG="2 1 3 6 5 8"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```
## Author

	•	Sayed Ahmed Husain

## This project has helped me learn about:

	•	The use of basic algorithms
	•	The use of sorting algorithms
	•	The use of stacks
