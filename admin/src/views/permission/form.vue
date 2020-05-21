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
        <el-form-item label="名称" prop="name">
          <el-input class="p-form__input" 
            placeholder="请输入" 
            clearable
            v-model="form.name"
            :maxlength="formRules.name[1].max" >
            <span class="p-input-count" slot="suffix">
            {{form.name.length}} / {{formRules.name[1].max}}
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
        <el-form-item label="路径" prop="path">
          <el-input class="p-form__input" 
            clearable
            placeholder="请输入" 
            v-model="form.path"></el-input>
        </el-form-item> 
        <el-form-item label="icon" prop="icon">
          <el-input class="p-form__input" 
            placeholder="请输入" 
            clearable
            v-model="form.icon"></el-input>
        </el-form-item>       
        <el-form-item label="auth" prop="code">
          <el-input class="p-form__input" 
            placeholder="请输入"
            clearable
            v-model="form.code" ></el-input>
        </el-form-item>
        <el-form-item label="跳转" prop="redirect">
          <el-input class="p-form__input" 
            placeholder="请输入" 
            clearable
            v-model="form.redirect"></el-input>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio
              v-for="item, index in options.status"
              :key="item.value"
              :label="item.value">{{item.label}}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="菜单栏显示" prop="hidden" >
          <el-radio-group 
            v-model="form.hidden" 
            size="small">
            <el-radio-button
              v-for="item, index in options.hidden"
              :key="item.value"
              :label="item.value">{{item.label}}</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="级别" prop="level" >
          <el-radio-group 
            v-model="form.level"
            size="small">
            <el-radio-button
              v-for="item, index in options.level"
              :key="item.value"
              :label="item.value">{{item.label}}</el-radio-button>
          </el-radio-group>
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

export default {
  name: 'permission-form',
  components: {},
  data() {
    return {
      processing: false,
      form: {
        id: '',
        name: '',
        pid: null,
        path: '',
        icon: '',
        code: '',
        redirect: '',
        status: 1,
        hidden: 2,
        level: 3,
      },
      options: {
        status: [{
          value: 1,
          label: '启用'
        },{
          value: 2,
          label: '停用'
        }],
        hidden: [{
          value: 1,
          label: '显示'
        },{
          value: 2,
          label: '隐藏'
        }],
        level: [{
          value: 1,
          label: '一级'
        },{
          value: 2,
          label: '二级'
        },{
          value: 3,
          label: '三级'
        }],
        nodes: []
      },
      //验证规则
      formRules: {
        name: [{
          required: true,
          message: "请输入",
          trigger: 'blur'
        }, {
          min: 2,
          max: 32,
          message: "必须输入2-32个字符",
          trigger: 'blur'
        }],
        pid: [{
          required: true,
          message: "请选择",
          trigger: 'blur'
        }],
        level: [{
          required: true,
          message: "请选择",
          trigger: 'blur'
        }],
        status: [{
          required: true,
          message: "请选择状态",
          trigger: 'blur'
        }],
        hidden: [{
          required: true,
          message: "请选择",
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
    getNodes() {
      const self = this
      fetchGet('/sys/permission/list', {level:2}).then(response => {
        const status = response.data.state
        const res = response.data.data.rows
        if (status) {
          self.options.nodes = [{id:0, name:'顶级菜单'}, ...res]
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
      fetchGet('/sys/permission/item', {
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
          fetchPost('/sys/permission/save', self.form).then(response => {
            const status = response.data.state
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
      this.$router.push('/permission/list')
    }
  },
  created() {
    this.getNodes();
    if (this.isEdit) {
      this.getItem();
    }
  }
}
</script>