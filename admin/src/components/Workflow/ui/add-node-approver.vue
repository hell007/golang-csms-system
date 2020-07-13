<template>
  <el-dialog
    :append-to-body="true"
    class="xa-dialog"
    :title="title"
    :visible.sync="visible">
    <div class="approver-panel">
      <div class="approver-panel-item">
        <h5 class="approver-panel-hd">
          选择审批对象
        </h5>
        <div class="approver-panel-bd">
          <el-radio-group 
            v-model="currentApp">
            <el-radio 
              v-for="(a, index) in approvers"
              :label="a.value"
              @change="setApprover(a)">{{ a.label }}</el-radio>
          </el-radio-group>
          <div class="approver-actions"
            v-if="currentApp === 'target_label' && !showAddRole">
            <el-button 
              type="primary" 
              size="small"
              icon="el-icon-plus"
              @click="addRole">添加角色</el-button>
          </div>
          <div
            v-if="showAddRole"
            class="ant-row-flex ant-row-flex-space-around ant-row-flex-middle condition-group">
            角色名&nbsp;&nbsp;
            <div
              class="ant-select ant-select-enabled"
              style="min-width: 150px;">
              <div class="ant-input-number-input-wrap">
                <input
                  v-model="form.actionerRules[0].labelNames"
                  class="ant-input-number-input"
                  placeholder="输入角色名">
              </div>
            </div>
          </div>
          <div class="approver-actions"
            v-if="currentApp === 'target_management'">
            <span class="approver-actions-label">发起人的</span>
            <el-select 
              size="small"
              v-model="form.actionerRules[0].level" 
              placeholder="请选择">
              <el-option
                v-for="item in managements"
                :key="item.id"
                :label="item.name"
                :value="item.id">
              </el-option>
            </el-select>
          </div>
        </div>
      </div>

      <div class="approver-panel-item">
        <h5 class="approver-panel-hd">
          多人审批时采用的审批方式
        </h5>
        <div class="approver-panel-bd">
          <el-radio-group 
            v-model="currentAction">
            <el-radio 
              v-for="(a, index) in actTypes"
              :label="a.value"
              @change="setAction(a)">{{ a.label }}</el-radio>
          </el-radio-group>
        </div>
      </div>
      <p>{{ form }}</p>
    </div>
    <span slot="footer" class="dialog-footer">
      <el-button 
        type="info"
        @click="cancel">取 消</el-button>
      <el-button 
        type="primary"
        @click="save">保 存</el-button>  
    </span>
  </el-dialog>
</template>
<script>
export default {
  name: 'add-node-approver',
  props: {
    show: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: ''
    },
    properties: {
      type: Object,
      default: undefined
    }
  },
  data: () => ({
    visible: false,
    showAddRole: false,
    currentApp: 'target_management',
    currentAction: 'or',
    temp: {},
    approvers: [
      { label: '主管', value: 'target_management'},
      { label: '角色', value: 'target_label'}
    ],
    managements: [
      {id: 1, name: '直接主管'}, 
      { id: 2, name: '第二级主管'}
    ],
    actTypes: [
      { label: '或签（一名审批人同意或拒绝即可）', value: 'or'},
      { label: '会签（须所有审批人同意）', value: 'and'}
    ],
    form: {
      actionerRules: [{
        type: 'target_management',
        level: 1,
        isEmpty: false,
        autoUp: true,
        actType: 'or'
      }]
    }
  }),
  watch: {
    'show' (val) {
      this.visible = val
    },
    visible (val) {
      this.$emit('update:show', val)
    }
  },
  mounted () {
    this.form = this.properties
    this.init()
    Object.assign(this.temp, this.form)
  },
  methods: {
    init () {
      this.form = this.form ? this.form : {
        actionerRules: [
          {
            type: 'target_management',
            level: 1,
            isEmpty: false,
            autoUp: true,
            actType: 'or'
          }
        ]
      }
      var rule = this.form.actionerRules && this.form.actionerRules[0]
      if (rule) {
        this.currentApp = rule.type
        this.currentAction = rule.actType
        if (rule.labelNames) this.showAddRole = true
      }
    },
    save () {
      var rule = this.form.actionerRules[0]
      switch (rule.type) {
        case 'target_label':
          if (!rule.labelNames || rule.labelNames === '') {
            alert('角色不能为空')
            return
          }
          break
      }
      this.visible = false
      Object.assign(this.temp, this.form)
      this.$emit('setProperties', this.form)
    },
    cancel () {
      this.visible = false
      this.form = {}
      Object.assign(this.form, this.temp)
      this.init()
      this.$emit('setProperties', this.form)
    },
    setApprover (row) {
      this.currentApp = row.value
      if (row.value === 'target_label') {

      } else {
        this.showAddRole = false
      }
      this.form.actionerRules = []
      switch (row.value) {
        case 'target_management':
          this.form.actionerRules.push({
            type: 'target_management',
            level: 1,
            isEmpty: false,
            autoUp: true
          })
          break
        case 'target_label':
          this.form.actionerRules.push({
            type: 'target_label',
            labelNames: '',
            labels: '',
            isEmpty: false,
            memberCount: 1,
            actType: 'or'
          })
          break
        default:
      }
    },
    setAction (act) {
      this.currentAction = act.value
      this.form.actionerRules[0].actType = act.value
    },
    addRole () {
      this.showAddRole = true
    }
  }
}
</script>
<style lang="scss" scoped>
.approver {
  
  &-panel {

    &-item {
      padding:0 15px;
      margin-bottom:15px;
      border-bottom:1px solid $color-border-gray;
    }

    &-hd {
      font-size: 14px;
      color: #333;
      margin-bottom:0;
    }

    &-bd {
      padding:15px 0;
    }
  }

  &-actions {
    padding-top:10px;

    &-label {
      margin-right:10px;
    }
  }
}  
</style>
