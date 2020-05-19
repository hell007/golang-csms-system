// https://d3js.org.cn/document/
import * as d3 from 'd3'
import d3Tip from 'd3-tip'

export function setLine(opt) {
  const config = {
    container: null,
    containerWidth: 1200,
    containerHeight: 600,
    margin: {
      top: 80,
      right: 80,
      bottom: 30,
      left: 60
    },
    colors: [
      '#6bcd07',
      '#fbd029',
      '#fe8800',
      '#fe0000',
      '#970454',
      '#62001e'
    ],
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
  const colors = options.colors
  const title = options.title
  const labelPadding = 3

  const chart = d3
    .select(options.container)
    .attr('width', options.containerWidth)
    .attr('height', options.containerHeight)

  // 设最外包层在总图上的相对位置
  let g = chart
    .append('g')
    .attr('class', 'chart-line')
    .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')')

  const timeParse = d3.timeParse('%H:%M')

  const key = Object.keys(data[0])[1]
  const serieArr = data.map(function(d) {
    return {
      key: key,
      time: timeParse(d.time),
      value: d[key]
    }
  })
  const maxValue = d3.max(serieArr, function(d) {
    return d.value
  })
  const stepValue = 50 // 用于生成背景柱
  const rangeByStep = d3.range(0, maxValue, stepValue) // 用于生成背景柱

  const names = ['优', '良', '轻度污染', '中度污染', '重度污染', '严重污染']

  // 定义x轴
  let x = d3
    .scaleTime()
    .domain([serieArr[0].time, serieArr[data.length - 1].time])
    .range([0, width])

  // 定义y轴
  let y = d3
    .scaleLinear()
    .domain([0, maxValue])
    .range([height, 0])

  let z = d3.scaleOrdinal(d3.schemeCategory10) // 通用线条的颜色

  let line = d3
    .line()
    .curve(d3.curveMonotoneX)
    .x(function(d) {
      return x(d.time)
    })
    .y(function(d) {
      return y(d.value)
    })

  // 设置tip
  let tip = d3Tip()
    .attr('class', 'd3-tip')
    .offset([-10, 0])
    .html(function(d) {
      return (
        '<strong>时间 ' +
        d3.timeFormat('%H : %M')(d.time) +
        "<br>AQI 值:</strong> <span style='color:#ffeb3b'>" +
        d.value +
        '</span>'
      )
    })

  chart.call(tip)

  // 添加长方形方块，遮罩作用
  // 用遮罩实现线动画 水平移动
  chart
    .append('defs')
    .append('clipPath')
    .attr('id', 'clip')
    .append('rect')
    .attr('height', height)
    .attr('width', 0) 
    .transition()
    .duration(1000)
    .attr('width', width)

  // 设置y轴
  g.append('g')
    .attr('class', 'axis axis--y')
    .call(d3.axisLeft(y).tickValues(d3.range(0, maxValue, stepValue)))
    .append('text')
    .attr('y', -16)
    .attr('dy', '.71em')
    .style('text-anchor', 'middle')
    .style('fill', primary)
    .text('AQI 值')

  // 设置背景柱
  g.append('g')
    .attr('class', 'line--bg')
    .selectAll('rect')
    .data(rangeByStep)
    .enter()
    .append('rect')
    .attr('stroke', 'none')
    .attr('stroke-width', 0)
    .attr('fill', function(d, i) {
      return colors[i]
    })
    .attr('x', 1)
    .attr('width', width)
    .attr('height', function(d, i) {
      if (i !== rangeByStep.length - 1) {
        return y(maxValue - stepValue)
      } else {
        return y(rangeByStep[rangeByStep.length - 1])
      }
    })
    .attr('y', function(d, i) {
      if (i !== rangeByStep.length - 1) {
        return y(rangeByStep[i + 1])
      } else {
        return 0
      }
    })

  // 设置背景柱文字
  g.append('g')
    .attr('class', 'line-bg-text')
    .selectAll('.ylabel') // 生成右边文字
    .data(rangeByStep)
    .enter()
    .append('text')
    .attr('class', 'ylabel')
    .attr('fill', '#fff')
    .attr('font-size', '24px')
    .attr('x', width / 2)
    .attr('y', function(d, i) {
      if (i !== rangeByStep.length - 1) {
        return y(rangeByStep[i + 1])
      } else {
        return 0
      }
    })
    .attr('dy', function(d, i) {
      if (i !== rangeByStep.length - 1) {
        return y(maxValue - stepValue) / 2
      } else {
        return y(rangeByStep[rangeByStep.length - 1]) / 2 + 8
      }
    })
    .attr('text-anchor', 'middle')
    .text(function(d, i) {
      return names[i]
    })

  // 生成x轴
  g.append('g')
    .attr('class', 'axis axis--x')
    .attr('transform', 'translate(0,' + height + ')')
    .call(
      d3
      .axisBottom(x)
      .ticks(24)
      .tickFormat(d3.timeFormat('%H:%M'))
    )

  // xx轴背景线
  g.selectAll('.axis--x .tick') 
    .append('line')
    .attr('class', 'bg-line')
    .attr('stroke', 'rgba(255,255,255,0.5)')
    .attr('stroke-dasharray', '2,2')
    .attr('shape-rendering', 'crispEdges')
    .attr('transform', 'translate(' + 0 + ',' + -1 * height + ')')
    .attr('y2', height)
  g.select('.bg-line').remove()

  // 生成线条
  let serie = g
    .selectAll('.serie')
    .data([serieArr])
    .enter()
    .append('g')
    .attr('class', 'serie')

  // 绘画线条
  serie
    .append('path')
    .attr('clip-path', 'url(#clip)')
    .attr('class', 'line')
    .style('stroke', function(d) {
      return z(d[0].key)
    })
    .style('stroke-width', 2)
    .attr('fill', 'none')
    .attr('d', line)

  // 生成文字包层
  let label = serie
    .selectAll('.label')
    .data(function(d) {
      return d
    })
    .enter()
    .append('g')
    .on('mouseover', tip.show)
    .on('mouseout', tip.hide)
    .attr('cursor', 'pointer')
    .attr('class', 'label')
    .attr('transform', function(d, i) {
      return 'translate(' + x(d.time) + ',' + y(d.value) + ')'
    })

  // 生成数值文字
  label
    .append('text')
    .attr('dy', '.35em')
    .attr('fill', '#fff')
    .attr('text-anchor', 'middle')
    .text(function(d) {
      return d.value
    })

  // 生成背景白块
  label
    .insert('rect', 'text')
    .datum(function() {
      return this.nextSibling.getBBox()
    })
    .attr('fill', 'rgba(0,0,0,0.5)')
    .attr('rx', '5px')
    .attr('ry', '5px')
    .attr('x', function(d) {
      return d.x - labelPadding
    })
    .attr('y', function(d) {
      return d.y - labelPadding
    })
    .attr('width', function(d) {
      return d.width + 2 * labelPadding
    })
    .attr('height', function(d) {
      return d.height + 2 * labelPadding
    })

  // 标题
  chart
    .append('g')
    .attr('class', 'chart-title')
    .append('text')
    .attr('fill', '#000')
    .attr('font-weight', '700')
    .attr(
      'transform',
      'translate(' + (width / 2 + margin.left) + ',' + 20 + ')'
    )
    .attr('text-anchor', 'middle')
    .attr('x', 0)
    .attr('y', 0)
    .text(title)

}