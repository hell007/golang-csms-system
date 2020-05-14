<template>
  <el-card class="xa-block" shadow="hover">
    <div class="xa-box xa-content">
      <el-button @click="loading">loading</el-button>
      <hr>
      <el-button :plain="true" @click="message1">消息提示</el-button>
      <el-button :plain="true" @click="message2">消息提示(成功)</el-button>
      <el-button :plain="true" @click="message3">消息提示(警告)</el-button>
      <el-button :plain="true" @click="message4">消息提示(消息)</el-button>
      <el-button :plain="true" @click="message5">消息提示(错误)</el-button>
      <hr>
      <el-button :plain="true" @click="messageBox1">Message Box(alert)</el-button>
      <el-button :plain="true" @click="messageBox2">Message Box(confirm)</el-button>
      <el-button :plain="true" @click="messageBox3">Message Box(prompt)</el-button>
      <hr>
      <el-button :plain="true" @click="notify1">Notification 通知(错误)</el-button>
      <hr>
      <h4>dialog</h4>
      <el-button :plain="true" @click="dialogs.dialog = true">Dialog</el-button>
      <el-dialog title="提示" :visible.sync="dialogs.dialog" size="tiny">
        <span>这是一段信息</span>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.dialog = false">取 消</el-button>
          <el-button type="primary" @click="dialogs.dialog = false">确 定</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>dropdown</h4>
      <el-dropdown split-button :plain="true">
        更多菜单
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item>黄金糕</el-dropdown-item>
          <el-dropdown-item>狮子头</el-dropdown-item>
          <el-dropdown-item>螺蛳粉</el-dropdown-item>
          <el-dropdown-item>双皮奶</el-dropdown-item>
          <el-dropdown-item>蚵仔煎</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
      <el-button type="success" @click="dialogs.drop = true">Dialog中的下拉框</el-button>
      <el-dialog title="Dialog中的下拉框" :visible.sync="dialogs.drop">
        <el-dropdown split-button :plain="true">
          更多菜单
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>黄金糕</el-dropdown-item>
            <el-dropdown-item>狮子头</el-dropdown-item>
            <el-dropdown-item>螺蛳粉</el-dropdown-item>
            <el-dropdown-item>双皮奶</el-dropdown-item>
            <el-dropdown-item>蚵仔煎</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.drop = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>tooltip</h4>
      <el-tooltip class="item" effect="dark" content="Top Left 提示文字" placement="top-start">
        <el-button>上左</el-button>
      </el-tooltip>
      <el-button type="success" @click="dialogs.tooltip = true">dialog中的tooltip</el-button>
      <el-dialog title="Dialog中的tooltip" :visible.sync="dialogs.tooltip">
        <el-tooltip class="item" effect="dark" content="Top Left 提示文字" placement="top-start">
          <el-button>上左</el-button>
        </el-tooltip>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.tooltip = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>popover</h4>
      <el-popover ref="popover1" placement="top-start" title="标题" width="200" trigger="hover" content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。">
      </el-popover>
      <el-button v-popover:popover1>hover 激活</el-button>
      <el-button type="success" @click="dialogs.popover = true">Dialog中的popover(fixed)</el-button>
      <el-dialog title="Dialog中的下拉框" :visible.sync="dialogs.popover">
        <el-popover ref="popover2" placement="top-start" title="标题" width="200" trigger="hover" content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。">
        </el-popover>
        <el-button v-popover:popover2>hover 激活</el-button>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.popover = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>select</h4>
      <el-select v-model="value" placeholder="请选择">
        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
        </el-option>
      </el-select>
      <el-button type="success" @click="dialogs.select = true">Dialog中的select(正常)</el-button>
      <el-dialog title="Dialog中的下拉框" :visible.sync="dialogs.select">
        <el-select v-model="value" placeholder="请选择">
          <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.select = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>Cascader </h4>
      <el-cascader v-model="cascaderValue" :options="cascaderOptions">
      </el-cascader>
      <el-button type="success" @click="dialogs.cascader = true">Dialog中的cascader(fixed)</el-button>
      <el-dialog title="Dialog中的cascader" :visible.sync="dialogs.cascader">
        <el-cascader v-model="cascaderValue" :options="cascaderOptions">
        </el-cascader>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.cascader = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>Slider </h4>
      <el-slider v-model="slideValue" style="width: 200px;display:inline-block"></el-slider>
      <el-button type="success" @click="dialogs.slider = true">Dialog中的Slider (fixed)</el-button>
      <el-dialog title="Dialog中的Slider" :visible.sync="dialogs.slider">
        <el-slider v-model="slideValue"></el-slider>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.slider = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>TimeSelect </h4>
      <el-time-select v-model="tsValue" :picker-options="{
            start: '08:30',
            step: '00:15',
            end: '18:30'
          }" placeholder="选择时间">
      </el-time-select>
      <el-button type="success" @click="dialogs.timeSelect = true">Dialog中的TimeSelect (fixed)</el-button>
      <el-dialog title="Dialog中的 time select" :visible.sync="dialogs.timeSelect">
        <el-time-select v-model="tsValue" :picker-options="{
            start: '08:30',
            step: '00:15',
            end: '18:30'
          }" placeholder="选择时间">
        </el-time-select>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.timeSelect = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>timepicker</h4>
      <el-time-picker v-model="tpValue" placeholder="任意时间点">
      </el-time-picker>
      <el-button type="success" @click="dialogs.timePicker = true">Dialog中的TimePicker</el-button>
      <el-dialog title="Dialog中的timePicker" :visible.sync="dialogs.timePicker">
        <el-time-picker v-model="tpValue" placeholder="任意时间点">
        </el-time-picker>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.timePicker = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>timepicker range</h4>
      <el-time-picker is-range v-model="tprValue" placeholder="选择时间范围">
      </el-time-picker>
      <el-button type="success" @click="dialogs.timePicker2 = true">Dialog中的timepicker range</el-button>
      <el-dialog title="Dialog中的Slider" :visible.sync="dialogs.timePicker2">
        <el-time-picker is-range v-model="tprValue" placeholder="选择时间范围">
        </el-time-picker>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.timePicker2 = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>datetimepicke</h4>
      <el-date-picker v-model="dtpValue" type="datetime" placeholder="选择日期时间" align="right">
      </el-date-picker>
      <el-button type="success" @click="dialogs.dateTimePicker = true">Dialog中的dateTimePicker</el-button>
      <el-dialog title="Dialog中的Slider" :visible.sync="dialogs.dateTimePicker">
        <el-date-picker v-model="dtpValue" type="datetime" placeholder="选择日期时间" align="right">
        </el-date-picker>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.dateTimePicker = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>DatePicker </h4>
      <el-date-picker v-model="dpValue" type="date" placeholder="选择日期">
      </el-date-picker>
      <el-button type="success" @click="dialogs.datePicker  = true">Dialog中的DatePicker</el-button>
      <el-dialog title="DatePicker" :visible.sync="dialogs.datePicker">
        <el-date-picker v-model="dpValue" type="date" placeholder="选择日期">
        </el-date-picker>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.datePicker = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>ColorPicker </h4>
      <el-color-picker v-model="cpValue" show-alpha></el-color-picker>
      <el-button type="success" @click="dialogs.colorPicker = true">Dialog中的DateColorPicker (fixed)</el-button>
      <el-dialog title="ColorPicker " :visible.sync="dialogs.colorPicker">
        <el-color-picker v-model="cpValue" show-alpha></el-color-picker>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.colorPicker = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>autocomplete</h4>
      <el-autocomplete class="inline-input" :fetch-suggestions="querySearch" placeholder="请输入内容"></el-autocomplete>
      <el-button type="success" @click="dialogs.autocomplete = true">Dialog中的autocomplete (fixed)</el-button>
      <el-dialog title="autocomplete" :visible.sync="dialogs.autocomplete">
        <el-autocomplete class="inline-input" :fetch-suggestions="querySearch" placeholder="请输入内容"></el-autocomplete>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.autocomplete = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>复合型输入框</h4>
      <el-input placeholder="请输入内容" style="width: 300px;">
        <el-select v-model="value" slot="prepend" placeholder="请选择">
          <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
        <el-button slot="append" icon="el-icon-search"></el-button>
      </el-input>
      <el-button type="success" @click="dialogs.selectInInput = true">Dialog中的复合型输入框（fixed）</el-button>
      点击空白处不能关闭下拉
      <el-dialog title="autocomplete" :visible.sync="dialogs.selectInInput">
        <el-input placeholder="请输入内容">
          <el-select v-model="value" slot="prepend" placeholder="请选择">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
          <el-button slot="append" icon="el-icon-search"></el-button>
        </el-input>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.selectInInput = false">取 消</el-button>
        </span>
      </el-dialog>
      <hr>
      <h4>table</h4>
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column prop="date" label="日期" sortable>
        </el-table-column>
        <el-table-column prop="name" label="姓名">
        </el-table-column>
        <el-table-column prop="address" label="地址">
        </el-table-column>
        <el-table-column prop="tag" label="标签" :filters="[{ text: '家', value: '家' }, { text: '公司', value: '公司' }]" filter-placement="bottom-end">
          <template slot-scope="scope">
            <el-tag :type="scope.row.tag === '家' ? 'primary' : 'success'" close-transition>{{scope.row.tag}}</el-tag>
          </template>
        </el-table-column>
      </el-table>
      <el-button type="success" @click="dialogs.table = true">Dialog中的table</el-button>
      <el-dialog size="large" title="autocomplete" :visible.sync="dialogs.table">
        <el-table :data="tableData" border style="width: 100%">
          <el-table-column prop="date" label="日期" sortable width="180">
          </el-table-column>
          <el-table-column prop="name" label="姓名" width="180">
          </el-table-column>
          <el-table-column prop="address" label="地址">
          </el-table-column>
          <el-table-column prop="tag" label="标签" width="100" :filters="[{ text: '家', value: '家' }, { text: '公司', value: '公司' }]" filter-placement="bottom-end">
            <template slot-scope="scope">
              <el-tag :type="scope.row.tag === '家' ? 'primary' : 'success'" close-transition>{{scope.row.tag}}</el-tag>
            </template>
          </el-table-column>
        </el-table>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogs.table = false">取 消</el-button>
        </span>
      </el-dialog>
    </div>
  </el-card>
</template>
<script>
export default {
  name: 'unit-components',
  components: {
  },
  data: function() {
    return {
      cascaderValue: '',
      cascaderOptions: [{
          value: 'zhinan',
          label: '指南',
          children: [{
              value: 'shejiyuanze',
              label: '设计原则',
              children: [{
                  value: 'yizhi',
                  label: '一致'
                },
                {
                  value: 'fankui',
                  label: '反馈'
                },
                {
                  value: 'xiaolv',
                  label: '效率'
                },
                {
                  value: 'kekong',
                  label: '可控'
                }
              ]
            },
            {
              value: 'daohang',
              label: '导航',
              children: [{
                  value: 'cexiangdaohang',
                  label: '侧向导航'
                },
                {
                  value: 'dingbudaohang',
                  label: '顶部导航'
                }
              ]
            }
          ]
        },
        {
          value: 'zujian',
          label: '组件',
          children: [{
              value: 'basic',
              label: 'Basic',
              children: [{
                  value: 'layout',
                  label: 'Layout 布局'
                },
                {
                  value: 'color',
                  label: 'Color 色彩'
                },
                {
                  value: 'typography',
                  label: 'Typography 字体'
                },
                {
                  value: 'icon',
                  label: 'Icon 图标'
                },
                {
                  value: 'button',
                  label: 'Button 按钮'
                }
              ]
            },
            {
              value: 'form',
              label: 'Form',
              children: [{
                  value: 'radio',
                  label: 'Radio 单选框'
                },
                {
                  value: 'checkbox',
                  label: 'Checkbox 多选框'
                },
                {
                  value: 'input',
                  label: 'Input 输入框'
                },
                {
                  value: 'input-number',
                  label: 'InputNumber 计数器'
                },
                {
                  value: 'select',
                  label: 'Select 选择器'
                },
                {
                  value: 'cascader',
                  label: 'Cascader 级联选择器'
                },
                {
                  value: 'switch',
                  label: 'Switch 开关'
                },
                {
                  value: 'slider',
                  label: 'Slider 滑块'
                },
                {
                  value: 'time-picker',
                  label: 'TimePicker 时间选择器'
                },
                {
                  value: 'date-picker',
                  label: 'DatePicker 日期选择器'
                },
                {
                  value: 'datetime-picker',
                  label: 'DateTimePicker 日期时间选择器'
                },
                {
                  value: 'upload',
                  label: 'Upload 上传'
                },
                {
                  value: 'rate',
                  label: 'Rate 评分'
                },
                {
                  value: 'form',
                  label: 'Form 表单'
                }
              ]
            },
            {
              value: 'data',
              label: 'Data',
              children: [{
                  value: 'table',
                  label: 'Table 表格'
                },
                {
                  value: 'tag',
                  label: 'Tag 标签'
                },
                {
                  value: 'progress',
                  label: 'Progress 进度条'
                },
                {
                  value: 'tree',
                  label: 'Tree 树形控件'
                },
                {
                  value: 'pagination',
                  label: 'Pagination 分页'
                },
                {
                  value: 'badge',
                  label: 'Badge 标记'
                }
              ]
            },
            {
              value: 'notice',
              label: 'Notice',
              children: [{
                  value: 'alert',
                  label: 'Alert 警告'
                },
                {
                  value: 'loading',
                  label: 'Loading 加载'
                },
                {
                  value: 'message',
                  label: 'Message 消息提示'
                },
                {
                  value: 'message-box',
                  label: 'MessageBox 弹框'
                },
                {
                  value: 'notification',
                  label: 'Notification 通知'
                }
              ]
            },
            {
              value: 'navigation',
              label: 'Navigation',
              children: [{
                  value: 'menu',
                  label: 'NavMenu 导航菜单'
                },
                {
                  value: 'tabs',
                  label: 'Tabs 标签页'
                },
                {
                  value: 'breadcrumb',
                  label: 'Breadcrumb 面包屑'
                },
                {
                  value: 'dropdown',
                  label: 'Dropdown 下拉菜单'
                },
                {
                  value: 'steps',
                  label: 'Steps 步骤条'
                }
              ]
            },
            {
              value: 'others',
              label: 'Others',
              children: [{
                  value: 'dialog',
                  label: 'Dialog 对话框'
                },
                {
                  value: 'tooltip',
                  label: 'Tooltip 文字提示'
                },
                {
                  value: 'popover',
                  label: 'Popover 弹出框'
                },
                {
                  value: 'card',
                  label: 'Card 卡片'
                },
                {
                  value: 'carousel',
                  label: 'Carousel 走马灯'
                },
                {
                  value: 'collapse',
                  label: 'Collapse 折叠面板'
                }
              ]
            }
          ]
        },
        {
          value: 'ziyuan',
          label: '资源',
          children: [{
              value: 'axure',
              label: 'Axure Components'
            },
            {
              value: 'sketch',
              label: 'Sketch Templates'
            },
            {
              value: 'jiaohu',
              label: '组件交互文档'
            }
          ]
        }
      ],
      value: '选项1',
      options: [{
          value: '选项1',
          label: '黄金糕'
        },
        {
          value: '选项2',
          label: '双皮奶'
        },
        {
          value: '选项3',
          label: '蚵仔煎'
        },
        {
          value: '选项4',
          label: '龙须面'
        },
        {
          value: '选项5',
          label: '北京烤鸭'
        }
      ],
      slideValue: 12,
      tsValue: '',
      tpValue: '',
      dpValue: '',
      cpValue: '',
      tprValue: '',
      dtpValue: '',
      dialogs: {
        drop: false,
        dialog: false,
        tooltip: false,
        popover: false,
        select: false,
        cascader: false,
        slider: false,
        timeSelect: false,
        timePicker: false,
        timePicker2: false,
        datePicker: false,
        dateTimePicker: false,
        colorPicker: false,
        autocomplete: false,
        selectInInput: false,
        table: false
      },
      tableData: [{
          date: '2016-05-02',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄',
          tag: '家'
        },
        {
          date: '2016-05-04',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1517 弄',
          tag: '公司'
        },
        {
          date: '2016-05-01',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1519 弄',
          tag: '家'
        },
        {
          date: '2016-05-03',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1516 弄',
          tag: '公司'
        }
      ],
      autocomplete: [{
          value: '三全鲜食（北新泾店）',
          address: '长宁区新渔路144号'
        },
        {
          value: 'Hot honey 首尔炸鸡（仙霞路）',
          address: '上海市长宁区淞虹路661号'
        },
        {
          value: '新旺角茶餐厅',
          address: '上海市普陀区真北路988号创邑金沙谷6号楼113'
        },
        {
          value: '泷千家(天山西路店)',
          address: '天山西路438号'
        },
        {
          value: '胖仙女纸杯蛋糕（上海凌空店）',
          address: '上海市长宁区金钟路968号1幢18号楼一层商铺18-101'
        },
        {
          value: '贡茶',
          address: '上海市长宁区金钟路633号'
        },
        {
          value: '豪大大香鸡排超级奶爸',
          address: '上海市嘉定区曹安公路曹安路1685号'
        },
        {
          value: '茶芝兰（奶茶，手抓饼）',
          address: '上海市普陀区同普路1435号'
        },
        {
          value: '十二泷町',
          address: '上海市北翟路1444弄81号B幢-107'
        },
        {
          value: '星移浓缩咖啡',
          address: '上海市嘉定区新郁路817号'
        },
        {
          value: '阿姨奶茶/豪大大',
          address: '嘉定区曹安路1611号'
        },
        {
          value: '新麦甜四季甜品炸鸡',
          address: '嘉定区曹安公路2383弄55号'
        },
        {
          value: 'Monica摩托主题咖啡店',
          address: '嘉定区江桥镇曹安公路2409号1F，2383弄62号1F'
        },
        {
          value: '浮生若茶（凌空soho店）',
          address: '上海长宁区金钟路968号9号楼地下一层'
        },
        {
          value: 'NONO JUICE  鲜榨果汁',
          address: '上海市长宁区天山西路119号'
        },
        {
          value: 'CoCo都可(北新泾店）',
          address: '上海市长宁区仙霞西路'
        },
        {
          value: '快乐柠檬（神州智慧店）',
          address: '上海市长宁区天山西路567号1层R117号店铺'
        },
        {
          value: 'Merci Paul cafe',
          address: '上海市普陀区光复西路丹巴路28弄6号楼819'
        },
        {
          value: '猫山王（西郊百联店）',
          address: '上海市长宁区仙霞西路88号第一层G05-F01-1-306'
        },
        {
          value: '枪会山',
          address: '上海市普陀区棕榈路'
        },
        {
          value: '纵食',
          address: '元丰天山花园(东门) 双流路267号'
        },
        {
          value: '钱记',
          address: '上海市长宁区天山西路'
        },
        {
          value: '壹杯加',
          address: '上海市长宁区通协路'
        },
        {
          value: '唦哇嘀咖',
          address: '上海市长宁区新泾镇金钟路999号2幢（B幢）第01层第1-02A单元'
        },
        {
          value: '爱茜茜里(西郊百联)',
          address: '长宁区仙霞西路88号1305室'
        },
        {
          value: '爱茜茜里(近铁广场)',
          address: '上海市普陀区真北路818号近铁城市广场北区地下二楼N-B2-O2-C商铺'
        },
        {
          value: '鲜果榨汁（金沙江路和美广店）',
          address: '普陀区金沙江路2239号金沙和美广场B1-10-6'
        },
        {
          value: '开心丽果（缤谷店）',
          address: '上海市长宁区威宁路天山路341号'
        },
        {
          value: '超级鸡车（丰庄路店）',
          address: '上海市嘉定区丰庄路240号'
        },
        {
          value: '妙生活果园（北新泾店）',
          address: '长宁区新渔路144号'
        },
        {
          value: '香宜度麻辣香锅',
          address: '长宁区淞虹路148号'
        },
        {
          value: '凡仔汉堡（老真北路店）',
          address: '上海市普陀区老真北路160号'
        },
        {
          value: '港式小铺',
          address: '上海市长宁区金钟路968号15楼15-105室'
        },
        {
          value: '蜀香源麻辣香锅（剑河路店）',
          address: '剑河路443-1'
        },
        {
          value: '北京饺子馆',
          address: '长宁区北新泾街道天山西路490-1号'
        },
        {
          value: '饭典*新简餐（凌空SOHO店）',
          address: '上海市长宁区金钟路968号9号楼地下一层9-83室'
        },
        {
          value: '焦耳·川式快餐（金钟路店）',
          address: '上海市金钟路633号地下一层甲部'
        },
        {
          value: '动力鸡车',
          address: '长宁区仙霞西路299弄3号101B'
        },
        {
          value: '浏阳蒸菜',
          address: '天山西路430号'
        },
        {
          value: '四海游龙（天山西路店）',
          address: '上海市长宁区天山西路'
        },
        {
          value: '樱花食堂（凌空店）',
          address: '上海市长宁区金钟路968号15楼15-105室'
        },
        {
          value: '壹分米客家传统调制米粉(天山店)',
          address: '天山西路428号'
        },
        {
          value: '福荣祥烧腊（平溪路店）',
          address: '上海市长宁区协和路福泉路255弄57-73号'
        },
        {
          value: '速记黄焖鸡米饭',
          address: '上海市长宁区北新泾街道金钟路180号1层01号摊位'
        },
        {
          value: '红辣椒麻辣烫',
          address: '上海市长宁区天山西路492号'
        },
        {
          value: '(小杨生煎)西郊百联餐厅',
          address: '长宁区仙霞西路88号百联2楼'
        },
        {
          value: '阳阳麻辣烫',
          address: '天山西路389号'
        },
        {
          value: '南拳妈妈龙虾盖浇饭',
          address: '普陀区金沙江路1699号鑫乐惠美食广场A13'
        }
      ]
    }
  },
  methods: {
    loading: function() {
      var loading = this.$loading({
        target: parent.document.body
      })

      xa.utils.delay(2000).then(function() {
        loading.close()
      })
    },
    message1: function() {
      this.$message('这是一条消息提示')
    },
    message2: function() {
      this.$message.success('恭喜你，这是一条成功消息')
    },
    message3: function() {
      this.$message.warning('警告哦，这是一条警告消息')
    },
    message4: function() {
      this.$message('这是一条消息提示')
    },
    message5: function() {
      this.$message.error('错了哦，这是一条错误消息')
    },
    messageBox1: function() {
      this.$alert('这是一段内容', '标题名称', {
        confirmButtonText: '确定',
        callback: function(action) {
          console.log(action)
        }
      })
    },
    messageBox2: function() {
      this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        .then(function(action) {
          console.log(action)
        })
        .catch(function(action) {
          console.log(action)
        })
    },
    messageBox3: function() {
      this.$prompt('请输入邮箱', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          inputPattern: /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
          inputErrorMessage: '邮箱格式不正确'
        })
        .then(function(value) {
          console.log('输入的值', value)
        })
        .catch(function(action) {
          console.log(action)
        })
    },
    notify1: function() {
      this.$notify({
        title: '提示',
        message: '这是一条不会自动关闭的消息',
        type: 'success',
        duration: 6000
      })
    },
    querySearch: function(queryString, callback) {
      callback(
        queryString ?
        this.autocomplete.filter(function(item) {
          return item.value.indexOf(queryString) > -1
        }) :
        this.autocomplete
      )
    }
  }
}
</script>
<style rel="stylesheet/scss" lang="scss">
.el-button {
  margin-left:10px;
}
</style>



