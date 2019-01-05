# go-can-playground

[CAN bus](https://en.wikipedia.org/wiki/CAN_bus) is frame based protocol widely used in automotive, allowing to design single point of failure proof systems. There is no host system and all messages are broadcast.
Every single module in your modern car is handling communication in that way. That enables, us - geeks, doing amazing things there. 

There are many ways and implementations for CAN bus. For hacking purposes, I am using [MCP2515](https://www.microchip.com/wwwproducts/en/en010406) connected trough SPI to ARM based Linux system. Linux kernel comes with can module implementation, [linux-can](https://github.com/linux-can), which binds CAN data wire into Linux's network interface.

Golang because I like simplicity of this language. However, I would prefer the embedded programming, but since Golang does not provide it (there is promising, LLVM based Golang compiler - [tinygo](https://github.com/aykevl/tinygo) rising), I have to run on managed system.

## Table of content

### Command line utility `can-loopback-test`

Receiving the all frames from CAN bus, in the meantime sending debug message every 1s.