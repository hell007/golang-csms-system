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
        <el-form-item label="名称" prop="categoryName">
          <el-input class="p-form__input" 
            placeholder="请输入"
            clearable 
            v-model="form.categoryName"
            :maxlength="formRules.categoryName[1].max" >
            <span class="p-input-count" slot="suffix">
            {{form.categoryName.length}} / {{formRules.categoryName[1].max}}
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
              :key="item.categoryName"
              :label="item.categoryName"
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
        <el-form-item label="导航显示" prop="isNav" >
          <el-radio-group v-model="form.isNav">
            <el-radio-button
              v-for="item, index in options.navs"
              :key="item.value"
              :label="item.value">{{item.label}}</el-radio-button>
          </el-radio-group>
        </el-form-item>
      </div>
      <hr class="xa-separator" />
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
  name: 'category-form',
  components: {},
  data() {
    return {
      processing: false,
      form: {
        id: '',
        categoryName: '',
        pid: null,
        status: 1,
        isNav: 2,
      },
      options: {
        status: [{
          value: 1,
          label: '启用'
        },{
          value: 2,
          label: '停用'
        }],
        navs: [{
          value: 1,
          label: '显示'
        },{
          value: 2,
          label: '隐藏'
        }],
        nodes: []
      },
      //验证规则
      formRules: {
        categoryName: [{
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
        status: [{
          required: true,
          message: "请选择状态",
          trigger: 'blur'
        }],
        isNav: [{
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
      fetchGet('/goods/category/list', {pid:0}).then(response => {
        const status = response.data.state
        const res = response.data.data
        if (status) {
          self.options.nodes = [{id:0, categoryName:'顶级菜单'}, ...res]
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
      fetchGet('/goods/category/item', {
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

          fetchPost('/goods/category/save', self.form).then(response => {
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
      this.$router.push('/category/list')
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