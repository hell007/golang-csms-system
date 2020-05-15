// https://d3js.org.cn/document/
import * as d3 from 'd3'

function getChinaData() {
  return (async() => {
    try {
      const res = await fetch('http://cdn.a4z.cn/json/china.geojson')
      let data = await res.json()
      
    } catch (err) {
      console.log(err)
    }
  })()
}

export function setMap(opt) {
  const config = {
    container: null,
    containerWidth: 1200,
    containerHeight: 1000,
    margin: {
      top: 80,
      right: 20,
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
  if (!options.container) return

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

  let fill = d3.scaleOrdinal(colors) // 定义颜色

  let projection = d3
    .geoMercator() // 定义墨卡托地理投射(平面投射)
    .center([107, 31])
    .scale(d3.min([width, height]))
    .translate([width / 2, height / 2])

  let path = d3
    .geoPath() // 定义路径
    .projection(projection)

  let z = d3.scaleOrdinal(colors) // 通用线条的颜色

  let g = chart
    .append('g')
    .attr('transform', 'translate(' + margin.left + ',' + margin.top + ')') // 设最外包层在总图上的相对位置
    .style('fill-opacity', 0)

  let province = g
    .selectAll('path') // 绘画所有的省
    .data(data.features)
    .enter()
    .append('path')
    .attr('stroke', '#fff')
    .attr('stroke-width', 1)
    .attr('fill', function(d, i) {
      return z(i)
    })
    .attr('d', path)
    .on('mouseover', function(d, i) {
      d3.select(this).attr('fill', 'yellow')
    })
    .on('mouseout', function(d, i) {
      d3.select(this).attr('fill', z(i))
    })

  province
    .append('title') // 输出Title，mouseover显示
    .text(function(d) {
      return d.properties.name
    })

  g.transition()
    .duration(1000)
    .style('fill-opacity', 1) // 动画渐现

  chart
    .append('g') // 输出标题
    .attr('class', 'bar--title')
    .append('text')
    .attr('fill', '#000')
    .attr('font-size', '16px')
    .attr('font-weight', '700')
    .attr('text-anchor', 'middle')
    .attr('x', options.containerWidth / 2)
    .attr('y', 20)
    .text(title)
}