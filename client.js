const net = require('net')

class SocketPipe {
  constructor() {
    this.clientSocket = null
    this.proxySocket = null
  }
  _tryPipe() {
    if (this.clientSocket && this.proxySocket) {
      this.proxySocket.pipe(this.clientSocket).pipe(this.proxySocket)
      this.clientSocket = null
      this.proxySocket = null
      console.log('PROXIED')
    }
  }
  setClient(socket) {
    this.clientSocket = socket
    this._tryPipe()
  }
  setProxy(socket) {
    this.proxySocket = socket
    this._tryPipe()
  }
}

const socketPipe = new SocketPipe()

const connectProxy = () => {
  let hasError = false
  let proxySocket = net.connect(81, () => {
    const handle = (err, clientSocket) => {
      if (err) {
        setTimeout(() => {
          connectServer(handle)
        }, 2000)
        return
      }
      proxySocket.write('a.chole.io')
      proxySocket.on('data', (chunk) => {
        if (!proxySocket.used) {
          const authed = (chunk.indexOf('ok\r\n') == 0)
          if (authed) {
            socketPipe.setClient(clientSocket)
            socketPipe.setProxy(proxySocket)
          } else {
            console.log('Auth failed')
          }
          proxySocket.used = true
        } else {
          if (!proxySocket.transfered) {
            connectProxy()
            proxySocket.transfered = true
          }
        }
      })
    }
    connectServer(handle)
  }).on('error', (err) => {
    hasError = true
    console.error(1, err.message)
    setTimeout(() => {
      connectProxy()
    }, 2000)
  }).on('close', () => {
    if (!hasError) {
      setTimeout(() => {
        connectProxy()
      }, 2000)
    }
  })
  return proxySocket
}

const connectServer = (callback) => {
  const clientSocket = net.connect(8080, () => {
    callback(null, clientSocket)
  }).on('error', (err) => {
    console.error(2, err.message)
    callback(err)
  }).on('close', () => {
  })
  return clientSocket
}

connectProxy()
