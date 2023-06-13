# Basic Message Queue in Go

This is a basic implementation of a message queue system written in Go. It demonstrates a fundamental asynchronous communication protocol, allowing for a constant stream of data between applications.

## Project Structure
The project is structured as follows:

* **message.go**: This file defines the Message struct, which represents a message to be sent via the queue.
* **queue.go**: This file contains the Queue struct and methods related to it (i.e., sending and receiving messages).
* **main.go**: This file contains the producer and consumer functions, as well as the main function that ties everything together.

