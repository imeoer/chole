const net = require('net')

class Pipe {
  tryPipe() {
    if (this.from && this.to) {
      this.from.pipe(this.to).pipe(this.from)
      this.from = this.to = null
    }
  }
  setFrom(from) {
    this.from = from
    this.tryPipe()
  }
  setTo(to) {
    this.to = to
    this.tryPipe()
  }
}

const pipe = new Pipe()

const connectFrom = () => {
  const from = net.connect(81, '127.0.0.1', () => {
    pipe.setFrom(from)
    connectTo((to) => {
      pipe.setTo(to)
    })
  }).on('error', (err) => {
    console.log('[from]', err.message)
  }).on('close', () => {
    connectFrom()
  })
}

const connectTo = (cb) => {
  const to = net.connect(8080, '127.0.0.1', () => {
    cb(to)
  }).on('error', (err) => {
    console.log('[to]', err.message)
  })
}

connectFrom()
