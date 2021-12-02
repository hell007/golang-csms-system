<template>
  <el-card class="xa-block" shadow="hover">
    <el-form class="p-form"
      ref="postForm"
      :rules="formRules"
      :model="form"
      label-suffix="："
      label-width="120px">

      <div class="p-form__section">
        <div class="p-form__s-head">表单信息</div>
        <el-form-item label="角色名" prop="rolename">
          <el-input class="p-form__input" 
            placeholder="请输入" 
            clearable
            v-model="form.rolename"
            :maxlength="formRules.rolename[1].max" >
            <span class="p-input-count" slot="suffix">
            {{form.rolename.length}} / {{formRules.rolename[1].max}}
            </span>
          </el-input>
        </el-form-item>
        <el-form-item label="上级" prop="pid">
          <el-select 
            clearable
            v-model="form.pid" 
            placeholder="请选择">
            <el-option
              v-for="item in options.nodes"
              :key="item.name"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item> 
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio
              v-for="item, index in options.status"
              :key="item.value"
              :label="item.value">{{item.label}}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="描述" prop="remark">
          <el-input 
            class="p-form__input" 
            type="textarea" 
            placeholder="请输入"
            v-model="form.remark" ></el-input>
        </el-form-item>
      </div>
      <hr class="p-separator" />
      <el-form-item>
        <el-button type="primary"
          :loading="processing"
          @click="handleSubmit">提交</el-button>
        <el-button
          @click="handleCancle">返回</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script>
import {
  mapActions
} from 'vuex'

import {
  validateMobile
} from '@/utils/validate' //验证规则

export default {
  name: 'roleForm',
  components: {},
  data() {
    return {
      processing: false,
      form: {
        id: '',
        rolename: '',
        pid: 2,
        status: 1,// int  注意不是字符串
        remark:''
      },
      options: {
        status: [{
          value: 1,
          label: '启用'
        },{
          value: 2,
          label: '停用'
        }],
        nodes: [{
          id: 1,
          name: 'superadmin'
        },{
          id: 2,
          name: 'admin'
        }]
      },
      //验证规则
      formRules: {
        rolename: [{
          required: true,
          message: "请输入姓名",
          trigger: 'blur'
        }, {
          min: 2,
          max: 20,
          message: "必须输入2-20个字符",
          trigger: 'blur'
        }],
        status: [{
          required: true,
          message: "请选择状态",
          trigger: 'blur'
        }]
      }
    }
  },
  //实时计算属性
  computed: {
    //判断是否处于编辑状态
    isEdit() {
      this.form.id = this.$route.query.id ? this.$route.query.id : null
      return this.form.id
    }
  },
  methods: {
    ...mapActions(['getRole', 'saveRole']),
    //根据id获取数据
    getItem() {
      const self = this
      self.getRole(self.form.id).then(response => {
        const status = response.data.state
        const res = response.data.data
        const message = response.data.msg
        if (status) {
          self.form = res
        } else {
          self.$notify({
            title: '失败',
            message: message,
            type: 'error'
          })
        } 
        self.loading = false
      })
    },
    //表单提交
    handleSubmit() {
      const self = this
      this.$refs.postForm.validate(valid => {
        if (valid) {
          self.processing = true
          self.saveRole(this.form).then(response => {
            const status = response.data.state
            const res = response.data.data
            const message = response.data.msg
            if (status) {
              self.$notify({
                title: '成功',
                message: message,
                type: 'success',
                duration: 2000
              })
            } else {
              self.$notify({
                title: '失败',
                message: message,
                type: 'error'
              })
            } 
            self.processing = false
          })
        } else {
          self.$alert('请正确输入！', '提示', {
            confirmButtonText: '确定'
          });
          return false
        }
      })
    },
    handleCancle() {
      this.$router.push('/role/list')
    }
  },
  created() {
    if (this.isEdit) {
      this.getItem();
    }
  }
}
</script>