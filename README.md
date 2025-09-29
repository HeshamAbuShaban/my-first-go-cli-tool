# My First Go CLI Tool 🚀

A lightweight command‑line application written in **Go**, built as my first step into the world of CLI development.  
The project has grown from a simple skeleton into a structured CLI with **subcommands, flags, and help messages**.

---

## ✨ Features
- Modular Go codebase with multiple subcommands
- Each subcommand has its own **flags** and `--help` usage
- Global `--help` shows available commands
- Beginner‑friendly structure for learning CLI design
- Fast and portable binary once built
- A foundation for future, more advanced CLI tools

---

## 📦 Installation

Clone the repository:
```bash
git clone https://github.com/HeshamAbuShaban/my-first-go-cli-tool.git
cd my-first-go-cli-tool
```

Build the binary:
```bash
go build -o mycli
```

Optionally, install it globally:
```bash
go install
```

Make sure your Go bin directory (usually `%USERPROFILE%\go\bin` on Windows) is in your PATH.

---

## 🚀 Usage

### Global Help
```bash
mycli --help
```

### Subcommands

#### Greet
```bash
mycli greet --name Hesham
mycli greet --name Hesham --shout
mycli greet --help
```

#### Sum
```bash
mycli sum --a 5 --b 7
mycli sum --help
```

#### Time
```bash
mycli time
mycli time --utc
mycli time --format "2006-01-02 15:04"
mycli time --help
```

If you’re still in development, you can also run it directly:
```bash
go run . greet --name Hesham
```

---

## 🛠️ Development Notes
- This is my **first Go project** exploring CLI concepts.
- The repo now demonstrates:
  - Subcommand dispatching
  - Scoped flag parsing with `flag.NewFlagSet`
  - Custom usage messages
- Future improvements may include:
  - Global flags (e.g. `--verbose`)
  - Better error handling
  - Unit tests for each subcommand

---

## 📜 License
This project is open source and available under the [MIT License](LICENSE).
