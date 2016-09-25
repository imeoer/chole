<style scoped>
  canvas {
    width: 100%;
    height: 150px;
    position: absolute;
    bottom: -50px;
    z-index: 0;
  }
  #line1 {
    opacity: 0.2;
  }
  #line2 {
    opacity: 0.3;
  }
</style>

<template>
  <div class="chart-wrap">
    <canvas id="line1"></canvas>
    <canvas id="line2"></canvas>
  </div>
</template>

<script>
const draw = (id, speed, density, color) => {
  const TAU = Math.PI * 2
  const res = 0.005 // percentage of screen per x segment
  const outerScale = 0.1 / density
  let inc = 0

  const c = document.getElementById(id)
  const ctx = c.getContext('2d')

  const grad = ctx.createLinearGradient(0, 0, 0, c.height * 4)
  grad.addColorStop(0, 'rgba(223, 191, 32, 1)')
  grad.addColorStop(1, 'rgba(0, 0, 0, 0)')

  const drawWave = () => {
    const w = c.offsetWidth
    const h = 100
    const cx = w * 0.5
    const cy = h * 0.5
    ctx.clearRect(0, 0, w, h)
    const segmentWidth = w * res
    ctx.fillStyle = color
    ctx.beginPath()
    ctx.moveTo(0, cy)
    for (let i = 0, endi = 1 / res; i <= endi; i++) {
      const _y = cy + Math.sin((i + inc) * TAU * res * density) * cy * Math.sin(i * TAU * res * density * outerScale)
      const _x = i * segmentWidth
      ctx.lineTo(_x, _y)
    }
    ctx.lineTo(w, h)
    ctx.lineTo(0, h)
    ctx.closePath()
    ctx.fill()
  }

  const loop = () => {
    inc -= speed
    drawWave()
    requestAnimationFrame(loop)
  }

  loop()
}

export default {
  ready() {
    draw('line1', 0.2, 5, '#0be242')
    draw('line2', 0.4, 8, '#0be242')
  }
}
</script>
