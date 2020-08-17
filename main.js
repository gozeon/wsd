const server = require('express')()
const WebSocket = require('ws')

const wss = new WebSocket.Server({
  server: server,
  port: 8080,
})

wss.on('connection', function (ws, req) {
  // 给他个表示别
  ws.id = req.headers['sec-websocket-key']
  console.log(ws)

  ws.on('message', function (message) {
    wss.clients.forEach((w) => {
      console.log(w.id)
    })
    ws.send(message)

    console.log(ws.id, JSON.stringify(message))
  })

  ws.on('pong', function () {
    console.log('pong')
  })
})

wss.on('close', function (ws) {})

server.get('/', (req, res) => {
  res.sendFile(__dirname + '/index.html')
})

server.listen(3000)
