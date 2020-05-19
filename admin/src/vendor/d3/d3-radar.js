// https://d3js.org.cn/document/
import * as d3 from 'd3'

export function setRadar(opt) {
  const config = {
    container: null,
    containerWidth: 1200,
    containerHeight: 700,
    margin: {
      top: 80,
      right: 80,
      bottom: 60,
      left: 60
    },
    primary: '',
    title: '',
    data: []
  }

  const options = Object.assign(config, opt)
  const data = options.data
  if (!options.container || data.length == 0) return

  const margin = options.margin
  const width = options.containerWidth - margin.left - margin.right
  const height = options.containerHeight - margin.top - margin.bottom
  const primary = options.primary
  const title = options.title
  const outerRadius = Math.min(width, height) * 0.5

  const chart = d3
    .select(options.container)
    .attr('width', options.containerWidth)
    .attr('height', options.containerHeight)

  let g = chart
    .append('g')
    .attr('class', 'chart-radar')
    .attr(
      'transform',
      'translate(' +
      options.containerWidth / 2 +
      ',' +
      options.containerHeight / 2 +
      ')'
    )

  const names = {
    chinese: '语文',
    math: '数学',
    physics: '物理',
    chemistry: '化学',
    english: '英语'
  }

  const keys = Object.keys(data[0]).slice(1)
  //console.log('keys==',keys)
  const series = data.map(d => {
    return keys.map(key => {
      return {
        key: key,
        className: d.className,
        value: d[key]
      }
    })
  })
  //console.log('series==',series)


  // 定义x轴
  let x = d3
    .scaleBand() 
    .range([0, 2 * Math.PI])
    .align(0)

  // 定义y轴
  let y = d3.scaleLinear().range([0, outerRadius]) 

  // 通用线条的颜色
  let z = d3.scaleOrdinal(d3.schemeCategory10) 

  const maxSerieValue = d3.max(series, function(s) {
    return d3.max(s, function(d) {
      return d.value
    })
  })

  // x与y轴的值域
  x.domain(keys) 
  y.domain([0, maxSerieValue + 10])

  // 添加圆形遮罩
  let defs = chart.append('defs') 
  defs
    .selectAll('clipPath')
    .data(d3.range(data.length))
    .enter()
    .append('clipPath')
    .attr('id', function(d, i) {
      return 'clip_' + i
    })
    .append('circle')
    .attr('r', 0)
    .transition()
    .duration(500)
    .delay(function(d, i) {
      return i * 500
    })
    .attr('r', outerRadius)
  
  // 画y轴圈圈及文字
  let yAxis = g
    .append('g') 
    .attr('text-anchor', 'start')

  let yTick = yAxis
    .selectAll('g')
    .data(y.ticks(10).reverse())
    .enter()
    .append('g')

  yTick
    .append('circle')
    .attr('fill', '#ddd')
    .attr('stroke', '#999')
    .attr('fill-opacity', 0.3)
    .attr('r', y)

  yTick
    .append('text')
    .attr('x', 6)
    .attr('y', function(d) {
      return -y(d)
    })
    .attr('dy', '0.35em')
    .attr('fill', 'none')
    .attr('stroke', '#fff')
    .attr('stroke-width', 5)
    .text(y.tickFormat(10, 'r'))

  yTick
    .append('text')
    .attr('x', 6)
    .attr('y', function(d) {
      return -y(d)
    })
    .attr('dy', '0.35em')
    .text(y.tickFormat(10, 'r'))
  
  // 绘画所有轴线条
  let tick = g
    .selectAll('.tick') 
    .data(
      keys.map(key => {
        return [{
          angle: 0,
          radius: 0
        }, {
          angle: x(key),
          radius: y(maxSerieValue + 10)
        }]
      })
    )
    .enter()
    .append('g')
    .attr('class', 'tick')
  
  // 开始绘画所有的轴线条
  tick
    .append('path') 
    .attr('class', 'tick-line')
    .style('stroke', '#fff')
    .style('stroke-width', 2)
    .attr('fill', 'none')
    .attr(
      'd',
      d3
      .lineRadial()
      .angle(function(d) {
        return d.angle
      })
      .radius(function(d) {
        return d.radius
      })
      .curve(d3.curveLinearClosed)
    )
  
  // 绘所有轴添加类型名
  g.selectAll('.tick-type') 
    .data(keys)
    .enter()
    .append('text')
    .attr('class', 'tick-type')
    .attr('text-anchor', function(d) {
      return x(d) > Math.PI ? 'end' : 'start'
    })
    .attr('x', function(d) {
      return Math.sin(x(d)) * (outerRadius + 10)
    })
    .attr('y', function(d) {
      return -1 * Math.cos(x(d)) * (outerRadius + 10)
    })
    .text(function(d) {
      return names[d]
    })

  
  // 绘画雷达线条
  let serie = g
    .selectAll('.serie') 
    .data(series)
    .enter()
    .append('g')
    .attr('class', 'serie')
    .attr('clip-path', function(d, i) {
      return 'url(#clip_' + i + ')'
    })
  
  // 开始绘画雷达线条
  serie
    .attr('class', 'radar-serie')
    .append('path') 
    .attr('class', 'radar-line')
    .style('stroke', function(d) {
      return z(d[0].className)
    })
    .attr('fill', function(d) {
      return z(d[0].className)
    })
    .attr('fill-opacity', 0.2)
    .attr(
      'd',
      d3
      .lineRadial()
      .angle(function(d) {
        return x(d.key)
      })
      .radius(function(d) {
        return y(d.value)
      })
      .curve(d3.curveLinearClosed)
    )
    .append('title') // 输出Title，mouseover显示
    .text(function(d) {
      let str = d
        .map(t => {
          return names[t.key] + ': ' + t.value + '分'
        })
        .join('\n')
      return d[0].className + '\n' + str
    })
  
  // 画legend
  let legend = chart
    .append('g') 
    .attr('font-family', 'sans-serif')
    .attr('font-size', 10)
    .attr('transform', 'translate(-65,65)')
    .attr('text-anchor', 'end')
    .selectAll('g')
    .data(data)
    .enter()
    .append('g')
    .attr('transform', function(d, i) {
      return 'translate(0,' + i * 20 + ')'
    })
    .on('click', function(d, i){
      //console.log(d.className+'===',d)
      const rs = document.querySelector('.radar-serie')
      rs[i].style.dispaly = 'none'
    })

  legend
    .append('rect')
    .attr('x', width - 19)
    .attr('width', 19)
    .attr('height', 19)
    .attr('fill', function(d) {
      return z(d.className)
    })

  legend
    .append('text')
    .attr('x', width - 24)
    .attr('y', 9.5)
    .attr('dy', '0.32em')
    .text(function(d) {
      return d.className
    })
  
  // 标题
  chart
    .append('g') 
    .attr('class', 'chart-title')
    .append('text')
    .attr('fill', '#000')
    .attr('font-size', '16px')
    .attr('font-weight', '700')
    .attr('text-anchor', 'middle')
    .attr('x', options.containerWidth / 2)
    .attr('y', 20)
    .text(title)
}