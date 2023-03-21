# gRPC Message Service

![gRPC Logo](https://grpc.io/img/logos/grpc-icon-color.png)

## Project Overview
This project demonstrates a **simple gRPC service** to send and receive messages using various gRPC communication patterns. gRPC is a high-performance, open-source, universal RPC framework designed to make it easier to build distributed systems.

Tech Stack
Language: Golang
Framework: gRPC
Protocol: Protocol Buffers (proto3)

![Output of the Project](/assets/output.png)

## Application
This application serves as a foundation for creating more complex microservices using gRPC. The gRPC service supports four different types of RPC methods:

- Unary RPC: A simple request and response. The client sends a request to the server and waits for a response.
- Server Streaming RPC: The client sends a request to the server, and the server responds with a stream of messages.
- Client Streaming RPC: The client sends a stream of messages to the server, and the server responds with a single message.
- Bidirectional Streaming RPC: Both the client and the server send a stream of messages to each other.

In real-world scenarios, these communication patterns can be used for various purposes such as load balancing, real-time updates, chat applications, data streaming, and more.

### Example Usage

Here are some screenshots of the gRPC Message Service in action:

![Unary RPC example](/assets/1.png)

**Figure 1.** Unary RPC example.

![Server Streaming RPC example](/assets/2.png)

**Figure 2.** Server / ClientStreaming RPC example.


![Bidirectional Streaming RPC example](/assets/4.png)

**Figure 3.** Bidirectional Streaming RPC example.

![Client Streaming RPC example](/assets/3.png)

## Future of gRPC in Microservices
gRPC offers several advantages over traditional REST-based APIs, making it an excellent choice for microservices architecture:

- Performance: gRPC uses HTTP/2 as its transport protocol, which is more efficient than HTTP/1.1 used by REST. It also uses Protocol Buffers for serialization, resulting in smaller payloads and faster serialization/deserialization.
- Language-agnostic: gRPC supports multiple languages, making it easy to create microservices using different programming languages and still have them communicate seamlessly.
- Strongly-typed contracts: gRPC enforces strict contracts between services using Protocol Buffers, ensuring that services are always in sync and reducing the risk of runtime errors.
- Bi-directional streaming: gRPC's support for bi-directional streaming allows for more efficient communication patterns, particularly for applications that require real-time updates or large data transfers.

As more organizations adopt microservices architecture, gRPC is well-positioned to become a leading framework for building high-performance, scalable, and resilient distributed systems.



## To contribute to this project, please follow these steps:

- Fork the repository by clicking on the "Fork" button at the top of the page.
- Clone the forked repository to your local machine using git clone https://github.com/mudityadev/gRPC_Message_Service_Bidirectional.git
- Create a new branch for your changes using git checkout -b branch-name
- Make your changes to the code and commit them using git commit -m "commit message"
- Push your changes to your forked repository using git push origin branch-name
- Create a Pull Request (PR) by clicking on the "New pull request" button on your forked repository page.
### Thank you for your contribution! ☺️