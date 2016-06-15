const net = require('net')
const Manage = require('./manage')
const Duplex = require('./duplex')

const parseDomain = (chunk) => {
  const header = String(chunk.slice(0, chunk.indexOf('\r\n\r\n') - 1)).toLowerCase()
  const matchs = (/\r\nhost:(.*)\r\n/).exec(header)
  if (matchs && matchs[1]) {
    return matchs[1].trim()
  }
}

const manage = new Manage()

net.createServer((from) => {
  manage.setIdle(from)
  const duplex = new Duplex()
  duplex.writable.pipe(from)
  from.on('data', (chunk) => {
    const domain = parseDomain(chunk)
    duplex.readable.write(chunk)
    if (manage.hasIdle(from)) {
      manage.pipeFrom(duplex)
      manage.delIdle(from)
    }
  }).on('end', (chunk) => {
    duplex.readable.end(chunk)
  }).on('error', (err) => {
    console.log('[from]', err.message)
  })
}).listen(80)

net.createServer((to) => {
  to.on('error', (err) => {
    console.log('[to]', err.message)
  })
  manage.pipeTo(to)
}).listen(81)
