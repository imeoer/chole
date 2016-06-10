const net = require('net')
const dns = require('dns')

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

const parseDomain = (chunk) => {
  const header = String(chunk.slice(0, chunk.indexOf('\r\n\r\n') - 1)).toLowerCase()
  const matchs = (/\r\nhost:(.*)\r\n/).exec(header)
  if (matchs && matchs[1]) {
    return matchs[1].trim()
  }
}

const clientServer = net.createServer((clientSocket) => {
  socketQueue.pushClient(clientSocket)
  clientSocket.on('data', (chunk) => {
    console.log(parseDomain(chunk))
  }).on('error', (err) => {
    console.error(3, err.message)
  }).on('close', () => {
  })
}).on('error', (err) => {
  console.error(4, err.message)
}).listen(80, () => {
  address = clientServer.address()
  console.log('Listening on %s', address.port)
})
