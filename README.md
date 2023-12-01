# go-chat

A simple Go chat application using Web Socket.

## How to Use

1. Clone this repository.

2. Run the application.

```sh
go run main.go
```

3. Try it out using client application like Postman.

## How to Test the Application using Postman

1. Open Postman.

2. Choose `File` > `New...` > `WebSocket`.

3. Create a new request called `first user`.

4. Enter the Web Socket URL with this format `ws://localhost:1323/ws?username=YOUR_USERNAME&room=ROOM_NUMBER`.

Example: `ws://localhost:1323/ws?username=nadir&room=1`

5. Save the request then try to connect to the Web Socket by clicking `Connect` button.

6. Create another request called `second user`.

7. Enter the Web Socket URL with this format `ws://localhost:1323/ws?username=YOUR_USERNAME&room=ROOM_NUMBER`.

Example: `ws://localhost:1323/ws?username=ridan&room=2`

8. Save the request then try to connect to the Web Socket by clicking `Connect` button.

9. In first request, try to send a message to the user in room 2.

```json
{
  "username": "nadir",
  "room": "2",
  "text": "hello ridan!"
}
```

10. The message will be received in second request. In the second request, try to send a message back to the user in room 1.

```json
{
  "username": "ridan",
  "room": "1",
  "text": "hello nadir, I am good"
}
```
