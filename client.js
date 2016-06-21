const net = require('net')
const Manage = require('./manage')
const Duplex = require('./duplex')

const manage = new Manage()

const uuid = '12345'

const connectFrom = () => {
  const duplex = new Duplex()
  net.connect(81, '127.0.0.1', function() {
    const from = this
    duplex.writable.pipe(from)
    manage.setIdle(from)
    from.write(`${uuid} hello\n`)
    connectTo(to => manage.pipeTo(to))
  }).on('data', function(chunk) {
    const from = this
    duplex.readable.write(chunk)
    if (manage.hasIdle(from)) {
      manage.pipeFrom(duplex)
      if (manage.lessIdle()) connectFrom()
      manage.delIdle(from)
    }
  }).on('end', (chunk) => {
    duplex.readable.end(chunk)
  }).on('error', (err) => {
    console.log('[from]', err.message)
  }).on('close', () => {
    if (manage.lessIdle()) connectFrom()
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
