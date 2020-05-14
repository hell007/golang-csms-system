// https://d3js.org.cn/document/
import * as d3 from 'd3'
import d3Tip from 'd3-tip'

export function setBar(opt) {
  const config = {
    container: null,
    containerWidth: 1200,
    containerHeight: 500,
    margin: {
      top: 80,
      right: 20,
      bottom: 30,
      left: 60
    },
    bgs: ['#ccc', '#ddd'],
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
  const bgs = options.bgs
  const colors = options.colors
  const title = options.title

  const chart = d3
    .select(options.container)
    .attr('width', options.containerWidth)
    .attr('height', options.containerHeight)

  // 设置x轴  
  let x = d3
    .scaleBand()
    .rangeRound([0, width])
    .padding(0.1)
    .domain(
      data.map(function(d) {
        return d.letter
      })
    )

  // 设置y轴
  let y = d3
    .scaleLinear()
    .rangeRound([height, 0])
    .domain([
      0,
      d3.max(data, function(d) {
        return d.frequency
      })
    ])

  // 设最外包层在总图上的相对位置
  let g = chart
    .append('g')
    .attr('class', 'chart-bar')
    .attr('transform', 'translate(' + margin.left + ', ' + margin.top + ')')

  // 生成x轴
  g.append('g')
    .attr('class', 'axis axis--x')
    .attr('transform', 'translate(0,' + height + ')')
    .call(d3.axisBottom(x))

  // 生成y轴
  g.append('g')
    .attr('class', 'axis axis--y')
    .call(d3.axisLeft(y).ticks(10, '%'))
    .append('text')
    .attr('y', -16)
    .attr('dy', '.71em')
    .style('text-anchor', 'middle')
    .style('fill', primary)
    .text('外出率 (%)')


  // 设置背景柱
  const barWidth = (width / data.length) * 0.5
  const stepArray = d3.ticks(0, d3.max(data, d => d.frequency), 10)

  g.append('g')
    .attr('class', 'bar--bg')
    .selectAll('rect')
    .data(d3.range(stepArray.length - 1))
    .enter()
    .append('rect')
    .attr('stroke', 'none')
    .attr('stroke-width', 0)
    .attr('fill', function(d, i) {
      return bgs[i % 2]
    })
    .attr('x', 1)
    .attr('width', width)
    .attr('height', function(d, i) {
      return y(stepArray[i]) - y(stepArray[i + 1])
    })
    .attr('y', function(d, i) {
      return y((i + 1) * stepArray[1])
    })

  // 设置tip
  let tip = d3Tip()
    .attr('class', 'd3-tip')
    .offset([-10, 0])
    .html(function(d) {
      return (
        '<strong>星期' +
        d.letter +
        "<br>外出率:</strong> <span style='color:" + primary + "'>" +
        (d.frequency * 100).toFixed(2) +
        '%</span>'
      )
    })

  chart.call(tip)

  // 画柱图
  g.selectAll('.bar')
    .data(data)
    .enter()
    .append('rect')
    .on('mouseover', tip.show)
    .on('mouseout', tip.hide)
    .attr('class', 'bar')
    .attr('fill', function(d, i) {
      return colors[i]
    })
    .attr('x', function(d) {
      return x(d.letter)
    })
    .attr('y', height) // 控制动画由下而上
    .attr('width', x.bandwidth())
    .attr('height', 0) // 控制动画由下而上
    .transition()
    .duration(200)
    .ease(d3.easeBounceInOut)
    .delay(function(d, i) {
      return i * 200
    })
    .attr('y', function(d) {
      return y(d.frequency)
    })
    .attr('height', function(d) {
      return height - y(d.frequency)
    })

  // 输出柱图上的数值
  g.append('g')
    .attr('class', 'bar--text')
    .selectAll('text')
    .data(data)
    .enter()
    .append('text')
    .attr('fill', '#fff')
    .attr('font-size', '14px')
    .attr('text-anchor', 'center')
    .attr('x', function(d, i) {
      return x(d.letter)
    })
    .attr('y', function(d) {
      return y(d.frequency) + 5
    })
    .attr('dx', barWidth / 2)
    .attr('dy', '1em')
    .text(function(d) {
      return (d.frequency * 100).toFixed(2) + '%'
    })
    .on('mouseover', tip.show)
    .on('mouseout', tip.hide)

  // 标题
  chart
    .append('g')
    .attr('class', 'cahrt-title')
    .append('text')
    .attr('fill', '#000')
    .attr('font-size', '16px')
    .attr('font-weight', '700')
    .attr('text-anchor', 'middle')
    .attr('x', options.containerWidth / 2)
    .attr('y', 20)
    .text(title)
}