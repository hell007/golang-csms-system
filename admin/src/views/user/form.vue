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
        <el-form-item label="姓名" prop="username">
          <el-input class="p-form__input" 
            placeholder="请输入" 
            clearable
            v-model="form.username"
            :maxlength="formRules.username[1].max" >
            <span class="p-input-count" slot="suffix">
            {{form.username.length}} / {{formRules.username[1].max}}
            </span>
          </el-input>
        </el-form-item>
        <el-form-item label="角色组" prop="roleId">
            <el-select 
              clearable
              v-model="form.roleId" 
              placeholder="请选择">
              <el-option
                v-for="item in options.roles"
                :key="item.rolename"
                :label="`${item.rolename} => ${item.rolenote}`"
                :value="item.id">
              </el-option>
            </el-select>
          </el-form-item>
        <el-form-item label="手机号" prop="mobile">
          <el-input class="p-form__input" 
            placeholder="请输入" 
            clearable
            maxlength="11"
            v-model="form.mobile"></el-input>
        </el-form-item>        
        <el-form-item label="邮箱" prop="email">
          <el-input class="p-form__input" 
            placeholder="请输入"
            clearable
            v-model="form.email" ></el-input>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio
              v-for="item, index in options.status"
              :key="item.value"
              :label="item.value">{{item.label}}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="创建时间" prop="createTime" >
          <el-date-picker class="p-form__picker"
            v-model="form.createTime" 
            type="datetime" 
            format="yyyy-MM-dd HH:mm:ss" 
            placeholder="选择日期时间">
          </el-date-picker>
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

import {fetchGet, fetchPost} from '@/api'
import {validateMobile} from '@/utils/validate'

export default {
  name: 'user-form',
  components: {},
  data() {
    return {
      processing: false,
      form: {
        id: '',
        username: '',
        roleId: null,
        mobile: '',
        email: '',
        status: 1,// int  注意不是字符串
        createTime:''
      },
      options: {
        status: [{
          value: 1,
          label: '启用'
        },{
          value: 2,
          label: '停用'
        }],
        roles: []
      },
      //验证规则
      formRules: {
        username: [{
          required: true,
          message: "请输入姓名",
          trigger: 'blur'
        }, {
          min: 2,
          max: 20,
          message: "必须输入2-20个字符",
          trigger: 'blur'
        }],
        roleId: [{
          required: true,
          message: "请选择",
          trigger: 'blur'
        }],
        mobile: [{
          required: true,
          message: "请输入手机号码",
          trigger: 'blur'
        }, {
          validator: validateMobile,
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
    //角色
    getRoles() {
      const self = this
      const listQuery = {
        pageNumber: 1,
        pageSize: 10,
        status: 1
      }

      fetchGet('/sys/role/list', listQuery).then(response => {
        const status = response.data.state
        const res = response.data.data.rows
        if (status) {
          self.options.roles = res
        }
      }).catch(ex => {
        self.$notify({
          title: '请求错误',
          message: ex,
          type: 'error'
        })
      })
    },
    //根据id获取数据
    getItem() {
      const self = this

      fetchGet('/sys/user/item', {
        id: self.form.id
      }).then(response => {
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
      }).catch(ex => {
        self.$notify({
          title: '请求错误',
          message: ex,
          type: 'error'
        })
      })
    },
    //表单提交
    handleSubmit() {
      const self = this
      this.$refs.postForm.validate(valid => {
        if (valid) {
          self.processing = true

          fetchPost('/sys/user/save', self.form).then(response => {
            const status = response.data.state
            //const res = response.data.data
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
          }).catch(ex => {
            self.$notify({
              title: '请求错误',
              message: ex,
              type: 'error'
            })
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
      this.$router.push('/user/list')
    }
  },
  created() {
    this.getRoles();
    if (this.isEdit) {
      this.getItem();
    }
  }
}
</script>