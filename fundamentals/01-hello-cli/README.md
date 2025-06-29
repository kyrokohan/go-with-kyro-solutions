# 01 – Hello CLI

Welcome to your very first Go exercise!  
You’ll build a tiny command-line program that greets the user by name.

---

## 🌟 Required goal

| # | Requirement | Quick check |
|---|-------------|-------------|
| 1 | Accept a **`--name`** flag (long form only) and print `Hello, <name>!` to stdout. | `go run . --name Alice` → `Hello, Alice!` |

That’s it—keep it as small and tidy as possible.

---

## 🚀 Stretch goals (optional, pick any you like)

1. Add a shorthand **`-n`** flag that does the same thing as `--name`.
2. Default to **“World”** if no flag is supplied (`go run .` → `Hello, World!`).
3. Colorize the name in cyan using an ANSI escape sequence.
4. Override `flag.Usage` to print friendlier `-h / --help` output.
5. Respect the `NO_COLOR` environment variable (skip colors when it’s set).
6. Add a `--version` flag that prints `hello-cli v0.1.0`.
7. Write table-driven unit tests for your greeting logic (`go test`).
