const server = require('express')()
const WebSocket = require('ws')

const wss = new WebSocket.Server({
  server: server,
  port: 8080,
})

wss.on('connection', function (ws) {
  ws.on('message', function (message) {
    console.log('received', message)
    ws.send(message)
  })
})

wss.on('message', function (mes) {
  console.log(mes)
})

server.get('/', (req, res) => {
  res.sendFile(__dirname + '/index.html')
})

server.listen(3000)
