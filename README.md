
# # TCP Chat Project (NetCat-like Server-Client Architecture)

This project involves recreating the **NetCat** utility in a Server-Client architecture, with a chat application that allows communication via TCP connections. The server listens on a specific port, handling multiple clients, while the clients can connect to the server, send messages, and interact in a group chat.

The chat application must mimic the behavior of **NetCat** but with added functionality for group communication. The project is built using **Go** and demonstrates core networking, concurrency, and real-time communication skills.




## Features

### Server

- Accepts **multiple client connections** using **TCP** protocol.
- Enforces **name requirement** for clients when they connect.
- **Controls the maximum number of connections** (max 10 clients).
- Forwards **messages to all connected clients** in the chat.
- Logs messages with timestamps and the user who sent the message.
- Sends **previous chat logs** to new clients when they join.
- **Notifies all clients** when someone joins or leaves the chat.
- Handles error conditions gracefully.
- If no port is specified, defaults to port **8989**.

### Client

- Connects to the server, entering the **name** and starting the chat.
- Sends **messages** to the server that are broadcast to all connected clients.
- Receives messages from other clients and displays them in real time.
- If the client leaves, the server notifies the rest of the clients.

### Communication

- Messages are formatted as: `[timestamp][username]: [message]`.
- The server notifies the clients when someone **joins** or **leaves** the chat.
- Empty messages should not be broadcasted.

### Example Interaction

**Server Output:**


**Client1 (Yenlik) Interaction:**


**Client2 (Lee) Interaction:**

## Instructions

### Running the Project

1. To run the server:
    ```
    $ go run .
    Listening on the port: 8989
    ```

2. To run the server on a specific port:
    ```
    $ go run . <port>
    Listening on the port: <port>
    ```
3. To connect as a client:
    ```
    $ nc <server-ip> <port>
    Welcome to TCP-Chat!
    [ENTER YOUR NAME]: <name>
    ```

    Enter your name and start chatting. If the server has other connected clients, you will see their messages and can communicate with them.

### Requirements

- **Go-routines** for handling multiple clients concurrently.
- Use of **channels** or **Mutexes** for synchronization.
- Maximum of **10 connections** allowed.
- Error handling for both server and client sides.
- **Test files** for unit testing the server and client connections.




## Usage/Examples

1. **Run the server**:

```bash
 $ go run .
    Listening on the port: 8989
```

2. **Connect as a client**:

```
 $ nc localhost 8989

```

3. **Chat with others**, and send/receive messages.


