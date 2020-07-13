<template>
  <div class="flow-page" style="background:#fff;position:fixed;left:0;top:0;bottom:0;right:0;">
  <!-- <div class="flow-page"> -->  
    <div class="flow-nav">
      <div class="flow-nav-title">
        {{ options.title }}
      </div>
      <div class="flow-nav-steps">
        <el-steps 
          :space="200" 
          :active="active" 
          finish-status="success">
          <el-step 
            title="基础设置"
            @click.native="active=0"></el-step>
          <el-step 
            title="表单设计"
            @click.native="active=1"></el-step>
          <el-step 
            title="流程设计"
            @click.native="active=2"></el-step>
        </el-steps>
      </div>
      <div class="flow-nav-opts">
        <el-button 
          type="info" 
          size="small"
          plain
          @click="preview">预 览</el-button>
        <el-button 
          type="primary" 
          size="small"
          @click="save">发 布</el-button>
      </div>
    </div>
    
    <section class="flow-design"
      v-if="active == 0">
      <WfForm 
        :title="options.title"/>
    </section> 

    <section class="flow-tpl" 
      v-if="active == 1" 
      @mouseup="moveEnd" 
      @mousemove="move">
      <div class="flow-tpl-main">
        <WfWidgets/>
        <WfCanvas :list="options.components"/>
        <WfSettings/>
      </div>
      <WfDragging/>
    </section> 

    <section class="flow-design" 
      v-if="active == 2">
      <!-- <div class="flow-zoom">
        <div class="zoom-out">
          <i class="el-icon-minus"></i>
        </div>
        <span>100%</span>
        <div class="zoom-in">
          <i class="el-icon-plus"></i>
        </div>
      </div> -->
      <div ref="design-scale" class="flow-design-scale">
        <StartNode />
        <Node
          v-for="(item, index) in items"
          :key="index"
          :node="item"
          @addnode="addnode"
          @delNode="delNode(item)"/>
        <EndNode />
      </div>
    </section>
    
    <!-- 预览弹窗 -->
    <el-dialog
      :append-to-body="true"
      class="xa-dialog"
      title="预览"
      :visible.sync="dialog.view">
      <pre>{{ JSON.stringify(options.node, null, 4) }}</pre>
    </el-dialog>

    <!-- 发布错误提示 -->
    <ErrorsModal
      :show.sync="dialog.error"
      :data="errors"/>
  </div>
</template>

<script> 
import Vue from 'vue'
window.drag = new Vue()

import Node from '@/components/Workflow/ui/node'
Vue.component('Node', Node) 

import StartNode from './start-node' 
import EndNode from './end-node'
import ErrorsModal from './errors-modal'
import WfForm from '../tpl/wf-form'
import WfWidgets from '../tpl/wf-widgets'
import WfCanvas from '../tpl/wf-canvas'
import WfSettings from '../tpl/wf-settings'
import WfDragging from '../tpl/wf-dragging'
import { iteratorData, addNewNode, delNode, checkData } from './process'
export default {
  name: 'work-flow',
  components: {
    StartNode,
    EndNode,
    ErrorsModal,
    WfForm,
    WfWidgets,
    WfCanvas,
    WfSettings,
    WfDragging
  },
  props: {
    data: {
      type: Object,
      default: undefined
    }
  },
  data: () => ({
    active: 1,
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
  watch: {
    data: {
      handler (val) {
        this.options = val
      },
      deep: true
    }
  },
  methods: {
    move(e) {
      if (this.isStart) {
        document.querySelector('html').classList.add('wf-cursor-move')
        let obj = {
          componentName: this.componentView.componentName,
          clientX: e.clientX,
          clientY: e.clientY
        }
        drag.$emit("moveInCanvas", obj)
        drag.$emit('move', e)
      }
    },
    moveEnd(e) {
      if (this.isStart) {
        let obj = {
            componentView: this.componentView
        }
        drag.$emit("moveEnd", obj)
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
      console.log('save-->', 'node=', this.options)

      var errors = checkData(this.options.node)
      if (errors.length > 0) {
        this.dialog.error = true
        this.errors = errors
        return
      }
      this.$emit('ok', this.options)
    },
    preview () {
      var errors = checkData(this.options.node)
      console.log(errors)
      if (errors.length > 0) {
        this.dialog.error = true
        this.errors = errors
        return
      }
      this.dialog.view = true
    }
  },
  created () {
    let self = this
    
    drag.$on('moveStart', function(obj){
      self.isStart = true
      self.componentView = obj.componentView
    })

    drag.$on('save', function(components){
      self.options.components = components
    })
  },
  mounted () {
    this.options = this.data
    this.iteratorData(this.options.node)
  }
}
</script>
<style lang="scss">
@import "../../../styles/_global.scss";

.flow {

  &-nav {
    @include flex-row();
    padding:0 20px;
    padding-bottom: 10px;
    align-items: center;
    font-size: 14px;

    &-title {
      flex:2;
      color:$color-primary;
      font-size: 16px;
    }

    &-steps {
      flex:4;
      cursor: pointer;
    }

    &-opts {
      text-align:right;
    }
  }

  &-tpl {
    width:100%;
    height:100%;
    background-color: #f5f5f7;

    &-main {
      position:relative;
      display:flex;
      flex-direction:row;
      algin-items:center;
      min-width:1200px;
      height:100%;
      min-height:670px;
      //padding:20px;
    }
  }

  &-design {
    width: 100%;
    height: 100%;
    padding-bottom:200px;
    overflow:auto;
    background-color: #f5f5f7;

    &-scale {
      transform: scale(1);
      display: inline-block;
      position: relative;
      width: 100%;
      height: 100%;
      padding: 54.5px 0;     
      align-items: flex-start;     
      justify-content: center;
      flex-wrap: wrap;
      min-width: min-content;
      transform: scale(1); 
      transform-origin: 50% 0px 0px;
    }
  }

  &-zoom {
    display: flex;
    position: fixed;
    align-items: center;
    justify-content: space-between;
    height: 40px;
    width: 125px;
    right: 40px;
    margin-top: 30px;
    z-index: 1;

    &-out,
    &-in {
      width: 30px;
      height: 30px;
      text-align:center;
      line-height:30px;
      background: #fff;
      color: #333;
      cursor: pointer;
    }
  }
}
</style>
