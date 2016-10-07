<style scoped>
  canvas {
    width: 100%;
    height: 100%;
    bottom: 0;
    z-index: 0;
  }
  #line {
  }
</style>

<template>
  <div class="chart-wrap">
    <canvas id="line"></canvas>
  </div>
</template>

<script>
import Chart from 'chart.js'

const createChart = () => {
  const ctx = document.getElementById('line')
  const startingData = {
    labels: [1, 2, 3, 4, 5, 6, 7],
    datasets: [
      {
        fillColor: 'rgba(220,220,220,0.2)',
        strokeColor: 'rgba(220,220,220,1)',
        pointColor: 'rgba(220,220,220,1)',
        pointStrokeColor: '#fff',
        data: [65, 59, 80, 81, 56, 55, 40]
      },
      {
        fillColor: 'rgba(220,220,220,0.2)',
        strokeColor: 'rgba(220,220,220,1)',
        pointColor: 'rgba(220,220,220,1)',
        pointStrokeColor: '#fff',
        data: [28, 48, 40, 19, 86, 27, 90]
      }
    ]
  }
  let latestLabel = startingData.labels[0]
  const chart = new Chart(ctx)
  setInterval(function() {
    chart.data.labels.splice(0, 1)
    chart.data.datsets.forEach(function(dataset) {
      dataset.data.splice(0, 1)
    })

    chart.update()

    chart.data.labels.push('new label')
    chart.data.datsets.forEach(function(dataset, index) {
      dataset.data.push([Math.random() * 100, Math.random() * 100], ++latestLabel)
    })

    chart.update()
  }, 100)
}

export default {
  props: {
    speed: {
      type: Number,
      required: true
    }
  },
  ready() {
    createChart()
  }
}
</script>
