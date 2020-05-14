// https://d3js.org.cn/document/
import * as d3 from 'd3'

export function setVoronoi(opt) {
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

  let g = chart
    .append('g')
    .attr('class', 'chart-voronoi')
    .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')')
    .on('touchmove mousemove', moved) // 设最外包层在总图上的相对位置

  let sites = d3
    .range(100) // 生成所有的点的坐标
    .map(function(d) {
      return [Math.random() * width, Math.random() * height]
    })

  let voronoi = d3
    .voronoi() // 定义维诺图
    .extent([
      [-1, -1],
      [width + 1, height + 1]
    ])

  let z = d3.scaleOrdinal(d3.schemeCategory10) // 通用线条的颜色

  let polygon = g
    .append('g') // 画多边形
    .attr('class', 'polygons')
    .selectAll('path')
    .data(voronoi.polygons(sites))
    .enter()
    .append('path')
    .call(redrawPolygon)

  let link = g
    .append('g') // 画Delaunay三角网(三角形的每条边)
    .attr('class', 'links')
    .selectAll('line')
    .data(voronoi.links(sites))
    .enter()
    .append('line')
    .call(redrawLink)

  let site = g
    .append('g') // 画生成元(小圆点)
    .attr('class', 'sites')
    .selectAll('circle')
    .data(sites)
    .enter()
    .append('circle')
    .call(redrawSite)

  function moved() {
    // 下面是mouseover重画
    sites[0] = d3.mouse(this)
    redraw()
  }

  function redraw() {
    let diagram = voronoi(sites)
    polygon = polygon.data(diagram.polygons()).call(redrawPolygon)
    link = link.data(diagram.links())
    link.exit().remove()
    link = link
      .enter()
      .append('line')
      .merge(link)
      .call(redrawLink)
    site = site.data(sites).call(redrawSite)
  }

  function redrawPolygon(polygon) {
    polygon
      .attr('d', function(d) {
        return d ? 'M' + d.join('L') + 'Z' : null
      })
      .attr('fill', function(d, n) {
        return z(n)
      })
  }

  function redrawLink(link) {
    link
      .attr('x1', function(d) {
        return d.source[0]
      })
      .attr('y1', function(d) {
        return d.source[1]
      })
      .attr('x2', function(d) {
        return d.target[0]
      })
      .attr('y2', function(d) {
        return d.target[1]
      })
  }

  function redrawSite(site) {
    site
      .attr('r', function(d, i) {
        return i === 0 ? 4 : 2.5
      })
      .attr('cx', function(d) {
        return d[0]
      })
      .attr('cy', function(d) {
        return d[1]
      })
  }
}