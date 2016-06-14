const net = require('net')
const stream = require('stream')

const parseDomain = (chunk) => {
  const header = String(chunk.slice(0, chunk.indexOf('\r\n\r\n') - 1)).toLowerCase()
  const matchs = (/\r\nhost:(.*)\r\n/).exec(header)
  if (matchs && matchs[1]) {
    return matchs[1].trim()
  }
}

class Duplex extends stream.Duplex {
  constructor(opts) {
    super(opts)
    this.readable = new stream.PassThrough()
    this.writeable = new stream.PassThrough()
    this.readable.on('data', (chunk) => {
      this.push(chunk)
    }).on('end', () => {
      this.push(null)
    })
  }
  _read(size) {}
  _write(chunk, encoding, next) {
    this.writeable.end(chunk)
  }
}

class Pipe {
  constructor() {
    this.froms = []
    this.tos = []
  }
  tryPipe() {
    if (this.froms.length && this.tos.length) {
      const from = this.froms.shift()
      const to = this.tos.shift()
      from.pipe(to).pipe(from)
    }
  }
  setFrom(from) {
    this.froms.push(from)
    this.tryPipe()
  }
  setTo(to) {
    this.tos.push(to)
    this.tryPipe()
  }
}

const pipe = new Pipe()

net.createServer((from) => {
  const duplex = new Duplex()
  duplex.writeable.pipe(from)
  duplex.on('error', () => {})
  from.on('data', (chunk) => {
    const domain = parseDomain(chunk)
    if (domain) {
      if (!from.used) {
        pipe.setFrom(duplex)
        from.used = true
      }
      duplex.readable.write(chunk)
    }
  }).on('end', () => {
    duplex.readable.end()
  }).on('error', (err) => {
    console.log('[from]', err.message)
  })
}).listen(80)

net.createServer((to) => {
  to.on('error', (err) => {
    console.log('[to]', err.message)
  })
  pipe.setTo(to)
}).listen(81)
