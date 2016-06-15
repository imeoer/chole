const stream = require('stream')

module.exports = class Duplex extends stream.Duplex {
  constructor(opts) {
    super(opts)
    this.readable = new stream.PassThrough()
    this.writable = new stream.PassThrough()
    this.readable.on('readable', () => {
      this._read()
    })
  }
  _read(size) {
    let chunk
    while ((chunk = this.readable.read(size)) !== null) {
      this.push(chunk)
    }
  }
  _write(chunk, encoding, next) {
    this.writable.write(chunk)
    next()
  }
}
