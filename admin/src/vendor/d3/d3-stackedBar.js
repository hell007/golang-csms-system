// https://d3js.org.cn/document/
import * as d3 from 'd3'
import d3Tip from 'd3-tip'

export function setStackedBar(opt) {
  const config = {
    container: null,
    containerWidth: 1200,
    containerHeight: 500,
    margin: {
      top: 80,
      right: 80,
      bottom: 30,
      left: 60
    },
    colors: [
      '#19d4ae',
      '#5867c3',
      '#f7ba2a',
      '#13ce66',
      '#8acc6d',
      '#20a0ff',
      '#fa6e86'
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

  const chart = d3
    .select(options.container)
    .attr('width', options.containerWidth)
    .attr('height', options.containerHeight)

  let x = d3
    .scaleBand()
    .rangeRound([0, width])
    .paddingInner(0.1)

  let y = d3.scaleLinear().rangeRound([height, 0])

  let z = d3.scaleOrdinal().range(d3.schemeCategory10)

  let keys = Object.keys(data[0]).slice(1)
  const names = {
    q1: '第一季度',
    q2: '第二季度',
    q3: '第三季度',
    q4: '第四季度'
  }

  let series = d3
    .stack()
    .keys(keys)
    .offset(d3.stackOffsetDiverging)(data)

  let tip = d3Tip() // 设置tip
    .attr('class', 'd3-tip')
    .offset([-10, 0])
    .html(function(d) {
      let total = d.data.q1 + d.data.q2 + d.data.q3 + d.data.q4
      return (
        '<strong>' +
        d.data.date +
        '年</strong><br>' +
        '<span style="color:' +
        z(keys[0]) +
        '">' +
        names.q1 +
        ': ' +
        d.data.q1 +
        ' 万</span><br>' +
        '<span style="color:' +
        z(keys[1]) +
        '">' +
        names.q2 +
        ': ' +
        d.data.q2 +
        ' 万</span><br>' +
        '<span style="color:' +
        z(keys[2]) +
        '">' +
        names.q3 +
        ': ' +
        d.data.q3 +
        ' 万</span><br>' +
        '<span style="color:' +
        z(keys[3]) +
        '">' +
        names.q4 +
        ': ' +
        d.data.q4 +
        ' 万</span><br>' +
        '<span style="color:#fff">年总: ' +
        total +
        ' 万</span>'
      )
    })

  chart.call(tip)

  x.domain(
    data.map(function(d) {
      return d.date
    })
  )
  y.domain([
    d3.min(series, function(serie) {
      return d3.min(serie, function(d) {
        return d[0]
      })
    }),
    d3.max(series, function(serie) {
      return d3.max(serie, function(d) {
        return d[1]
      })
    })
  ])

  let g = chart
    .append('g')
    .attr('class', 'chart-stacked')
    .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')') // 设最外包层在总图上的相对位置

  g.append('g') // 画柱状图
    .selectAll('g')
    .data(series)
    .enter()
    .append('g')
    .attr('fill', function(d) {
      return z(d.key)
    })
    .selectAll('rect')
    .data(function(d) {
      return d
    })
    .enter()
    .append('rect')
    .on('mouseover', tip.show)
    .on('mouseout', tip.hide)
    .attr('width', x.bandwidth)
    .attr('cursor', 'pointer')
    .attr('x', function(d) {
      return x(d.data.date)
    })
    .attr('height', 0)
    .attr('y', y(0))
    .transition()
    .duration(750)
    .delay(function(d, i) {
      return i * 10
    })
    .attr('y', function(d) {
      return y(d[1])
    })
    .attr('height', function(d) {
      return y(d[0]) - y(d[1])
    })

  g.append('g') // 画x轴
    .attr('class', 'axis')
    .attr('transform', 'translate(0,' + y(0) + ')')
    .call(d3.axisBottom(x))

  g.append('g') // 画y轴
    .attr('class', 'axis')
    .call(d3.axisLeft(y).ticks(null, 's'))
    .append('text')
    .attr('x', 2)
    .attr('y', y(y.ticks().pop()) + 0.5)
    .attr('dy', '0.32em')
    .attr('fill', '#000')
    .attr('font-weight', 'bold')
    .attr('text-anchor', 'start')
    .text('利润(万)')

  let legend = g
    .append('g') // 画legend
    .attr('font-family', 'sans-serif')
    .attr('font-size', 10)
    .attr('transform', 'translate(65,0)')
    .attr('text-anchor', 'end')
    .selectAll('g')
    .data(keys.slice())
    .enter()
    .append('g')
    .attr('transform', function(d, i) {
      return 'translate(0,' + i * 20 + ')'
    })

  legend
    .append('rect')
    .attr('x', width - 19)
    .attr('width', 19)
    .attr('height', 19)
    .attr('fill', z)

  legend
    .append('text')
    .attr('x', width - 24)
    .attr('y', 9.5)
    .attr('dy', '0.32em')
    .text(function(d) {
      return names[d]
    })

  chart
    .append('g') // 输出标题
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