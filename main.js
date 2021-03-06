#! /usr/local/bin/node

const commandLineArgs = require('command-line-args')
const optionDefinitions = [
  { name: 'http_port', alias: 'p', type: Number },
  { name: 'http_host', alias: 'h', type: String },
  { name: 'socket_port', alias: 'b', type: Number },
  { name: 'socket_host', alias: 'n', type: String },
]
const argvOptions = commandLineArgs(optionDefinitions)
const options = Object.assign(
  {
    http_port: 3000,
    http_host: '0.0.0.0',
    socket_port: 8080,
    socket_host: '0.0.0.0',
  },
  argvOptions
)

const server = require('express')()
const WebSocket = require('ws')

const wss = new WebSocket.Server({
  server: server,
  port: options.socket_port,
  host: options.socket_host,
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

server.listen(options.http_port, options.http_host, () => {
  const ip = require('internal-ip').v4.sync()
  console.log(
    require('boxen')(
      `
    Http Listen: http://${options.http_host}:${options.http_port}
    Http Listen: http://${ip}:${options.http_port}

    Ws Listen: ws://${options.socket_host}:${options.socket_port}
    Ws Listen: ws://${ip}:${options.socket_port}

    Dashboard: http://${options.http_host}:${options.http_port}/dashboard
    Dashboard: http://${ip}:${options.http_port}/dashboard

    Demo: http://${options.http_host}:${options.http_port}/demo
    Demo: http://${ip}:${options.http_port}/demo
    `,
      { padding: 1, margin: 1, borderColor: 'green' }
    )
  )
})
