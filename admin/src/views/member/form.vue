<template>
  <el-card class="xa-block" shadow="hover">
    <el-form class="p-form"
      label-suffix="："
      label-width="120px">
      <div class="p-form__section">
        <div class="p-form__s-head">会员信息</div>
        <el-form-item label="姓名">
          <span>{{form.name}}</span>
        </el-form-item>
        <el-form-item label="性别">
          <span>{{form.gender}}</span>
        </el-form-item>
        <el-form-item label="手机号">
          <span>{{form.mobile}}</span>
        </el-form-item>
        <el-form-item label="状态">
          <span v-for="item in options.status">
            {{form.status == item.value ? item.label : ''}}
          </span>
          <span>&nbsp;&nbsp;&nbsp;&nbsp;</span>
          <el-button 
            size="small" 
            type="danger" 
            icon="el-icon-edit"
            @click="dialog.visible = true">
            编辑
          </el-button>
        </el-form-item>
        <el-form-item label="创建时间">
          <span>{{form.createTime | parseTime}}</span>
        </el-form-item>
        <el-form-item label="登录时间">
          <span>{{form.loginTime | parseTime}}</span>
        </el-form-item>
        <el-form-item class="xa-text-right">
          <el-button
            type="primary"
            size="small"
            icon="el-icon-back"
            @click="handleCancle">返回</el-button>
        </el-form-item>
      </div>
    </el-form>

    <!-- 会员地址 -->
    <hr class="xa-separator" />
    <div class="p-form__section">
      <div class="p-form__s-head">会员地址</div>
      <div class="xa-block">
        <el-table
          class="p-table"
          :data="list"
          border
          ref="table-product"
          size="small">
          <el-table-column 
            label="收货人姓名" 
            prop="consignee"
            align="center">
          </el-table-column>
          <el-table-column 
            label="联系方式" 
            align="center">
            <template slot-scope="scope">
              {{scope.row.phone}} {{scope.row.tell}}
            </template>
          </el-table-column>
          <el-table-column
            label="详细地址"
            align="center">
            <template slot-scope="scope">
              {{scope.row.province}}{{scope.row.city}}{{scope.row.district}}{{scope.row.address}}
            </template>
          </el-table-column>
          <el-table-column 
            label="邮编"
            prop="zipcode" 
            width="100" 
            align="center">
          </el-table-column>
        </el-table>
      </div>
    </div>

    <el-dialog
      title="编辑会员"
      :visible.sync="dialog.visible"
      class="xa-dialog xa-dialog--sm">
      <el-form
        class="p-form"
        ref="postForm"
        :rules="formRules"
        :model="form"
        label-width="120px"
        label-suffix="："
        size="small">
        <div class="p-form__section">
          <div class="p-form__s-head">会员信息</div>
          <el-form-item label="姓名">
            <span>{{form.name}}</span>
          </el-form-item>
          <el-form-item
            prop="status"
            label="状态">
            <el-radio-group v-model="form.status">
              <el-radio-button
                v-for="item, index in options.status"
                :key="`state-${item.value}`"
                :label="item.value">{{item.label}}</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item 
            prop="note"
            label="备注">
            <el-input v-model="form.note" type="textarea"></el-input>
          </el-form-item>
        </div>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button
          type="primary"
          :loading="dialog.processing"
          @click="handleSubmit">确定</el-button>
         <el-button 
          @click="dialog.visible = false">取消</el-button> 
      </span>
    </el-dialog>  
  </el-card>
</template>

<script>
import {
  mapActions
} from 'vuex'

export default {
  name: 'memberform',
  components: {},
  data() {
    return {
      form: {},
      list: [],
      dialog: {
        visible: false,
        processing: false,
      },
      options: {
        status: [{
          value: 1,
          label: '已认证'
        },{
          value: 2,
          label: '未认证'
        },{
          value: 3,
          label: '已注销'
        }]
      },
      //验证规则
      formRules: {
        status: [{
          required: true,
          message: "请选择",
          trigger: 'blur'
        }],
        note: [{
          required: true,
          message: "请输入注销原因",
          trigger: 'blur'
        }]
      }
    }
  },
  computed: {
    isEdit() {
      this.form.id = this.$route.query.id ? this.$route.query.id : null
      return this.form.id
    }
  },
  methods: {
    ...mapActions(['getMember', 'saveMember']),
    getItem() {
      const self = this
      self.getMember(self.form.id).then(response => {
        const status = response.data.state
        const res = response.data.data
        if (status) {
          self.form = res.member
          self.list = res.list
        } else {
          self.$notify({
            title: '失败',
            message: response.data.message,
            type: 'error'
          })
        } 
        self.loading = false
      })
    },
    handleSubmit() {
      const self = this
      self.$refs.postForm.validate(valid => {
        if (valid) {
          let form = self.form
          self.dialog.processing = true
          self.$confirm(`确定要更改【${form.name}】会员?`, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          })
          .then(function(action) { 
            self.saveMember(form).then(response => {
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
              self.dialog.processing = false
            })
          })
          .catch(function(action) {})

        } else {
          self.$alert('请正确输入！', '提示', {
            confirmButtonText: '确定'
          });
          return false
        }
      })
    },
    handleCancle() {
      this.$router.push('/member/list')
    }
  },
  created() {
    if (this.isEdit) {
      this.getItem();
    }
  }
}
</script>