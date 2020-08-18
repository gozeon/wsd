const server = require('express')()
const WebSocket = require('ws')

const wss = new WebSocket.Server({
  server: server,
  port: 8080,
})

wss.on('connection', function (ws, req) {
  // 给他个标识
  ws.id = req.headers['sec-websocket-key']

  ws.on('message', function (message) {
    wss.clients.forEach((w) => {
      // 转发除了自己的所有连接
      if (w.id !== ws.id) {
        w.send(message)
      }
    })
  })
})

server.get('/dashboard', (req, res) => {
  res.sendFile(__dirname + '/dashboard.html')
})

server.get('/demo', (req, res) => {
  res.sendFile(__dirname + '/demo.html')
})

server.listen(3000)
