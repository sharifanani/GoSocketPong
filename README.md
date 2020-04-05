# GoSocketPong
A project for learning to work with sockets in Go (while learning Go)

## What the project does
The aim of the project is to do a little bit of exploratory work to learn go, while doing that in the context of network programming.

Project structure:
```
.
├── LICENSE
├── README.md
└── src
    ├── client
    │   ├── //client
    │   └── client.go
    └── server
        ├── //server
        └── server.go
```

Two small executables are compiled: the `client` and the `server`. The `server` listens for incoming network requests and 
prints them out when it gets them, and client sends them at 1 sec intervals.

### Why do this?
It's about learning GoLang, and helping develop an easily scalable VPN platform