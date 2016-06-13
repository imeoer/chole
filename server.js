const net = require('net')
const stream = require('stream')
const Duplex = stream.Duplex
const PassThrough = stream.PassThrough

const parseDomain = (chunk) => {
  const header = String(chunk.slice(0, chunk.indexOf('\r\n\r\n') - 1)).toLowerCase()
  const matchs = (/\r\nhost:(.*)\r\n/).exec(header)
  if (matchs && matchs[1]) {
    return matchs[1].trim()
  }
}

class SocketDuplex extends Duplex {
  constructor(opts) {
    super(opts)
    this.readable = new PassThrough()
    this.writeable = new PassThrough()
    this.readable.on('data', (chunk) => {
      this.push(chunk)
    }).on('end', () => {
      this.push(null)
    })
  }
  _read(size) {}
  _write(chunk, encoding, next) {
    this.writeable.write(chunk, encoding, next)
  }
}

class SocketQueue {
  constructor() {
    this.clientSockets = []
    this.proxySockets = []
  }
  _tryPipe() {
    if (this.clientSockets.length && this.proxySockets.length) {
      const clientSocket = this.clientSockets.shift()
      const proxySocket = this.proxySockets.shift()
      clientSocket.pipe(proxySocket).pipe(clientSocket)
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

const proxyServer = net.createServer((proxySocket) => {
  proxySocket.on('data', (chunk) => {
    if (!proxySocket.used) {
      const domain = String(chunk)
      if (domain == 'a.chole.io') {
        proxySocket.write('ok')
        socketQueue.pushProxy(proxySocket)
      } else {
        proxySocket.write('no')
      }
      proxySocket.used = true
    }
  }).on('error', (err) => {
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
  const socketDuplex = new SocketDuplex()
  socketDuplex.writeable.pipe(clientSocket)
  clientSocket.on('data', (chunk) => {
    const domain = parseDomain(chunk)
    console.log(domain)
    if (domain == 'a.chole.io') {
      if (!clientSocket.used) {
        socketQueue.pushClient(socketDuplex)
        clientSocket.used = true
      }
      socketDuplex.readable.write(chunk)
    }
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
