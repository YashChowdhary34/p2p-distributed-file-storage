# Distributed File System in Go

A peer-to-peer, content-addressable distributed file storage system written in Go. This project leverages a custom TCP-based p2p library to enable storage, retrieval, and encrypted transfer of files across nodes in a distributed network.

## Features

- **Peer-to-Peer Communication:** Uses a custom TCP transport for node-to-node communication.
- **Content-Addressable Storage:** Files are stored using a hash-based addressing scheme.
- **Encryption:** Supports AES encryption for secure file transfers.
- **Bootstrap & Discovery:** Nodes can bootstrap to a network using provided bootstrap addresses.
- **Extensibility:** Modular design allows adding new protocols or storage backends.

## Architecture & Design

The system is built around a few core components:

- **File Server:**  
  Each node runs a file server that handles storing files locally and serves them over the network on demand.

  - **Server Initialization:** In `main.go`, nodes are created on different ports (e.g., `:3000`, `:7000`, `:5000`) and connected using bootstrap addresses.
  - **Message Handling:** The server uses Go’s `gob` encoder/decoder for serializing messages such as `MessageStoreFile` and `MessageGetFile`.

- **Storage Layer:**  
  The storage layer implements file storage with a content-addressable scheme:

  - **Path Transformation:** A function (CASPathTransformFunc) converts a file key into a structured folder path based on a SHA1 hash.
  - **File Operations:** Supports writing (both plain and encrypted), reading, and deletion of files from disk.

- **Cryptography:**  
  In `crypto.go`, functions generate unique node IDs, compute MD5 hashes for file keys, and perform encryption/decryption using AES in CTR mode.
- **Peer-to-Peer Library:**  
  The project depends on an external p2p package which provides TCP transport, connection handling, and message dispatching.

## Installation

### Prerequisites

- [Go (version 1.16 or later)](https://golang.org/dl/)
- Git

### Clone the Repository

```bash
git clone <repository-url>
cd <repository-directory>
```

## Conclusion

This project demonstrates a robust, modular approach to building a peer-to-peer distributed file system in Go. With features such as secure file transfers and content-addressable storage, it serves as an effective foundation for further experimentation and development in distributed systems.
