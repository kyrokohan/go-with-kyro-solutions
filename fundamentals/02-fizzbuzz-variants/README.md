# 02 – FizzBuzz Variants

Build classic *FizzBuzz* (and a few twists) to practice loops and conditionals.

---

## 🌟 Required goals

| # | Requirement | Quick check |
|---|-------------|-------------|
| 1 | Print numbers **1–100**, replacing multiples of **3** with `Fizz`, multiples of **5** with `Buzz`, and multiples of both with `FizzBuzz`. | `go run .` → `… 14 FizzBuzz 16 … 98 Fizz 99 Buzz 100` |
| 2 | Accept a **`--max`** flag (long form only) to choose the upper bound (default **100**). | `go run . --max 20` → output ends at `20` |

---

## 🚀 Stretch goals (optional)

1. Add **`--fizz`** and **`--buzz`** flags to customise the divisors (default 3 and 5).
2. Add a **`--format`** flag (`plain`, `csv`, `json`) to choose the output style—for example, `csv` prints `1,2,Fizz,4,…` and `json` prints `["1","2","Fizz","4",…]`.
3. Colorize `Fizz` in green and `Buzz` in yellow using ANSI escape codes.
4. Write table-driven unit tests for your substitution logic (`go test`). 