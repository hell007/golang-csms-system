<template>
  <div class="condition-group">
    <span class="condition-group-param">
      参数 {{ pos + 1 }}
    </span>
    <div class="condition-group-list">
      <el-row :gutter="20">
        <el-col :span="12">
          <el-input
            size="small"
            v-model="data1.paramKey"
            placeholder="请输入key,如: day"
            clearable>
          </el-input>
        </el-col>
        <el-col :span="12">
          <el-input
            size="small"
            v-model="data1.paramLabel"
            placeholder="请输入label,如: 请假天数"
            clearable>
          </el-input>
        </el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-select 
            v-model="n" 
            style="width:100%"
            size="small"
            clearable 
            placeholder="请选择">
            <el-option
              v-for="item in items"
              :key="item.key"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-col>
        <el-col :span="12">
          <el-input-number 
            v-if="data1.key !== 'between'"
            size="small" 
            v-model="inputVal"></el-input-number>
        </el-col>
      </el-row>
      <!-- <div class="condition-panel-range-set">
        <ConditionRange
          :name.sync="data1.key"
          :label.sync="data1.label"
          :items="items"
        />
        <ConditionInputNumber
          v-if="data1.key !== 'between'"
          :value.sync="inputVal"
        />
        <ConditionRangeBetween
          v-if="data1.key === 'between'"
          :data.sync="data1"
        />
      </div> -->
    </div>
    <span class="condition-group-opt"
      @click="del">
      <i class="el-icon-delete"/></i>
    </span>
  </div>
</template>
<script>
export default {
  name: 'condition-group',
  components: {
  },
  props: {
    data: {
      type: Object,
      default: undefined
    },
    pos: {
      type: Number,
      default: undefined
    }
  },
  data () {
    return {
      inputVal: undefined,
      select: [],
      data1: {
        paramKey: '',
        paramLabel: ''
      },
      items: [
        { key: 'lt', label: '小于', value: 'upperBound' },
        { key: 'le', label: '小于等于', value: 'upperBoundEqual' },
        { key: 'eq', label: '等于', value: 'boundEqual' },
        { key: 'gt', label: '大于', value: 'lowerBound' },
        { key: 'ge', label: '大于等于', value: 'lowerBoundEqual' },
        { key: 'between', label: '介于(两个数之间)', value: '' }
      ]
    }
  },
  computed: {
    key () {
      return this.data1.key
    }
  },
  watch: {
    data: {
      handler (val) {
        this.data1 = val
      },
      deep: true
    },
    data1: {
      handler (val) {
        this.$emit('update:data', val)
      },
      deep: true
    },
    key (val, oldval) { // 监测key,变化之后要重新赋值
      if (oldval === undefined) return
      this.data1.upperBound = ''
      this.data1.upperBoundEqual = ''
      this.data1.boundEqual = ''
      this.data1.lowerBound = ''
      this.data1.lowerBoundEqual = ''
      this.inputVal = undefined
    },
    inputVal (val) {
      if (!val) return
      var key = this.data1.key

    console.log('===', key, val)
      if (!key) {
        this.$alert('先设置操作符')
        return
      }
      switch (key) {
        case 'lt':
          this.data1.upperBound = val
          break
        case 'le':
          this.data1.upperBoundEqual = val
          break
        case 'eq':
          this.data1.boundEqual = val
          break
        case 'gt':
          this.data1.lowerBound = val
          break
        case 'ge':
          this.data1.lowerBoundEqual = val
          break
        case 'between':
          this.data1.boundEqual = val
          break
        default:
      }
    }
  },
  mounted () {
    this.data1 = this.data
    this.setValue(this.data1)
  },
  methods: {
    del () {
      this.$emit('del')
    },
    setValue (val) {
      if (!val) return
      if (val.key && val.key !== '') {
        switch (val.key) {
          case 'lt':
            this.inputVal = val.upperBound
            val.label = '小于'
            break
          case 'le':
            this.inputVal = val.upperBoundEqual
            val.label = '小于等于'
            break
          case 'eq':
            this.inputVal = val.boundEqual
            val.label = '等于'
            break
          case 'gt':
            this.inputVal = val.lowerBound
            val.label = '大于'
            break
          case 'ge':
            this.inputVal = val.lowerBoundEqual
            val.label = '大于等于'
            break
          case 'between':
            val.label = '介于(两个数之间)'
            break
          default:
        }
      } else {
        if (val.lowerBound && val.lowerBound !== '') {
          val.key = 'gt'
          this.inputVal = val.lowerBound
          val.label = '大于'
        }
        if (val.upperBound && val.upperBound !== '') {
          val.key = 'lt'
          this.inputVal = val.upperBound
          val.label = '小于'
        }
      }
    }
  }
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

.condition {

  &-group {
    @include flex-row();
    align-items:center;
    padding:10px 0;

    &-param {
      flex: 4;
      text-align:left;
    }

    &-list {
      flex: 18;
    }

    &-opt {
      flex:2;
      text-align:center;
    }
  }

}  
</style>