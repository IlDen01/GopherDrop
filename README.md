# GopherDrop

**GopherDrop** is a lightweight, lightning-fast peer-to-peer file transfer tool designed for local area networks (LAN). Written in Go, it eliminates the need for cloud services or USB drives when you need to move files between machines in the same WiFi or Ethernet network.

![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

```text
   ____             _               ____                  
  / ___| ___  _ __ | |__   ___ _ __|  _ \ _ __ ___  _ __  
 | |  _ / _ \| '_ \| '_ \ / _ \ '__| | | | '__/ _ \| '_ \ 
 | |_| | (_) | |_) | | | |  __/ |  | |_| | | | (_) | |_) |
  \____|\___/| .__/|_| |_|\___|_|  |____/|_|  \___/| .__/ 
             |_|                                    |_|    
```

## ✨ Features

-   🚀 **Zero Config**: No server required. Devices find each other automatically using UDP broadcasting.
-   🏷️ **Identity**: Automatically detects your computer's hostname or lets you set a custom nickname.
-   🔍 **Peer Discovery**: Scan your network and see a list of available receivers with their names and IP addresses.
-   🛡️ **Reliable Transfer**: Uses TCP for file transmission with a mandatory "OK" handshake to ensure data integrity.
-   ⚡ **Concurrent**: Built with Go routines for efficient non-blocking network I/O.

## 🛠️ How It Works

GopherDrop uses a two-stage protocol:
1.  **Discovery (UDP)**: The Sender broadcasts an `announce` message. Active Receivers listen on port `10609` and respond with a `response` message containing their identity.
2.  **Transfer (TCP)**: Once a receiver is selected, the Sender establishes a direct TCP connection on port `10610` to stream the file data.

## 🚀 Getting Started

### Quick Install (Pre-built Binaries)
You don't need Go installed to run GopherDrop. Simply download the latest binary for your operating system from the [Releases](https://github.com/IlDen01/GopherDrop/releases) page:

*   **Windows**: `gopherdrop-windows-amd64.exe`
*   **Linux**: `gopherdrop-linux-amd64`
*   **macOS (Apple Silicon)**: `gopherdrop-macos-arm64` (for M1/M2/M3 chips)
*   **macOS (Intel)**: `gopherdrop-macos-amd64`

#### 💡 Note for Linux/macOS users:
After downloading, you may need to give the binary execution permissions:
```bash
chmod +x gopherdrop-macos-arm64
```

### Alternative: Build from Source
If you prefer to build it yourself:

1.  **Prerequisites**: [Go](https://go.dev/doc/install) 1.20+
2.  **Clone & Build**:
```bash
git clone https://github.com/IlDen01/GopherDrop.git
cd GopherDrop
go build -o gopherdrop
```

## 🖥️ Usage
Run the application on at least two machines within the same network:

1.  **Start Receiver**: Run `./gopherdrop` and select option `2`.
2.  **Start Sender**: Run `./gopherdrop`, select option `1`, and enter the file path.
3.  **Select Device**: Choose the receiver from the auto-discovered list.

## 📂 Project Structure
-   `/protocol`: Defines the JSON communication schema.
-   `/utils`: Core logic for IP detection, TCP file streaming, and handshakes.
-   `main.go`: Entry point and interactive CLI menu.
-   `sender.go` / `receiver.go`: High-level network orchestration.

## 🤝 Contributing
Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/IlDen01/GopherDrop/issues).

## 📄 License
This project is [MIT](https://opensource.org/licenses/MIT) licensed.
