# sidl

> A fast, interactive CLI for inspecting Twilio SIDs from your terminal.

`sidl` is a lightweight command-line tool that helps developers quickly
identify and retrieve metadata about Twilio Service Identifiers (SIDs).
Instead of searching documentation or guessing resource types, sidl lets
you paste a SID and immediately see what it represents.

---

## ✨ Features

- **Interactive TUI** -- Clean terminal interface for exploring SID
  types.
- **Instant SID Lookup** -- Paste a SID to retrieve its associated
  resource information.
- **Batch Support** -- Analyze multiple SIDs at once.
- **Embedded Data** -- No external configuration or data files
  required; everything is compiled into the binary.
- **Cross-Platform Binaries** -- Prebuilt releases for macOS and
  Linux.

---

## 🎥 Demo

https://github.com/user-attachments/assets/ba8b54cd-7072-4c84-89ac-4490f29f95cf

---

## 📦 Installation

### Homebrew (macOS/Linux)

```bash
brew tap kwame-Owusu/sidl
brew install sidl
```

---

## 🚀 Usage

After installation, run:

```bash
sidl
```

You will enter the interactive terminal interface where you can:

- Paste a single SID
- Paste multiple SIDs for batch lookup
- Navigate resource types

---

## 🛠 Development

### Prerequisites

- Go (latest stable version recommended)

### Clone the Repository

```bash
git clone https://github.com/kwame-Owusu/sidl.git
cd sidl
```

### Build

```bash
go build -o sidl main.go
```

### Run

```bash
./sidl
```

---

## 📦 Releases

Prebuilt binaries are available on the GitHub Releases page.

**Supported platforms:**

- macOS
  - `sidl-darwin-amd64`
  - `sidl-darwin-arm64`
- Linux
  - `sidl-linux-amd64`

---

## 🤝 Contributing

Contributions are welcome and appreciated.

1.  Fork the repository

2.  Create a feature branch

    ```bash
    git checkout -b feature/my-feature
    ```

3.  Commit your changes

4.  Push to your fork

5.  Open a Pull Request

Please ensure your changes are well-tested and clearly documented.

---
