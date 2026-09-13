# 🥞 StackSort

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE.md)

**StackSort** is a high-efficiency dual-stack sorting algorithm generator and verifier implemented in Go. It computes the absolute minimal sequence of push, swap, and rotate operations required to sort an unorganized list of integers across two auxiliary stack structures (`Stack A` and `Stack B`).

---

## ⚡ Key Highlights

- **Dual Binary Architecture**: Includes both `stacksort` (instruction generator) and `checker` (instruction validation engine).
- **Instruction Optimization Constraints**: Guaranteed instruction thresholds:
  - 3 elements: \(\le 3\) instructions
  - 5 elements: \(\le 12\) instructions
  - 100 elements: \(\le 1500\) instructions
  - 500 elements: \(\le 11500\) instructions
- **Primitive Stack Operations**: Implements atomic operations:
  - **Push**: `pa` (push top B to A), `pb` (push top A to B).
  - **Swap**: `sa` (swap top 2 A), `sb` (swap top 2 B), `ss` (simultaneous `sa` + `sb`).
  - **Rotate**: `ra` (shift up A), `rb` (shift up B), `rr` (simultaneous `ra` + `rb`).
  - **Reverse Rotate**: `rra` (shift down A), `rrb` (shift down B), `rrr` (simultaneous `rra` + `rrb`).

---

## 📋 Table of Contents

- [Key Highlights](#-key-highlights)
- [System Architecture](#-system-architecture)
- [Sorting & Verification Flow](#-sorting--verification-flow)
- [Setup & Execution](#-setup--execution)
- [Project Directory Structure](#-project-directory-structure)
- [License](#-license)

---

## 🏗️ System Architecture

```mermaid
flowchart TD
    A[Unsorted Integer String Payload] --> B[StackSort Engine - PushSwapC/main.go]
    A --> C[Verification Checker Engine - CheckerC/main.go]
    
    B --> D[PushSort Optimization Handler]
    D --> E1[Instruction Calculator - PushInstructions]
    D --> E2[Stack State Mutator - PushStacks]
    
    E1 --> F[STDOUT Instruction Stream - pa, pb, ra, sa, rra]
    E2 --> F
    F --> C
    
    C --> G[CheckerStacks Engine]
    G --> H{Stack A Sorted & Stack B Empty?}
    H -- Yes --> I[Output: OK]
    H -- No --> J[Output: KO]
```

---

## 🖥️ Live Terminal & Pipeline Verification Preview

Below is a live shell trace demonstrating StackSort generating instructions for an unsorted set of numbers and piping the instructions into `checker` for validation:

```text
$ ARG="4 67 3 1 2 9 8 5"
$ ./stacksort "$ARG"

pb
pb
pb
sa
ra
pa
pa
pa
ra

$ ./stacksort "$ARG" | ./checker "$ARG"
OK

$ ./stacksort "$ARG" | wc -l
9
```

---

## 📐 Sorting & Verification Flow

```mermaid
sequenceDiagram
    participant User
    participant Generator as stacksort CLI
    participant Checker as checker CLI

    User->>Generator: ./stacksort "2 1 3 6 5 8"
    Generator->>Generator: Initialize Stack A with integers, Stack B empty
    Generator->>Generator: Execute partitioning & chunk sorting algorithms
    Generator-->>User: Stream generated instructions (pb, ra, sa, pa, ...)
    
    User->>Checker: echo -e "sa\nrra\npa" | ./checker "2 1 3 6 5 8"
    Checker->>Checker: Execute input instructions sequentially on Stack A/B
    alt Stack A is Sorted & Stack B is Empty
        Checker-->>User: Output "OK"
    else Invalid Order or Non-Empty Stack B
        Checker-->>User: Output "KO"
    end
```

---

## 🚀 Setup & Execution

### Prerequisites

- **Go**: Version 1.20 or newer installed.

---

### Build & Run

1. **Clone Repository**:
   ```bash
   git clone https://github.com/sahmedhusain/stacksort.git
   cd stacksort
   ```

2. **Compile Binaries**:
   ```bash
   go build -o stacksort PushSwapC/main.go
   go build -o checker CheckerC/main.go
   ```

3. **Generate Sorting Instructions**:
   ```bash
   ./stacksort "2 1 3 6 5 8"
   ```

4. **Verify Instruction Correctness with Checker**:
   ```bash
   ARG="2 1 3 6 5 8"; ./stacksort "$ARG" | ./checker "$ARG"
   ```
   *Expected Output: `OK`*

5. **Measure Instruction Count**:
   ```bash
   ARG="2 1 3 6 5 8"; ./stacksort "$ARG" | wc -l
   ```

---

## 📂 Project Directory Structure

```
stacksort/
├── go.mod                     # Go module manifest (module stacksort)
├── README.md                  # Documentation
├── test.sh                    # Verification test script
├── PushSwapC/
│   └── main.go                # Entrypoint for instruction generator binary
├── CheckerC/
│   └── main.go                # Entrypoint for validation checker binary
└── Functions/
    ├── PushSort/              # Core sorting algorithm routines
    ├── PushInstructions/      # Instruction formatting and output string builders
    ├── PushStacks/            # Generator stack operations (push, swap, rotate)
    └── CheckerStacks/         # Checker stack operations and verification rules
```

---

## 📄 License

Distributed under the MIT License. See [LICENSE](LICENSE.md) for details.
