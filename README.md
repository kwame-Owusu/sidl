# sidl

## A CLI tool to retrieve information about Twilio SIDs

📌 Overview

**sidl** is a terminal-based utility designed to help developers quickly look up details about Twilio SIDs (Service Identifiers). Whether you're working with messaging services, phone numbers, or other Twilio resources, sidl provides a fast and straightforward way to get the information you need.

### ⚙️ Features

Interactive TUI: Navigate through a terminal user interface to explore available SIDs.

SID Lookup: Input a SID to retrieve its associated information.

Multi-SID Support: Paste and test multiple SIDs simultaneously for batch lookups.

### 🎥 Demo

https://github.com/user-attachments/assets/ba8b54cd-7072-4c84-89ac-4490f29f95cf

### Embedded Data: No external files required; all necessary data is embedded within the binary.

🚀 Installation
Homebrew (macOS/Linux)

```
brew tap kwame-Owusu/sidl
brew install sidl

```

WebInstall (Universal)
note: will add script to allow installing from web install

### 🧪 Usage

After installation, run sidl from your terminal:

```
sidl

```

Navigate through the interactive interface to explore and look up Twilio SIDs.

### 🛠️ Development

To build the project locally, you must:

**Clone the repository**:

```
git clone https://github.com/kwame-Owusu/sidl.git

cd sidl
```

**Build the binary**:

```
go build -o sidl main.go

```

**Run the application**:

```
./sidl

```

### 📦 Releases

Binaries for various platforms are available on the Releases
page:

macOS: sidl-darwin-amd64 or arm64

Linux: sidl-linux-amd64

### 🤝 Contributing

Contributions are welcome! Please fork the repository, create a new branch, and submit a pull request with your changes.

### 📝 License

This project is licensed under the MIT License.
