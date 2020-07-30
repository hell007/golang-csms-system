<template>
<section ref="approver" class="modal"
  v-bind:class="{'modal-show':visible}">
  <section class="modal-right">
    <header class="modal-header">{{ title }}</header>
    <section class="modal-body">
      <div class="modal-panel">
        <div class="modal-panel-item">
          <h5 class="modal-panel-hd">设置下一节点人</h5>
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
        <div class="modal-panel-item">
          <h5 class="modal-panel-hd">多人审批时采用的审批方式</h5>
          <div class="modal-panel-bd">
            <el-radio-group
              class="el-radio-column"
              v-model="form.action">
              <el-radio 
                v-for="(a, index) in actions"
                :key="a.label"
                :label="a.value">{{ a.label }}</el-radio>
            </el-radio-group>
          </div>
        </div>
        <div class="modal-panel-item">
          <h5 class="modal-panel-hd">是否同步发送短信提醒</h5>
          <div class="modal-panel-bd">
            <el-radio-group
              class="el-radio-column"
              v-model="form.msg">
              <el-radio 
                v-for="(a, index) in msgs"
                :key="a.label"
                :label="a.value">{{ a.label }}</el-radio>
            </el-radio-group>
          </div>
        </div>
        <div class="modal-panel-item"
          v-if="form.approver=='2'">
          <h5 class="modal-panel-hd">审批人为空时</h5>
          <div class="modal-panel-bd">
            <el-radio-group
              class="el-radio-row"
              v-model="form.auto">
              <el-radio
                label="1">自动通过</el-radio>
            </el-radio-group>
          </div>
        </div>
        <div class="modal-panel-item">
          <h5 class="modal-panel-hd">节点处理时长</h5>
          <div class="modal-panel-bd">
            <el-radio-group
              class="el-radio-row"
              v-model="form.time">
              <el-radio 
                v-for="(a, index) in times"
                :key="a.label"
                :label="a.value">{{ a.label }}</el-radio>
            </el-radio-group>
          </div>
        </div>
        <div class="modal-panel-item">
          <h5 class="modal-panel-hd">节点处理超时</h5>
          <div class="modal-panel-bd">
            <el-radio-group
              class="el-radio-column"
              v-model="form.timeout">
              <el-radio 
                v-for="(a, index) in timeouts"
                :key="a.label"
                :label="a.value">{{ a.label }}</el-radio>
            </el-radio-group>
          </div>
        </div>
        <div class="modal-panel-item"
          v-if="form.approver=='1'">
          <h5 class="modal-panel-hd">提醒短信发送频率</h5>
          <div class="modal-panel-bd">
            <el-radio-group
              class="el-radio-row"
              v-model="form.frequency">
              <el-radio 
                v-for="(a, index) in  frequencys"
                :key="a.label"
                :label="a.value">{{ a.label }}</el-radio>
            </el-radio-group>
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
</section>
</template>

<script>
import personnel from '@/components/Personnel'   
export default {
  name: 'add-node-approver',
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
      approvers: [{ 
        label: '指定成员', 
        value: '1'
      }, { 
        label: '人员职务', 
        value: '2'
      }],
      actions: [{
        label: '会签（须所有审批人同意）',
        value: 'and'
      },{
        label: '或签（一名审批人同意或拒绝即可）',
        value: 'or'
      }],
      msgs: [{
        label: '是',
        value: '1'
      }, {
        label: '否',
        value: '0'
      }],
      times: [{
        label: '12小时',
        value: '12'
      }, {
        label: '24小时',
        value: '24'
      }, {
        label: '48小时',
        value: '48'
      }, {
        label: '5天',
        value: '120'
      }],
      timeouts: [{
        label: '不处理',
        value: '1'
      }, {
        label: '给处理人发送提醒短信',
        value: '2'
      }, {
        label: '给处理人上级发送提醒短信',
        value: '3'
      }],
      frequencys: [{
        label: '只发一次',
        value: '1'
      }, {
        label: '每12小时',
        value: '2'
      }, {
        label: '每24小时',
        value: '3'
      }, {
        label: '每48小时',
        value: '4'
      }],
      personnels: {
        show: false,
        edit: false,
        showResult: false,
        columns: []
      },
      form: {
        approver: '1',
        action: 'and',
        msg: '1',
        auto: '1',
        time: '12',
        timeout: '1',
        frequency: '1'
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
    save () {
      this.visible = false
      Object.assign(this.temp, this.form)

      if(this.title=='抄送人') {
        this.node.notifier.properties = this.form
      } else {
        this.node.properties = this.form
      }
    },
    cancel () {
      this.visible = false
    },
    init () {
      if(this.title=='抄送人') {
        this.form = this.node.notifier.properties ? this.node.notifier.properties : {
            approver: '1',
            action: 'and',
            msg: '1',
            auto: '1',
            time: '12',
            timeout: '1',
            frequency: '1'
          }
      } else {
        this.form = this.node.notifier.properties ? this.node.notifier.properties : {
            approver: '1',
            action: 'and',
            msg: '1',
            auto: '1',
            time: '12',
            timeout: '1',
            frequency: '1'
          }
      }
    },
    showModal(bool) {
      const modal = this.$refs.approver
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
@import '../modal.scss';
</style>
