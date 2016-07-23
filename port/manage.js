module.exports = class Manage {
  constructor() {
    this.froms = []
    this.tos = []
    this.idles = new Map()
  }
  _pipe() {
    if (this.froms.length && this.tos.length) {
      const from = this.froms.shift()
      const to = this.tos.shift()
      from.pipe(to).pipe(from)
      console.log(this.froms.length, this.tos.length)
    }
  }
  pipeFrom(from) {
    this.froms.push(from)
    this._pipe()
  }
  pipeTo(to) {
    this.tos.push(to)
    this._pipe()
  }
  setIdle(from) {
    this.idles.set(from, true)
  }
  delIdle(from) {
    this.idles.delete(from)
  }
  hasIdle(from) {
    return this.idles.has(from)
  }
  lessIdle() {
    return this.idles.size < 5
  }
}
