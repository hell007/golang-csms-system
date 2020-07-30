<template>
  <div class="condition-node-box">
    <div class="auto-judge">
      <!-- <div class="sort-left">
        <i class="el-icon-arrow-left"></i>
      </div> -->
      <div class="title-wrapper">
        <span class="editable-title">{{ node.name }}</span>
        <span class="priority-title">优先级{{ node.sort }}</span>
        <!-- <span class="copy">
          <i class="el-icon-document-copy"></i>
        </span> -->
        <span class="close" 
          @click="delConditionNode">
          <i class="el-icon-close"></i>
        </span>
      </div>
      <!-- <div class="sort-right">
        <i class="el-icon-arrow-right"></i>
      </div> -->
      <div
        class="content"
        @click="setProperties">
        <div>{{ desc }}</div>
      </div>
    </div>
    <AddNodeBtn
      :node="node"
      @addnode="addnode"/>
    <AddNodeCondition
      :show.sync="show"
      :properties="node.properties"
      @on-success="setPropertiesOK"/>
  </div>
</template>
<script>
import AddNodeBtn from './add-node-btn'
import AddNodeCondition from './add-node-condition'
export default {
  name: 'condition-node-box',
  components: {
    AddNodeBtn,
    AddNodeCondition
  },
  props: {
    text: {
      type: String,
      default: '请设置条件'
    },
    node: {
      type: Object,
      default: undefined
    }
  },
  data: () => ({
    show: false,
    desc: ''
  }),
  mounted () {
    console.log('condition-node-box',this.node)
    this.desc = this.getText()
    if (!this.node.properties) {
      this.node.properties = {
        conditions: [[]]
      }
    }
  },
  methods: {
    addnode () {
      this.$emit('addnode')
    },
    delConditionNode () {
      this.$emit('delConditionNode')
    },
    setProperties () {
      this.show = true
    },
    setPropertiesOK (properties) {
      this.node.properties = properties
      this.$emit('addConditionFactor', this.node)
      // this.desc.set(this.getText())
      this.desc = this.getText()
    },
    getText () {
      var text = '请设置条件'
      if (!this.node.properties) {
        return text
      }
      text = ''
      this.node.properties.conditions[0].forEach(cond => {
        var temp = ''
        temp += cond.paramLabel
        switch (cond.key) {
          case 'lt':
            temp += '<' + cond.upperBound
            break
          case 'le':
            temp += '≤' + cond.upperBoundEqual
            break
          case 'eq':
            temp += '=' + cond.boundEqual
            break
          case 'gt':
            temp += '>' + cond.lowerBound
            break
          case 'ge':
            temp += '≥' + cond.lowerBoundEqual
            break
          case 'between':
            temp = ''
            if (cond.lowerBound && cond.lowerBound !== '') {
              temp = cond.lowerBound + '<'
            } else {
              temp = cond.lowerBoundEqual + '≤'
            }
            temp += cond.paramLabel
            if (cond.upperBound && cond.upperBound !== '') {
              temp += '<' + cond.upperBound
            } else {
              temp += '≤' + cond.upperBoundEqual
            }
            break
          default:
        }
        temp += ' 且 '
        text += temp
      })
      text = text.substring(0, text.length - 2)
      return text
    }
  }
}
</script>
<style lang="scss" scoped>
.auto-judge {
  position: relative;
  width: 220px;
  min-height: 72px;
  padding: 14px 19px;
  border-radius: 4px;
  border:1px solid #fff;
  text-align: left;
  cursor: pointer;
  background: #fff;
  font-size: 14px;
  color: #191f25;
  box-shadow:0 2px 5px rgba(0, 0, 0, 0.045);

  .editable-title {
    @include ellipsis();
    display: inline-block;
    max-width: 120px;
  }
  
  .priority-title {
    display: inline-block;
    float: right;
    margin-right: 10px;
    color: rgba(25,31,37,.56);
  }
  
  .close,
  .copy {
    width: 14px;
    height: 14px;
    display: none;
    position: absolute;
    right: -2px;
    top: -2px;
    font-size: 14px;
    text-align: center;
    line-height: 20px;
    z-index: 2;
    color:#999;
    cursor:pointer;
  }

  .copy {
    right: 20px;
  }

  .title-wrapper {
    position: relative;
    font-size: 12px;
    color: #15bc83;
    text-align: left;
    line-height: 16px; 
  }

  .content {
    font-size: 14px;
    color: #191f25;
    text-align: left;
    margin-top: 6px;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
  }

  .sort-left, 
  .sort-right {
    position: absolute;
    top: 0;
    bottom: 0;
    display: none;
    z-index: 1;

    &:hover {
      background: #efefef;
      border-color: rgba(25,31,37,.08);
    }
  }

  .sort-right {
    right: 0;
    border-left: 1px solid #f6f6f6;
  }
  
  &.active,
  &:hover {
    border-color:rgb(50, 150, 250);

    .close,
    .copy {
      display:inline-block;
    } 

    .priority-title {
      display:none;
    }

    .sort-left, 
    .sort-right { 
      display: flex;
      align-items: center;
    }

    .sort-left {
      left: 0;
      border-right: 1px solid #f6f6f6;
    }
  }
}
</style>