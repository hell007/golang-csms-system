// https://d3js.org.cn/document/
import * as d3 from 'd3'
import d3Tip from 'd3-tip'

export function setChord(opt) {
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
      '#98abc5', 
      '#7b6888', 
      '#a05d56', 
      '#ff8c00'
    ],
    primary: '',
    title: '',
    data: {}
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
  const names = data.names
  const matrix = data.matrix
  const outerRadius = Math.min(width, height) * 0.5 - 40
  const innerRadius = outerRadius - 30

  const chart = d3
    .select(options.container)
    .attr('width', options.containerWidth)
    .attr('height', options.containerHeight)

  let chord = d3
    .chord() // 弦
    .padAngle(0.1)
    .sortSubgroups(d3.descending)(matrix)

  let arc = d3
    .arc() // 用于画弧
    .innerRadius(innerRadius)
    .outerRadius(outerRadius)

  let ribbon = d3
    .ribbon() // 用于画中间色带
    .radius(innerRadius)

  let color = d3
    .scaleOrdinal() // 四种颜色
    .domain(d3.range(4))
    .range(colors)

  let tip = d3Tip() // 设置tip
    .attr('class', 'd3-tip chord-tip')
    .offset([-10, 0])
    .html(function(d) {
      if (d.source.index !== d.source.subindex) {
        return (
          '<strong>' +
          names[d.source.index] +
          '->' +
          names[d.source.subindex] +
          ":</strong><span style='color:#ffeb3b'> " +
          (d.source.value / 10000).toFixed(2) +
          '万' +
          ' </span>件<br><strong>' +
          names[d.target.index] +
          '->' +
          names[d.target.subindex] +
          ":</strong><span style='color:#ffeb3b'> " +
          (d.target.value / 10000).toFixed(2) +
          '万</span> 件'
        )
      } else {
        return (
          '<strong>' +
          names[d.source.index] +
          '->' +
          names[d.source.subindex] +
          ":</strong><span style='color:#ffeb3b'> " +
          (d.source.value / 10000).toFixed(2) +
          '万</span> 件'
        )
      }
    })

  chart.call(tip)

  let g = chart
    .append('g') // 构建弦
    .attr('class', 'chart-chord')
    .attr(
      'transform',
      'translate(' +
      options.containerWidth / 2 +
      ',' +
      (height + margin.top + margin.bottom) / 2 +
      ')'
    )
    .datum(chord)

  let group = g
    .append('g') // 画刻度组
    .attr('class', 'groups')
    .selectAll('g')
    .data(function(chords) {
      return chords.groups
    })
    .enter()
    .append('g')

  group
    .append('path') // 画弧
    .style('fill', function(d) {
      return color(d.index)
    })
    .style('stroke', function(d) {
      return d3.rgb(color(d.index)).darker()
    })
    .attr('d', arc)

  group
    .append('g') // 输出名字
    .attr('class', 'group-name')
    .attr('transform', function(d) {
      return (
        'rotate(' +
        (((d.endAngle - (d.endAngle - d.startAngle) / 2) * 180) / Math.PI -
          90) +
        ') translate(' +
        (outerRadius + 60) +
        ',0)'
      )
    })
    .append('text')
    .attr('fill', function(d) {
      return d3.rgb(color(d.index)).darker()
    })
    .attr('x', 8)
    .attr('dy', '.35em')
    .attr('transform', function(d) {
      return d.endAngle - (d.endAngle - d.startAngle) / 2 > Math.PI ? 'rotate(180) translate(-16)' : null
    })
    .style('text-anchor', function(d) {
      return d.endAngle - (d.endAngle - d.startAngle) / 2 > Math.PI ? 'end' : null
    })
    .style('font-size', '16px')
    .style('font-weight', '700')
    .text(function(d) {
      return names[d.index]
    })

  let groupTick = group
    .selectAll('.group-tick') // 画刻度
    .data(function(d) {
      return groupTicks(d, 1e3)
    })
    .enter()
    .append('g')
    .attr('class', 'group-tick')
    .attr('transform', function(d) {
      return (
        'rotate(' +
        ((d.angle * 180) / Math.PI - 90) +
        ') translate(' +
        outerRadius +
        ',0)'
      )
    })

  groupTick
    .append('line') // 画刻度线
    .attr('x2', 6)

  groupTick // 输出刻度上的文字
    .filter(function(d) {
      return d.value % 5e3 === 0
    })
    .append('text')
    .attr('x', 8)
    .attr('dy', '.35em')
    .attr('transform', function(d) {
      return d.angle > Math.PI ? 'rotate(180) translate(-16)' : null
    })
    .style('text-anchor', function(d) {
      return d.angle > Math.PI ? 'end' : null
    })
    .text(function(d) {
      return (d.value / 10000).toFixed(2) + '万'
    })

  g.append('g') // 输出彩带
    .attr('class', 'ribbons')
    .selectAll('path')
    .data(function(chords) {
      return chords
    })
    .enter()
    .append('path')
    .on('mouseover', tip.show)
    .on('mouseout', tip.hide)
    .attr('d', ribbon)
    .style('fill', '#fff')
    .style('stroke', '#fff')
    .transition()
    .duration(50)
    .delay(function(d, i) {
      return i * 50
    })
    .style('fill', function(d) {
      return color(d.target.index)
    })
    .style('cursor', 'pointer')
    .style('stroke', function(d) {
      return d3.rgb(color(d.target.index)).darker()
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
    .text('某快递公司昨天来往于北上广深(含同城快递)的快递件数(件)')

  // Returns an array of tick angles and values for a given group and step.
  function groupTicks(d, step) {
    var k = (d.endAngle - d.startAngle) / d.value
    return d3.range(0, d.value, step).map(function(value) {
      return {
        value: value,
        angle: value * k + d.startAngle
      }
    })
  }

}