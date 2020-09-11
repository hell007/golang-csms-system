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
        <el-form-item label="头像"  prop="avatar">
          <el-upload
            class="p-uploader"
            style="width:80px;height:80px"
            :action="uploadApi"
            :show-file-list="false"
            :on-success="handleAvatarUpload"
            :before-upload="handleBeforeUpload"
            accept="image/jpeg, image/png">
            <img class="p-uploader__img" v-if="form.avatar" :src="form.avatar">
            <i v-else class="el-icon-plus p-uploader__icon"></i>
          </el-upload>
          <div class="p-form__tip">建议上传PNG格式透明背景图片，图片大小80px * 80px</div>
        </el-form-item>

        <el-form-item label="姓名" prop="name">
          <el-input class="p-form__input" 
            placeholder="请输入" 
            v-model="form.name"
            :maxlength="formRules.name[1].max" >
            <span class="p-input-count" slot="suffix">
            {{form.name.length}} / {{formRules.name[1].max}}
            </span>
          </el-input>
        </el-form-item>

        <el-form-item label="性别" prop="sex">
          <el-radio-group v-model="form.sex">
            <el-radio
              v-for="item, index in options.sex"
              :key="item.value"
              :label="item.value">{{item.label}}</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="多选框" prop="checkbox">
          <el-checkbox-group
            class="p-form__checkbox"
            v-model="form.checkbox">
            <el-checkbox
              v-for="item in options.boxs"
              :label="item"
              :key="item">
              {{item}}
            </el-checkbox>
          </el-checkbox-group>
          <div class="p-form__tip">提示说明文字</div>
        </el-form-item>

        <el-form-item label="角色" prop="roleId">
            <el-select 
              clearable
              v-model="form.roleId" 
              placeholder="请选择">
              <el-option
                v-for="item in options.roles"
                :key="item.roleName"
                :label="item.roleName"
                :value="item.roleId">
              </el-option>
            </el-select>
        </el-form-item>

        <el-form-item label="手机号" prop="mobile">
          <el-input class="p-form__input" 
            placeholder="请输入" 
            maxlength="11"
            v-model="form.mobile"></el-input>
        </el-form-item>   

        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio
              v-for="item, index in options.status"
              :key="item.value"
              :label="item.value">{{item.label}}</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="超级用户" prop="isSuper" >
          <el-radio-group v-model="form.isSuper">
            <el-radio-button label="1">是</el-radio-button>
            <el-radio-button label="2">否</el-radio-button>
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

        <el-form-item label="营业时间">
          <el-select
            v-model="form.weeks"
            style="width:400px"
            multiple
            placeholder="请选择">
            <el-option
              v-for="item in options.weeks"
              :key="item.label"
              :label="item.label"
              :value="item.value"></el-option>
          </el-select>
          <el-time-select
            style="margin-left:20px;"
            clearable
            placeholder="起始时间"
            v-model="form.startTime1"
            :picker-options="{
              start: '08:30',
              step: '00:15',
              end: '18:30'
            }"></el-time-select>
          <el-time-select
            style="margin-left:20px;"
            clearable
            placeholder="结束时间"
            v-model="form.endTime1"
            :picker-options="{
              start: '08:30',
              step: '00:15',
              end: '18:30',
              minTime: form.startTime1
            }"></el-time-select>
        </el-form-item>

        <el-form-item label="日期">
          <el-date-picker
            clearable
            v-model="form.time"
            type="date"
            placeholder="选择日期"></el-date-picker>
        </el-form-item>

        <el-form-item label="日期范围">
          <el-date-picker
            v-model="form.daterange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"></el-date-picker>
        </el-form-item>

        <el-form-item label="区域" prop="area">
          <el-cascader
            class="p-form__input"
            :options="options.area"
            v-model="form.area"
            @change="handleArea">
          </el-cascader>
        </el-form-item>

        <el-form-item label="详细地址" prop="address">
          <el-input 
            class="p-form__input" 
            type="textarea" 
            placeholder="请输入"
            v-model="form.address" ></el-input>
        </el-form-item>

        <el-form-item label="商品重量">
          <el-input
            clearable
            class="p-form__input-small"
            placeholder="请输入"></el-input>
          <el-select
            style="margin-left:20px"
            class="p-form__select-small"
            clearable
            v-model="form.unit"
            placeholder="请选择">
            <el-option
              v-for="item in options.units"
              :key="item.label"
              :label="item.label"
              :value="item.value"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="身份证号" prop="ids">
          <el-input class="p-form__input" 
            placeholder="请输入" 
            v-model="form.ids"
            :maxlength="formRules.ids[1].ids" >
            <span class="p-input-count" slot="suffix">
            {{form.ids.length}} / {{formRules.ids[1].max}}
            </span>
          </el-input>
        </el-form-item>

        <el-form-item label="身份证照片" prop="photos">
          <div class="p-uploader-inline">
            <el-upload
              class="p-uploader"
              style="width: 300px; height: 200px;"
              :action="uploadApi"
              :show-file-list="false"
              :on-success="handleIDFrontUpload"
              :before-upload="noop"
              accept="image/jpeg, image/png">
              <img class="p-uploader__img" v-if="form.photos[0]" :src="form.photos[0]">
              <i v-else class="el-icon-plus p-uploader__icon"></i>
            </el-upload>
            <div>正面</div>
          </div>
          <div class="p-uploader-inline">
            <el-upload
              class="p-uploader"
              style="width: 300px; height: 200px;"
              :action="uploadApi"
              :show-file-list="false"
              :on-success="handleIDBackUpload"
              :before-upload="noop"
              accept="image/jpeg, image/png">
              <img class="p-uploader__img" v-if="form.photos[1]" :src="form.photos[1]">
              <i v-else class="el-icon-plus p-uploader__icon"></i>
            </el-upload>
            <div>背面</div>
          </div>
        </el-form-item>

        <!-- 富文本编辑器使用 -->
        <el-form-item label="文章内容" prop="content" >
          <kindeditor>
            v-model="form.content" 
            :options="options.editor" >
          </kindeditor>  
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
import {
  mapActions
} from 'vuex'

import {
  noop
} from '@/utils'

import {
  validateMobile
} from '@/utils/validate' //验证规则

import kindeditor from '@/components/kindeditor' //富文本编辑器
import axios from 'axios'

export default {
  name: 'unit-form',
  components: {
    kindeditor
  },
  data() {
    return {
      uploadApi: 'http://127.0.0.1:9000/api/test/upload',
      areas:'../assets/project/scripts/area.json',
      processing: false,
      noop: noop,
      form: {
        id: '',
        avatar: '',
        name: '',
        sex: 1,
        checkbox: ['北京'],
        roleId: 1,
        mobile: '',
        content: '',
        status: 1,// int  注意不是字符串
        isSuper: 2,
        createTime:'',
        time: '',
        range: '',
        daterange: '',
        note: '',
        area: [],
        address: '',
        ids: '',
        photos: [],
        unit:''
      },
      options: {
        status: [{
          value: 1,
          label: '启用'
        },{
          value: 2,
          label: '停用'
        }],
        sex: [{
          value: '1',
          label: '男'
        },{
          value: '2',
          label: '女'
        }],
        roles: [{
          roleId: '1',
          roleName: 'admin'
        },{
          roleId: '2',
          roleName: 'goods'
        }],
        boxs: [
          '上海',
          '北京',
          '广州',
          '深圳',
          '2019',
          '2018',
          '2017',
          '2016'
        ],
        units: [
          {
            value: '11',
            label: '克'
          },
          {
            value: '12',
            label: '千克'
          }
        ],
        weeks: [
          {
            value: '1',
            label: '星期日'
          },
          {
            value: '2',
            label: '星期一'
          },
          {
            value: '3',
            label: '星期二'
          },
          {
            value: '4',
            label: '星期三'
          },
          {
            value: '5',
            label: '星期四'
          },
          {
            value: '6',
            label: '星期五'
          },
          {
            value: '7',
            label: '星期六'
          }
        ],
        editor: {
          uploadJson:'http://127.0.0.1:9000/api/test/upload',
          height:300
        },
        area:[]
      },
      //验证规则
      formRules: {
        avatar: [{
          required: true,
          message: "请上传头像"
        }],
        name: [{
          required: true,
          message: "请输入姓名",
          trigger: 'blur'
        }, {
          min: 2,
          max: 20,
          message: "必须输入2-20个字符",
          trigger: 'blur'
        }],
        sex: [{
          required: true,
          message: "请选择性别",
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
        area: [{
          required: true,
          message: "请选择地区",
          trigger: 'blur'
        }],
        status: [{
          required: true,
          message: "请选择状态",
          trigger: 'blur'
        }],
        isSuper: [{
          required: true,
          message: "请选择",
          trigger: 'blur'
        }],
        ids: [{
          required: true,
          message: "请输入身份证号",
          trigger: 'blur'
        }, {
          min: 18,
          max: 18,
          message: "身份证为18个字符",
          trigger: 'blur'
        }],
        photos: [{
          required: true,
          message: "请上传身份证照片"
        }]
      }
    }
  },
  //实时计算属性
  computed: {
    //判断是否处于编辑状态
    isEdit() {
      this.form.id = this.$route.query.id ? this.$route.query.id : null
      return this.form.id //根据参数id判断
      //return this.$route.meta.isEdit // 根据meta判断
      //return this.$route.path.indexOf('edit') !== -1 // 根据路由判断
    }
  },
  methods: {
    ...mapActions(['getUnit', 'saveUnit']),
    getAreaData() {
      const self = this
      axios.get(self.areas).then(function(response){
        if (response.status==200) {
          const data = response.data
          self.options.area = data
        }else {
          console.log(response.status)
        }
      }).catch(function(error){
        console.log(typeof+ error)
      })
    },
    //根据id获取数据
    getItem() {
      const self = this
      self.getUnit(self.form.id).then(response => {
        //console.log('getItem==', response)
        const status = response.data.state
        const res = response.data.data
        if (status) {
          self.form = res
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
    handleBeforeUpload (file) {
      /*const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isJPG && isLt2M;*/
    },
    handleAvatarUpload(res) {
      console.log('res===', res)
      if (res.success) {
        const file = res.data[0]
        this.form.avatar = `http://127.0.0.1:9000/${file.path}`;
      }
    },
    handleIDFrontUpload(res) {
      this.form.idfront = res.files.file
    },
    handleIDBackUpload(res) {
      this.form.idback = res.files.file
    },
    handleArea(val) {
      console.log('===',val)
    },
    //表单提交
    handleSubmit() {
      const self = this
      this.$refs.postForm.validate(valid => {
        if (valid) {
          self.processing = true
          self.saveUnit(this.form).then(response => {
            //console.log('handleSubmit==', response)
            const status = response.data.success
            const res = response.data.data
            if (status) {
              self.$notify({
                title: '成功',
                message: response.data.message,
                type: 'success',
                duration: 2000
              })
            } else {
              self.$notify({
                title: '失败',
                message: response.data.message,
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
      this.$router.push('/unit1/list')
    }
  },
  beforeCreate: noop,
  created() {
    this.getAreaData()
  },
  beforeMount: noop,
  mounted(){
    if (this.isEdit) {
      this.getItem();
    }
  },
  beforeDestory: noop,
  destoryed: noop
}
</script>