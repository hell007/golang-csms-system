<template>
  <div class="flow-page">  
    <ul class="flow-nav">
      <li v-bind:class="{'on':active==0}">
        <em>1</em>
        <span>基础设置</span>
      </li>
      <li v-bind:class="{'on':active==1}">
        <em>2</em>
        <span>表单设计</span>
      </li>
      <li v-bind:class="{'on':active==2}">
        <em>3</em>
        <span>流程设计</span>
      </li>
    </ul>
    
    <!-- 基础设置 -->
    <section class="flow-design"
      v-if="active == 0">
      <WfForm
        :form="options.form"/>
      <div class="flow-next">
        <el-button 
          type="primary"
          :disabled="isLocked"
          @click="active=1">下一步</el-button>
        <el-button>返&nbsp;&nbsp;&nbsp;&nbsp;回</el-button>  
      </div>  
    </section> 
    
    <!-- 表单设计 -->
    <section class="flow-tpl" 
      v-if="active == 1" 
      @mouseup="moveEnd" 
      @mousemove="move">
      <div class="flow-tpl-main">
        <WfWidgets/>
        <WfCanvas 
          :list="options.components"/>
        <WfSettings/>
      </div>
      <WfDragging/>
      <div class="flow-next">
        <el-button
          @click="active=0">上一步</el-button>
        <el-button 
          type="primary"
          @click="active=2">下一步</el-button>
      </div>
    </section> 
    
    <!-- 流程设计 -->
    <section class="flow-design" 
      v-if="active == 2">
      <div class="flow-design-scale">
        <StartNode />
        <Node
          v-for="(item, index) in items"
          :key="index"
          :node="item"
          @addnode="addnode"
          @delNode="delNode(item)"/>
        <EndNode />
        <div class="flow-next">
          <el-button
            @click="active=1">上一步</el-button>
          <el-button 
            type="primary"
            @click="save">发&nbsp;&nbsp;&nbsp;&nbsp;布</el-button>
        </div>
      </div>
    </section>
    
    <!-- 错误提示 -->
    <el-dialog
      :append-to-body="true"
      class="xa-dialog"
      title="错误提示"
      :visible.sync="dialog.view">
      <div>
        <pre>{{ JSON.stringify(options.node, null, 4) }}</pre>
      </div>
      <div slot="footer">
        <el-button 
          size="small"
          @click="dialog.view = false">我知道了</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script> 
import Vue from 'vue'
window.drag = new Vue()

import Node from '@/components/Workflow/ui/node'
Vue.component('Node', Node) 

import StartNode from '@/components/Workflow/ui/start-node' 
import EndNode from '@/components/Workflow/ui/end-node'
import WfForm from '@/components/Workflow/tpl/wf-form'
import WfWidgets from '@/components/Workflow/tpl/wf-widgets'
import WfCanvas from '@/components/Workflow/tpl/wf-canvas'
import WfSettings from '@/components/Workflow/tpl/wf-settings'
import WfDragging from '@/components/Workflow/tpl/wf-dragging'
import { iteratorData, addNewNode, delNode } from '@/components/Workflow/ui/process'

// 模拟数据
const data = {
  form: {
    title: '',
    date: '',
    person: '',
  },
  components: [{
    InTableCanvas: null,
    componentName: "textfield",
    defaultLable: "姓名",
    defaultProps: "请输入姓名",
    defaultRequired: true,
    idx: 0,
    name: "单行输入框",
    supports: ["label", "placeholder", "required"]
  }, {
    InTableCanvas: null,
    componentName: "tablefield",
    components: [{
      componentName: "textfield",
      defaultLable: "类型",
      defaultProps: "请输入类型",
      defaultRequired: true,
      idx: 2,
      name: "单行输入框",
      supports: ["label", "placeholder", "required"]
    }],
    defaultAction: "增加",
    defaultLable: "表格",
    idx: 1,
    name: "明细/表格",
    selected: "0",
    supports: ["label", "action"]
  }],
  node: {
    name: '发起方',
    type: 'start',
    content: '选择类型/人员/菜单/应用',
    nodeId: '1',
    properties: null,
    childNode: {
      name: '下一节点人',
      type: 'approver',
      content: '请选择',
      prevId: '1',
      nodeId: '2',
      properties: null,
      notifier: {
        name: '抄送人',
        type: "notifier",
        content: '请选择',
        prevId: '2',
        nodeId: '3',
        properties: null
      }
    }
  }
}

export default {
  name: 'flow-add',
  components: {
    StartNode,
    EndNode,
    WfForm,
    WfWidgets,
    WfCanvas,
    WfSettings,
    WfDragging
  },
  data: () => ({
    active: 2,
    isLocked: true,
    isStart: false,
    componentView: {},
    items: [],
    errors: [],
    dialog: {
      error: false,
      view: false
    },
    options: {}
  }),
  methods: {
    move(e) {
      if (this.isStart) {
        document.querySelector('html').classList.add('wf-cursor-move')
        let obj = {
          componentName: this.componentView.componentName,
          clientX: e.clientX,
          clientY: e.clientY
        }
        drag.$emit('moveInCanvas', obj)
        drag.$emit('move', e)
      }
    },
    moveEnd(e) {
      if (this.isStart) {
        let obj = {
            componentView: this.componentView
        }
        drag.$emit('moveEnd', obj)
        this.isStart = false
      }
    },
    iteratorData (data) {
      this.items = []
      iteratorData(this.items, data)
    },
    addnode (node) {
      addNewNode(node, this.options.node, this.items)
    },
    delNode (node) {
      delNode(node, this.options.node, this.items)
    },
    save () {
      console.log('save 基础设置->', this.options.form)
      console.log('save 表单设计->', this.options.components)
      console.log('save 流程设计->', this.items)
    }
  },
  created () {
    let self = this
    this.options = data

    drag.$on('setLocked', function(bool){
      self.isLocked = bool
    })
    
    drag.$on('moveStart', function(obj){
      self.isStart = true
      self.componentView = obj.componentView
    })

    drag.$on('save', function(components){
      self.options.components = components
    })
  },
  mounted () {
    // 生成流程设计数据
    this.iteratorData(this.options.node)
  }
}
</script>
<style lang="scss">
@import "../../styles/_global.scss";

.flow {

  &-nav {
    @include flex-row();
    padding:20px;
    align-items: center;
    justify-content: center;
    background-color:#fff;

    li {
      position:relative;
      @include flex-row();
      align-items: center;
      padding:8px 16px;
      margin-right:100px;
      border:1px solid $color-primary;
      border-radius:999px;
      font-size:14px;
      color:$color-primary;
      cursor:pointer;

      &::after {
        position:absolute;
        right:-99px;
        top:20px;
        display:inline-block;
        content:'';
        border-top:1px solid $color-primary;
        width:100px;
      }

      &::before {
        position:absolute;
        right:-101px;
        top:15px;
        display:inline-block;
        @include arrow(right,1em,1px);
        color:$color-primary;
      }

      &:last-child {

        &::before,
        &::after {
          display:none;
        }
      }

      &.on {
        background-color:$color-primary;
        color:#fff;

        em {
          border-color:#fff;
        }
      }

      em {
        font-style:normal;
        width:20px;
        height:20px;
        margin-right:6px;
        border:1px solid $color-primary;
        border-radius:999px;
        text-align:center;
        line-height:20px;
        font-size:12px;
      }
    }
  }

  &-tpl {
    width:100%;
    height:100%;

    &-main {
      position:relative;
      display:flex;
      flex-direction:row;
      algin-items:center;
      min-width:1200px;
      height:100%;
      min-height:670px;
    }
  }

  &-design {
    width: 100%;
    height: 100%;
    overflow:auto;

    &-scale {
      transform: scale(1);
      display: inline-block;
      position: relative;
      width: 100%;
      height: 100%;
      padding: 54px 0;     
      align-items: flex-start;     
      justify-content: center;
      flex-wrap: wrap;
      min-width: min-content;
      transform: scale(1); 
      transform-origin: 50% 0px 0px;
    }
  }

  &-next {
    text-align:center;
    margin-top:60px;
  }
}
</style>
