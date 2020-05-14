// https://d3js.org.cn/document/
import * as d3 from 'd3'
import d3Tip from 'd3-tip'

export function setDendrogram(opt) {
  const config = {
    container: null,
    containerWidth: 1200,
    containerHeight: 600,
    margin: {
      top: 80,
      right: 80,
      bottom: 80,
      left: 80
    },
    primary: '',
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

  const chart = d3
    .select(options.container)
    .attr('width', options.containerWidth)
    .attr('height', options.containerHeight)

  chart
    .append('defs')
    .append('clipPath') // 添加长方形方块，遮罩作用
    .attr('id', 'clip')
    .append('rect')
    .attr('height', height)
    .attr('width', 0) // 用遮罩实现线动画
    .transition()
    .duration(1000)
    .attr('width', width)

  let g = chart
    .append('g')
    .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')') // 设最外包层在总图上的相对位置

  let z = d3.scaleOrdinal(d3.schemeCategory10) // 通用线条的颜色

  let root = d3
    .tree() // 定义树
    .size([height, width])
    .separation(function(a, b) {
      return a.parent === b.parent ? 1 : 2
    })(d3.hierarchy(data))

  let linkLine = d3
    .linkHorizontal() // 连接线创建器
    .source(function(d) {
      return [d.source.y, d.source.x]
    })
    .target(function(d) {
      return [d.target.y, d.target.x]
    })

  g.selectAll('.link') // 创建每条连接线
    .data(root.links())
    .enter()
    .append('path')
    .attr('clip-path', 'url(#clip)')
    .attr('class', 'link')
    .style('stroke', '#666')
    .style('stroke-width', 1)
    .attr('fill', 'none')
    .attr('d', linkLine)

  let node = g
    .selectAll('.node') // 定位并创建到每个节点
    .data(root.descendants())
    .enter()
    .append('g')
    .attr('class', function(d) {
      return 'node' + (d.children ? ' node--internal' : ' node--leaf')
    })
    .attr('transform', function(d) {
      return 'translate(' + d.y + ',' + d.x + ')'
    })

  node
    .append('circle') // 画小圆
    .style('fill', function(d) {
      return z(d.depth)
    })
    .attr('r', 2.5)

  node
    .append('text') // 输出文字
    .attr('dy', 3)
    .attr('x', function(d) {
      return d.children ? -8 : 8
    })
    .style('fill', function(d) {
      return z(d.depth)
    })
    .style('text-anchor', function(d) {
      return d.children ? 'end' : 'start'
    })
    .text(function(d) {
      return d.data.name
    })

}