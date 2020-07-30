<template>
<section ref="sponsor" class="modal"
  v-bind:class="{'modal-show':visible}">
  <section class="modal-right">
    <header class="modal-header">{{ title }}</header>
    <section class="modal-body">
      <div class="modal-panel">
        <div class="modal-panel-item">
          <h5 class="modal-panel-hd">流程类型</h5>
          <div class="modal-panel-bd">
            <el-radio-group
              class="el-radio-row"
              v-model="form.flow">
              <el-radio 
                v-for="(a, index) in flows"
                :key="a.label"
                :label="a.value">{{ a.label }}</el-radio>
            </el-radio-group>
          </div>
        </div>
        <!-- 常规流程 -->
        <div class="modal-panel-item"
          v-if="form.flow=='1'">
          <h5 class="modal-panel-hd">设置发起人</h5>
          <div class="modal-panel-bd">
            <el-radio-group
              class="el-radio-row"
              v-model="form.approver">
              <el-radio 
                v-for="(a, index) in approvers"
                :key="a.label"
                :label="a.value">{{ a.label }}</el-radio>
            </el-radio-group>
            <dl v-if="form.approver=='1'">
              <dt>
                <el-button
                  v-if="form.persons && form.persons.length>0"
                  type="primary" 
                  size="mini"
                  @click="personnels.show = true">修改成员</el-button>
                <el-button
                  v-else
                  type="primary"
                  icon="el-icon-plus"
                  size="mini"
                  @click="personnels.show = true">添加成员</el-button>  
              </dt>
              <dd>
                <div class="modal-panel-list">
                  <el-tag
                    v-for="tag in form.persons"
                    :key="tag.name"
                    size="medium"
                    type="info"
                    color="#f7f7f7"
                    effect="plain"
                    closable>
                    {{tag.name}}
                  </el-tag>
                  <span class="del"
                    v-if="form.persons && form.persons.length>0">清空</span>
                </div>
              </dd>
            </dl>
            <dl v-if="form.approver=='2'">
              <dt>
                <el-button
                  v-if="form.roles && form.roles.length>0"
                  type="primary" 
                  size="mini"
                  @click="personnels.show = true">修改角色</el-button>
                <el-button
                  v-else
                  type="primary"
                  icon="el-icon-plus"
                  size="mini"
                  @click="personnels.show = true">添加角色</el-button>  
              </dt>
              <dd>
                <div class="modal-panel-list">
                  <el-tag
                    v-for="tag in form.roles"
                    :key="tag.name"
                    size="medium"
                    type="info"
                    color="#f7f7f7"
                    effect="plain"
                    closable>
                    {{tag.name}}
                  </el-tag>
                  <span class="del"
                    v-if="form.roles && form.roles.length>0">清空</span>
                </div>
              </dd>
            </dl>
          </div>
        </div>
        <div class="modal-panel-item"
          v-if="form.flow=='1'">
          <h5 class="modal-panel-hd">设置工作台菜单</h5>
          <div class="modal-panel-bd">
            <dl>
              <dt>
                <el-button
                  v-if="form.menus && form.menus.length>0"
                  type="primary" 
                  size="mini"
                  @click="dialog.menu = true">修改菜单</el-button>
                <el-button
                  v-else
                  type="primary"
                  size="mini"
                  @click="dialog.menu = true">设置菜单</el-button>  
              </dt>
              <dd>
                <div class="modal-panel-list">
                  <el-tag
                    v-if="form.menus && form.menus.length>0"
                    size="medium"
                    type="info"
                    color="#f7f7f7"
                    effect="plain"
                    closable>
                    <span 
                      v-for="tag,index in form.menus"
                      :key="tag">
                      <em v-if="index!=0">/</em> 
                      {{tag}}
                    </span>
                  </el-tag>
                  <span class="del"
                    v-if="form.menus && form.menus.length>0">清空</span>
                </div>
              </dd>
            </dl>
          </div>
        </div>
        <div class="modal-panel-item"
          v-if="form.flow=='1'">
          <h5 class="modal-panel-hd">设置流程图标</h5>
          <div class="modal-panel-bd">
            <dl>
              <dt>
                <el-button
                  type="primary" 
                  size="mini"
                  @click="dialog.icon = true">选择图标</el-button>
              </dt>
              <dd class="modal-panel-picture">
                <div class="icon">
                  <img src="../../../assets/images/workflow/icon.png"/>
                </div>
              </dd>
            </dl>
          </div>
        </div>
        <!-- 应用流程 -->
        <div class="modal-panel-item"
          v-if="form.flow=='2'">
          <h5 class="modal-panel-hd">设置发起流程</h5>
          <div class="modal-panel-bd">
            <dl>
              <dt>
                <el-button
                  v-if="form.applys && form.applys.length>0"
                  type="primary" 
                  size="mini"
                  @click="dialog.apply = true">修改应用</el-button>
                <el-button
                  v-else
                  type="primary"
                  size="mini"
                  @click="dialog.apply = true">选择应用</el-button>  
              </dt>
              <dd>
                <div class="modal-panel-list">
                  <el-tag
                    v-if="form.applys && form.applys.length>0"
                    size="medium"
                    type="info"
                    color="#f7f7f7"
                    effect="plain"
                    closable>
                    <span 
                      v-for="tag,index in form.applys"
                      :key="tag">
                      <em v-if="index!=0">/</em> 
                      {{tag}}
                    </span>
                  </el-tag>
                  <span class="del"
                    v-if="form.applys && form.applys.length>0">清空</span>
                </div>
              </dd>
            </dl>
          </div>
        </div>
        <!-- 数据驱动流程 -->
        <div class="modal-panel-item"
          v-if="form.flow=='3'">
          <h5 class="modal-panel-hd">设置流程角色</h5>
          <div class="modal-panel-bd">
            <dl>
              <dt>
                <el-button
                  v-if="form.roles && form.roles.length>0"
                  type="primary" 
                  size="mini"
                  @click="personnels.show = true">修改角色</el-button>
                <el-button
                  v-else
                  type="primary"
                  icon="el-icon-plus"
                  size="mini"
                  @click="personnels.show = true">添加角色</el-button>  
              </dt>
              <dd>
                <div class="modal-panel-list">
                  <el-tag
                    v-for="tag in form.roles"
                    :key="tag.name"
                    size="medium"
                    type="info"
                    color="#f7f7f7"
                    effect="plain"
                    closable>
                    {{tag.name}}
                  </el-tag>
                  <span class="del"
                    v-if="form.roles && form.roles.length>0">清空</span>
                </div>
              </dd>
            </dl>
          </div>
        </div>
        <div class="modal-panel-item"
          v-if="form.flow=='3'">
          <h5 class="modal-panel-hd">设置流程数据指标</h5>
          <div class="modal-panel-bd">
            <dl>
              <dt>
                <el-button
                  v-if="form.targets && form.targets.length>0"
                  type="primary" 
                  size="mini"
                  @click="dialog.target = true">修改指标</el-button>
                <el-button
                  v-else
                  type="primary"
                  icon="el-icon-plus"
                  size="mini"
                  @click="dialog.target = true">设置指标</el-button>  
              </dt>
              <dd>
                <div class="modal-panel-list">
                  <el-tag
                    v-if="form.targets && form.targets.length>0"
                    size="medium"
                    type="info"
                    color="#f7f7f7"
                    effect="plain"
                    closable>
                    <span 
                      v-for="tag,index in form.targets"
                      :key="tag">
                      <em v-if="index!=0">/</em> 
                      {{tag}}
                    </span>
                  </el-tag>
                </div>
              </dd>
            </dl>
          </div>
        </div>
        <div class="modal-panel-item"
          v-if="form.flow=='3'">
          <h5 class="modal-panel-hd">设置指标数据类型</h5>
          <div class="modal-panel-bd">
            <dl>
              <dt>
                <el-button
                  v-if="form.targetType"
                  type="primary" 
                  size="mini"
                  @click="dialog.targetType = true">修改类型</el-button>
                <el-button
                  v-else
                  type="primary"
                  icon="el-icon-plus"
                  size="mini"
                  @click="dialog.targetType = true">设置类型</el-button>  
              </dt>
              <dd>
                <div class="modal-panel-list">
                  <el-tag
                    v-if="form.targetType"
                    size="medium"
                    type="info"
                    color="#f7f7f7"
                    effect="plain"
                    closable>
                    {{form.targetType}}
                  </el-tag>
                </div>
              </dd>
            </dl>
          </div>
        </div>
        <div class="modal-panel-item"
          v-if="form.flow=='3'">
          <h5 class="modal-panel-hd">设置驱动表达式</h5>
          <div class="modal-panel-bd">
            <dl>
              <dt>
                <el-button
                  v-if="form.expression"
                  type="primary" 
                  size="mini"
                  @click="dialog.expression = true">修改表达式</el-button>
                <el-button
                  v-else
                  type="primary"
                  icon="el-icon-plus"
                  size="mini"
                  @click="dialog.expression = true">设置表达式</el-button>
              </dt>
              <dd>
                <div class="modal-panel-list">
                  <el-tag
                    v-if="form.expression"
                    size="medium"
                    type="info"
                    color="#f7f7f7"
                    effect="plain"
                    closable>
                    {{form.expression}}
                  </el-tag>
                </div>
              </dd>
            </dl>
          </div>
        </div>
      </div>
    </section>
    <footer class="modal-footer">
      <el-button
        size="small"
        @click="cancel">取 消</el-button>
      <el-button 
        type="primary" 
        size="small"
        @click="save">确 定</el-button>
    </footer>
  </section>

  <!-- 人员选择组件 -->
  <personnel
    :edit="personnels.edit"
    :show="personnels.show"
    :columns="personnels.columns"
    :showResult="personnels.showResult"
    @cancel="personCancel"
    @confirm="personConfirm" />
  
  <!-- 菜单弹窗 -->
  <el-dialog
    title="菜单位置设置"
    :show-close="false"
    :visible.sync="dialog.menu"
    :append-to-body="true"
    class="xa-dialog">  
    <div>
      <span>选择父级菜单：</span>
      <el-cascader
        style="width:300px"
        placeholder="请选择"
        :options="menuList"
        :props="{value: 'id',label: 'name'}"
        v-model="form.menus"
        filterable
        clearable
        ref="menuCascader"
        @change="menuChange"></el-cascader>
    </div>
    <div slot="footer">
      <el-button
        size="small"
        @click="dialog.menu =false">取 消</el-button>
      <el-button 
        type="primary" 
        size="small"
        @click="menuConfirm">确 定</el-button>  
    </div>
  </el-dialog>  

  <!-- 流程图标设置 -->
  <el-dialog
    title="流程图标设置"
    :show-close="false"
    :visible.sync="dialog.icon"
    :append-to-body="true"
    class="xa-dialog--sm">  
    <ul class="modal-panel-picture">
      <li class="icon">
        <img src="../../../assets/images/workflow/icon.png"/>
      </li>
    </ul>
    <div slot="footer">
      <el-button
        size="small"
        @click="dialog.icon =false">取 消</el-button>
      <el-button 
        type="primary" 
        size="small"
        @click="iconConfirm">确 定</el-button>  
    </div>
  </el-dialog>  

  <!-- 选择应用 -->
  <el-dialog
    title="选择应用"
    :show-close="false"
    :visible.sync="dialog.apply"
    :append-to-body="true"
    class="xa-dialog">  
    <div>
      <span>选择应用：</span>
      <el-cascader
        style="width:300px"
        placeholder="请选择"
        :options="menuList"
        :props="{value: 'id',label: 'name'}"
        v-model="form.applys"
        filterable
        clearable
        ref="applyCascader"
        @change="applyChange"></el-cascader>
    </div>
    <div slot="footer">
      <el-button
        size="small"
        @click="dialog.apply =false">取 消</el-button>
      <el-button 
        type="primary" 
        size="small"
        @click="applyConfirm">确 定</el-button>  
    </div>
  </el-dialog> 

  <!-- 选择流程数据指标 -->
  <el-dialog
    title="选择流程数据指标"
    :show-close="false"
    :visible.sync="dialog.target"
    :append-to-body="true"
    class="xa-dialog">  
    <div>
      <span>选择指标：</span>
      <el-cascader
        style="width:300px"
        placeholder="请选择"
        :options="menuList"
        :props="{value: 'id',label: 'name'}"
        v-model="form.targets"
        filterable
        clearable
        ref="targetCascader"
        @change="targetChange"></el-cascader>
    </div>
    <div slot="footer">
      <el-button
        size="small"
        @click="dialog.target =false">取 消</el-button>
      <el-button 
        type="primary" 
        size="small"
        @click="targetConfirm">确 定</el-button>  
    </div>
  </el-dialog> 

  <!-- 选择指标类型 -->
  <el-dialog
    title="选择指标类型"
    :show-close="false"
    :visible.sync="dialog.targetType"
    :append-to-body="true"
    class="xa-dialog">  
    <div>
      <span>选择指标类型：</span>
      <el-select 
        v-model="form.targetType" 
        style="width:300px"
        filterable 
        placeholder="请选择"
        clearable>
        <el-option>全部</el-option>
        <el-option
          v-for="item in targetTypeList"
          :key="item.id"
          :value="item.name">
          {{item.name}}
          <span style="color:#999;">({{item.desc}})</span>
        </el-option>
      </el-select>
    </div>
    <div slot="footer">
      <el-button
        size="small"
        @click="dialog.targetType =false">取 消</el-button>
      <el-button 
        type="primary" 
        size="small"
        @click="targetTypeConfirm">确 定</el-button>  
    </div>
  </el-dialog> 

  <!-- 设置表达式 -->
  <el-dialog
    title="设置表达式"
    :show-close="false"
    :visible.sync="dialog.expression"
    :append-to-body="true"
    class="xa-dialog">  
    <div>
      <el-row :gutter="20">
        <el-col :span="8">
          <div class="expression-line">
            {{form.targetType}}
          </div>
        </el-col>
        <el-col :span="6">
          <el-select 
            v-model="expressionKey" 
            style="width:160px"
            size="small"
            placeholder="请选择">
            <el-option
              v-for="item,index in expressionList"
              :key="index"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-col>
        <el-col :span="8">
          <el-input
            v-model="expressionVal"
            min="0"
            placeholder="请输入目标数值"
            size="small" />
        </el-col>
      </el-row>
      <div class="expression-result">
        当前表达式：{{form.targetType}} {{expressionKey}} {{expressionVal}}
      </div>
    </div>
    <div slot="footer">
      <el-button
        size="small"
        @click="dialog.expression =false">取 消</el-button>
      <el-button 
        type="primary" 
        size="small"
        @click="expressionConfirm">确 定</el-button>  
    </div>
  </el-dialog> 
</section>
</template>

<script>
import personnel from '@/components/Personnel'   
export default {
  name: 'add-node-sponsor',
  components: {
    personnel
  },
  props: {
    show: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: ''
    },
    node: {
      type: Object,
      default: null
    }
  },
  data () {
    return {
      visible: false,
      temp: {},
      flows: [{
        label: '常规流程',
        value: '1'
      }, {
        label: '应用流程',
        value: '2'
      }, {
        label: '数据驱动流程',
        value: '3'
      }],
      approvers: [{ 
        label: '指定成员', 
        value: '1'
      }, { 
        label: '人员职务', 
        value: '2'
      }],
      menuList:[{
        id: '1000',
        name: '销售话术',
        children: [{
          id: '1001',
          name: '销售话术列表',
        }, {
          id: '1002',
          name: '销售话术添加',
        }]
      }],
      targetTypeList: [{
        id: '1000',
        name: '移动开卡',
        desc: '指标值'
      },{
        id: '1001',
        name: '较上月平均',
        desc: '指标趋势'
      }],
      expressionKey: '',
      expressionVal: '',
      expressionList: [
        { label: '小于', value: '<' },
        { label: '小于等于', value: '<=' },
        { label: '等于', value: '==' },
        { label: '大于', value: '>' },
        { label: '大于等于', value: '>=' }
      ],
      personnels: {
        show: false,
        edit: false,
        showResult: false,
        columns: []
      },
      dialog: {
        menu: false,
        icon: false,
        apply: false,
        target: false,
        targetType: false,
        expression: false
      },
      form: {
        flow: '1',
        approver: '1',
        persons:[],
        roles:[],
        menus:[],
        applys: [],
        targets:[],
        targetType:'',
        expression: ''
      },
    }
  },
  watch: {
    'show' (val) {
      this.visible = val
      this.showModal(this.visible)
    },
    visible (val) {
      this.$emit('update:show', val)
    }
  },
  methods: {
    personCancel(){
      this.personnels.show = false
    },
    personConfirm(paramArr) {
      this.personnels.show = false
      console.log(paramArr);
    },
    menuChange(item) {
      this.form.menus = this.$refs['menuCascader'].getCheckedNodes()[0].pathLabels
    },
    menuConfirm() {
      this.dialog.menu = false
    },
    iconConfirm() {
      this.dialog.icon = false
    },
    applyChange(item) {
      this.form.applys = this.$refs['applyCascader'].getCheckedNodes()[0].pathLabels
    },
    applyConfirm() {
      this.dialog.apply = false
    },
    targetChange(item) {
      this.form.targets = this.$refs['targetCascader'].getCheckedNodes()[0].pathLabels
    },
    targetConfirm() {
      this.dialog.target = false
    },
    targetTypeConfirm() {
      this.dialog.targetType = false
    },
    expressionConfirm() {
      this.dialog.expression = false
      this.form.expression = this.form.targetType + this.expressionKey + this.expressionVal
    },
    save() {
      this.visible = false
      Object.assign(this.temp, this.form)

      let len = this.form.persons.length
      if(len==0) {
        this.$alert('请设置发起人！')
      } else if(len==1) {
        this.node.content = this.form.persons[0].name
      } else {
        this.node.content = this.form.persons.length + '人'
      }
    },
    cancel() {
      this.visible = false
    },
    init () {
      this.form = this.node.properties ? this.node.properties : {
        flow: '1',
        approver: '1',
        persons:[],
        roles:[],
        menus:[],
        applys: [],
        targets:[],
        targetType:'',
        expression: ''
      }
    },
    showModal(bool) {
      const modal = this.$refs.sponsor
      const body = document.querySelector('body')
      if(bool) {
        this.init()
        body.appendChild(modal)
      }else {
        body.removeChild(modal)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import "../../../styles/_global.scss";
@import '../modal.scss';

.expression {

  &-line {
    padding-bottom:8px;
    border-bottom:1px solid #ddd;
    color:$color-96;
  }

  &-result {
    color:$color-primary;
    padding-top:40px;
  }
}
</style>
