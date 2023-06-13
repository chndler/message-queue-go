# Basic Message Queue in Go

This is a basic implementation of a message queue system written in Go. It demonstrates a fundamental asynchronous communication protocol, allowing for a constant stream of data between applications.

## Project Structure
The project is structured as follows:

* **message.go**: This file defines the Message struct, which represents a message to be sent via the queue.
* **queue.go**: This file contains the Queue struct and methods related to it (i.e., sending and receiving messages).
* **main.go**: This file contains the producer and consumer functions, as well as the main function that ties everything together.

## Features
The basic features of this system are:

* Asynchronous communication: Producers can produce messages and consumers can consume messages independently.
* Multiple producers/consumers: The system supports multiple producers and consumers running concurrently.
* Error handling: The system handles errors such as when the queue is full or empty.

## Getting Started
To run the project:

1. Clone the repository.
2. Navigate to the directory containing the project.
3. Run the command `go run .`

## Limitations
This is a basic and simplified message queue system designed for educational purposes, and therefore it has several limitations:

* It doesn't guarantee the order of messages.
* It doesn't provide persistent storage; messages are stored in memory and will be lost if the program shuts down.
* It doesn't handle complex scenarios like managing when producers greatly outnumber consumers or vice versa.

## Future Improvements
Future iterations of this project could consider:

* Adding persistent storage for messages.
* Implementing mechanisms to guarantee the order of messages.
* Adding better support for scenarios with many more producers than consumers or vice versa.