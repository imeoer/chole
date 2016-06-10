const net = require('net')

class SocketQueue {
  constructor() {
    this.clientSockets = []
    this.proxySockets = []
  }
  _tryPipe() {
    if (this.clientSockets.length && this.proxySockets.length) {
      const clientSocket = this.clientSockets.shift()
      const proxyClientSocket = this.proxySockets.shift()
      clientSocket.pipe(proxyClientSocket).pipe(clientSocket)
      console.log('PROXIED')
    }
  }
  pushClient(socket) {
    this.clientSockets.push(socket)
    this._tryPipe()
  }
  pushProxy(socket) {
    this.proxySockets.push(socket)
    this._tryPipe()
  }
}

let socketQueue = new SocketQueue()

const proxyServer = net.createServer((proxyClientSocket) => {
  socketQueue.pushProxy(proxyClientSocket)
  proxyClientSocket.on('error', (err) => {
    console.error(1, err.message)
  }).on('close', () => {
  })
}).on('error', (err) => {
  console.error(2, err.message)
}).listen(81, () => {
  address = proxyServer.address()
  console.log('Listening on %s', address.port)
})

const clientServer = net.createServer((clientSocket) => {
  socketQueue.pushClient(clientSocket)
  clientSocket.on('error', (err) => {
    console.error(3, err.message)
  }).on('close', () => {
  })
}).on('error', (err) => {
  console.error(4, err.message)
}).listen(80, () => {
  address = clientServer.address()
  console.log('Listening on %s', address.port)
})
