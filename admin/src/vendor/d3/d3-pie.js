// https://d3js.org.cn/document/
import * as d3 from 'd3'

export function setPie(opt) {
  const config = {
    container: null,
    containerWidth: 1200,
    containerHeight: 600,
    margin: {
      top: 80,
      right: 20,
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
  const title = options.title

  const chart = d3
    .select(options.container)
    .attr('width', options.containerWidth)
    .attr('height', options.containerHeight)

  const radius = Math.min(width, height) / 2

  // 设最外包层在总图上的相对位置
  let g = chart
    .append('g')
    .attr(
      'transform',
      'translate(' +
      (width / 2 + margin.left) +
      ',' +
      (margin.top + radius) +
      ')'
    )

  let colors = d3
    .scaleOrdinal()
    .range(options.colors)

  // 定义单个圆弧
  let arc = d3
    .arc()
    .innerRadius(0)
    .padAngle(0)

  // 定义单个圆弧里面的线条开始点所在的弧
  let startPointArc = d3
    .arc()
    .outerRadius(radius - 10)
    .innerRadius(radius - 10)

  // 定义单个圆弧里面的percent文字
  let percentLabelArc = d3
    .arc()
    .outerRadius(radius - 60)
    .innerRadius(radius - 60)

  // 定义单个圆弧里面的population文字
  let populationLabelArc = d3
    .arc()
    .outerRadius(radius + 40)
    .innerRadius(radius + 40)

  // 定义饼图
  let pie = d3
    .pie()
    .sort(null)
    .value(function(d) {
      return d.population
    })

  const sumData = d3.sum(data, function(d) {
    return d.population
  })

  // 定义颜色值域 
  colors.domain(
    d3
    .map(data, function(d) {
      return d.age
    })
    .keys()
  )

  // 画环图
  g.selectAll('.arc')
    .data(pie(data))
    .enter()
    .append('path')
    .attr('cursor', 'pointer')
    .attr('class', 'arc')
    .attr('stroke', function(d) {
      return colors(d.data.age)
    })
    .style('fill', function(d) {
      return colors(d.data.age)
    })
    .each(function(d) {
      // 储存当前起始与终点的角度、并设为相等
      let temp = {...d,
        endAngle: d.startAngle
      }
      d.outerRadius = radius - 10
      this._currentData = temp
    })
    .on('mouseover', arcTween(radius + 20, 0))
    .on('mouseout', arcTween(radius - 10, 150))
    .transition()
    .duration(100)
    .delay(function(d, i) {
      return i * 100
    })
    .attrTween('d', function(next) {
      // 动态设置d属性、生成动画
      var i = d3.interpolate(this._currentData, next)
      this._currentData = i(0) // 重设当前角度
      return function(t) {
        return arc(i(t))
      }
    })

  // 构造圆弧
  const arcs = pie(data)

  // 创建每条连接线
  let linkLine = g.append('g')
  let links = []
  arcs.forEach(function(d) {
    // 输出age文字
    const startPoint = startPointArc.centroid(d)
    const endPoint = populationLabelArc.centroid(d)
    links.push({
      source: [startPoint[0], startPoint[1]],
      target: [endPoint[0], endPoint[1]]
    })
  })

  linkLine
    .selectAll()
    .data(links)
    .enter()
    .append('path')
    .attr('class', 'link-line')
    .style('stroke', '#999')
    .style('stroke-width', 1)
    .attr('fill', 'none')
    .attr(
      'd',
      d3
      .linkHorizontal()
      .source(function(d) {
        return d.source
      })
      .target(function(d) {
        return d.target
      })
    )

  let label = g.append('g')

  arcs.forEach(function(d) {
    // 输出percent文字
    const c = percentLabelArc.centroid(d)
    label
      .append('text')
      .attr('class', 'age-text')
      .attr('fill', '#fff')
      .attr('font-weight', '700')
      .attr('font-size', '14px')
      .attr('text-anchor', 'middle')
      .attr('x', c[0])
      .attr('y', c[1])
      .text(((d.data.population * 100) / sumData).toFixed(1) + '%')
  })

  arcs.forEach(function(d) {
    // 输出population文字
    var c = populationLabelArc.centroid(d)
    label
      .append('text')
      .attr('class', 'age-text')
      .attr('fill', '#000')
      .attr('font-size', '12px')
      .attr('text-anchor', function(d) {
        return c[0] >= 0 ? 'start' : 'end'
      })
      .attr('x', c[0])
      .attr('y', c[1])
      .text(
        d.data.age + '岁 : ' + (d.data.population / 10000).toFixed(2) + '万人'
      )
  })

  // 标题
  chart
    .append('g')
    .attr('class', 'cahrt-title')
    .append('text')
    .attr('fill', '#fff')
    .attr('font-size', '14px')
    .attr('font-weight', '700')
    .attr(
      'transform',
      'translate(' +
      (width / 2 + margin.left) +
      ',' +
      (margin.top + radius) +
      ')'
    )
    .attr('text-anchor', 'middle')
    .attr('x', 0)
    .attr('y', 0)
    .text(title)

  function arcTween(outerRadius, delay) {
    // 设置缓动函数
    return function() {
      d3.select(this)
        .transition()
        .delay(delay)
        .attrTween('d', function(d) {
          var i = d3.interpolate(d.outerRadius, outerRadius)
          return function(t) {
            d.outerRadius = i(t)
            return arc(d)
          }
        })
    }
  }
}