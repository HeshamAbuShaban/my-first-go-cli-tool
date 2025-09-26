# My First Go CLI Tool 🚀

A lightweight command‑line application written in **Go**, built as my first step into the world of CLI development.  
This project started with a simple skeleton and currently supports three core commands:

- **`greet`** → prints a friendly greeting (try it with your name!)  
- **`sum`** → adds numbers passed as arguments  
- **`time`** → displays the current system time  

---

## ✨ Features
- Modular Go codebase with multiple commands
- Beginner‑friendly structure for learning CLI design
- Fast and portable binary once built
- A foundation for future, more advanced CLI tools

---

## 📦 Installation

Clone the repository:
```bash
git clone https://github.com/HeshamAbuShaban/my-first-go-cli-tool.git
cd my-first-cli-tool
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

Run the tool with one of the available commands:

```bash
# Greet someone
mycli greet Hesham

# Add numbers
mycli sum 5 7 12

# Show current time
mycli time
```

If you’re still in development, you can also run it directly:
```bash
go run . greet Hesham
```

---

## 🛠️ Development Notes
- This is my **first Go project** exploring CLI concepts.
- The repo is intentionally simple, focusing on structure and learning.
- Future improvements may include flags, subcommands, and richer output.

---

## 📜 License
This project is open source and available under the [MIT License](LICENSE).
