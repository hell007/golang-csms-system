// https://d3js.org.cn/document/
import * as d3 from 'd3'
import d3Tip from 'd3-tip'

export function setPoints(opt) {
  const config = {
    container: null,
    containerWidth: 1200,
    containerHeight: 600,
    margin: {
      top: 80,
      right: 60,
      bottom: 80,
      left: 60
    },
    bgs: ['#ccc', '#ddd'],
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
  const bgs = options.bgs
  const title = options.title

  const chart = d3
    .select(options.container)
    .attr('width', options.containerWidth)
    .attr('height', options.containerHeight)

  let x = d3
    .scaleLinear()
    .domain([
      0,
      d3.max(data, function(d) {
        return d[0]
      })
    ])
    .range([0, width])
  let y = d3
    .scaleLinear()
    .rangeRound([0, height])
    .domain([
      d3.max(data, function(d) {
        return d[1]
      }),
      0
    ])
  let z = d3.scaleOrdinal(d3.schemeCategory10) // 通用线条的颜色

  let g = chart
    .append('g')
    .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')') // 设最外包层在总图上的相对位置

  let tip = d3Tip() // 设置tip
    .attr('class', 'd3-tip')
    .offset([-10, 0])
    .html(function(d) {
      return (
        "<strong>运动年限:</strong><span style='color:#ffeb3b'> " +
        d[0] +
        " </span><br><strong>健康指数:</strong><span style='color:#ffeb3b'> " +
        d[1] +
        ' </span>'
      )
    })

  chart.call(tip)

  g.append('g') // 设置x轴
    .attr('class', 'axis axis--x')
    .attr('transform', 'translate(0,' + height + ')')
    .call(d3.axisBottom(x))
    .append('text')
    .attr('x', width)
    .attr('y', 26)
    .attr('dy', '.71em')
    .style('text-anchor', 'end')
    .style('fill', '#000')
    .text('激烈运动年限 (年)')

  g.append('g') // 设置y轴
    .attr('class', 'axis axis--y')
    .call(d3.axisLeft(y))
    .append('text')
    .attr('y', -16)
    .attr('dy', '.71em')
    .style('text-anchor', 'start')
    .style('fill', '#000')
    .text('健康指数 (分)')

  g.append('g') // 输出点
    .selectAll('circle')
    .attr('class', 'data')
    .data(data)
    .enter()
    .append('circle')
    .on('mouseover', tip.show)
    .on('mouseout', tip.hide)
    .attr('cursor', 'pointer')
    .attr('fill', function(d) {
      return z(d[1])
    })
    .attr('cx', function(d) {
      return x(d[0])
    })
    .attr('cy', function(d) {
      return y(d[1])
    })
    .attr('r', 0)
    .transition()
    .duration(750)
    .delay(function(d, i) {
      return i * 10
    })
    .attr('r', 10)
  chart
    .append('g') // 输出标题
    .attr('class', 'chart--title')
    .append('text')
    .attr('fill', '#000')
    .attr('font-size', '16px')
    .attr('font-weight', '700')
    .attr('text-anchor', 'middle')
    .attr('x', options.containerWidth / 2)
    .attr('y', 20)
    .text(title)

}