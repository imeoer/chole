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

class SocketPipe {
  constructor() {
    this.clientMap = {}
    this.proxyMap = {}
  }
  _tryPipe(domain) {
    const clients = this.clientMap[domain]
    const proxys = this.proxyMap[domain]
    if (clients && clients.length && proxys && proxys.length) {
      const clientSocket = clients.shift()
      const proxySocket = proxys.shift()
      clientSocket.pipe(proxySocket).pipe(clientSocket)
      console.log('PROXIED')
    }
  }
  pushClient(domain, socket) {
    const clients = this.clientMap[domain]
    this.clientMap[domain] = clients ? clients.concat(socket) : [socket]
    this._tryPipe(domain)
  }
  pushProxy(domain, socket) {
    const proxys = this.proxyMap[domain]
    this.proxyMap[domain] = proxys ? proxys.concat(socket) : [socket]
    this._tryPipe(domain)
  }
}

const socketPipe = new SocketPipe()

const proxyServer = net.createServer((proxySocket) => {
  proxySocket.on('data', (chunk) => {
    if (!proxySocket.used) {
      const domain = String(chunk)
      proxySocket.write('ok\r\n')
      socketPipe.pushProxy(domain, proxySocket)
      proxySocket.used = true
    }
  }).on('error', (err) => {
    console.error(1, err.message)
  }).on('close', () => {
    console.log('CLIENT')
    for (let item in socketPipe.clientMap) {
      console.log(`${item}: ${socketPipe.clientMap[item].length}`)
    }
    console.log('PROXY')
    for (let item in socketPipe.proxyMap) {
      console.log(`${item}: ${socketPipe.proxyMap[item].length}`)
    }
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
    if (domain) {
      if (!clientSocket.used) {
        socketPipe.pushClient(domain, socketDuplex)
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
