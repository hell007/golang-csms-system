// https://d3js.org.cn/document/
import * as d3 from 'd3'
import d3Tip from 'd3-tip'

export function setGroupedBar(opt) {
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

  let x0 = d3
    .scaleBand()
    .rangeRound([0, width])
    .paddingInner(0.1)

  let x1 = d3.scaleBand().padding(0.05)

  let y = d3.scaleLinear().rangeRound([height, 0])

  let z = d3
    .scaleOrdinal()
    .range(['#98abc5', '#7b6888', '#a05d56', '#ff8c00'])

  let keys = Object.keys(data[0]).slice(1)
  const names = {
    q1: '第一季度',
    q2: '第二季度',
    q3: '第三季度',
    q4: '第四季度'
  }

  let tip = d3Tip() // 设置tip
    .attr('class', 'd3-tip')
    .offset([-10, 0])
    .html(function(d) {
      return (
        '<strong>' +
        names[d.key] +
        "<br>营收:</strong> <span style='color:#ffeb3b'>" +
        d.value +
        ' 万</span>'
      )
    })

  chart.call(tip)
  x0.domain(
    data.map(function(d) {
      return d.date
    })
  )
  x1.domain(keys).rangeRound([0, x0.bandwidth()])
  y.domain([
    0,
    d3.max(data, function(d) {
      return d3.max(keys, function(key) {
        return d[key]
      })
    })
  ]).nice()

  let g = chart
    .append('g')
    .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')') // 设最外包层在总图上的相对位置

  g.append('g') // 画柱状图
    .selectAll('g')
    .data(data)
    .enter()
    .append('g')
    .attr('transform', function(d) {
      return 'translate(' + x0(d.date) + ',0)'
    })
    .selectAll('rect')
    .data(function(d) {
      return keys.map(function(key) {
        return {
          key: key,
          value: d[key]
        }
      })
    }) // 把json数据转格式
    .enter()
    .append('rect')
    .on('mouseover', tip.show)
    .on('mouseout', tip.hide)
    .attr('x', function(d) {
      return x1(d.key)
    })
    .attr('cursor', 'pointer')
    .attr('width', x1.bandwidth())
    .attr('fill', function(d) {
      return z(d.key)
    })
    .attr('height', 0)
    .attr('y', height)
    .transition()
    .duration(750)
    .delay(function(d, i) {
      return i * 10
    })
    .attr('y', function(d) {
      return y(d.value)
    })
    .attr('height', function(d) {
      return height - y(d.value)
    })

  g.append('g') // 画x轴
    .attr('class', 'axis')
    .attr('transform', 'translate(0,' + height + ')')
    .call(d3.axisBottom(x0))

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
    .text('营收(万)')

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
    .attr('class', 'grouped-bar--title')
    .append('text')
    .attr('fill', '#000')
    .attr('font-size', '16px')
    .attr('font-weight', '700')
    .attr('text-anchor', 'middle')
    .attr('x', options.containerWidth / 2)
    .attr('y', 20)
    .text(title)
}