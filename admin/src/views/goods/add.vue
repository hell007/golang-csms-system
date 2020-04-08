<template>
  <el-tabs type="border-card">
    <el-tab-pane label="基本信息">
      <el-form class="p-form"
        ref="postForm"
        :rules="formRules"
        :model="goods"
        label-suffix="："
        label-width="120px">
        <div class="p-form__section">
          <div class="p-form__s-head">基本信息</div>
          <el-form-item label="商品名称" prop="goodsName">
            <el-input class="p-form__input" 
              placeholder="请输入" 
              clearable
              v-model="goods.goodsName"
              :maxlength="formRules.goodsName[1].max" >
              <span class="p-input-count" slot="suffix">
              {{goods.goodsName.length}} / {{formRules.goodsName[1].max}}
              </span>
            </el-input>
          </el-form-item>
          <el-form-item label="商品编码" prop="goodsSn">
            <el-input
              placeholder="请输入"
              clearable
              v-model="goods.goodsSn"
              :maxlength="formRules.goodsSn[1].max"
              class="p-form__input">
              <span class="p-input-count" slot="suffix">
                {{goods.goodsSn.length}} / {{formRules.goodsSn[1].max}}
              </span>
            </el-input>
            <div class="p-form__tip">
              如果您不输入商品编码，系统将自动生成一个唯一的商品编码。
            </div>
          </el-form-item>
          <el-form-item label="卖点" prop="promoteWord">
            <el-input
              type="textarea"
              placeholder="请输入"
              v-model="goods.promoteWord"
              class="p-form__input">
            </el-input>
          </el-form-item>
          <el-form-item label="上架" prop="isOnSale" >
            <el-radio-group v-model="goods.isOnSale">
              <el-radio-button
                v-for="item, index in options.sales"
                :key="`sales-${item.value}`"
                :label="item.value">{{item.label}}</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="主推" prop="isFirst" >
            <el-radio-group v-model="goods.isFirst">
              <el-radio-button
                v-for="item, index in options.firsts"
                :key="`first-${item.value}`"
                :label="item.value">{{item.label}}</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="热门" prop="isHot" >
            <el-radio-group v-model="goods.isHot">
              <el-radio-button
                v-for="item, index in options.hots"
                :key="`hot-${item.value}`"
                :label="item.value">{{item.label}}</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="商品分类" prop="categoryId">
            <el-select 
              clearable
              v-model="goods.categoryId" 
              placeholder="请选择">
              <el-option
                v-for="item in options.categorys"
                :key="item.categoryName"
                :label="item.categoryName"
                :value="item.id">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="商品单位" prop="unit">
            <el-select 
              clearable
              v-model="goods.unit" 
              placeholder="请选择">
              <el-option
                v-for="item in options.units"
                :key="'units-'+item.value"
                :label="item.label"
                :value="item.label">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="颜色" prop="color">
            <el-input class="p-form__input" 
              placeholder="请输入" 
              clearable
              v-model="goods.color">
            </el-input>
          </el-form-item>
          <el-form-item label="关键字" prop="keywords">
            <el-input
              type="textarea"
              placeholder="请输入"
              v-model="goods.keywords"
              class="p-form__input">
            </el-input>
            </el-input>
          </el-form-item>
          <el-form-item label="描述" prop="description">
            <el-input
              type="textarea"
              placeholder="请输入"
              v-model="goods.description"
              class="p-form__input">
            </el-input>
          </el-form-item>
          <el-form-item label="图文详情" prop="contents">
            <div class="p-form__rich">
              <kindeditor
                :options="options.editor"
                v-model="goods.contents" />
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary"
              :loading="processing"
              @click="handleGoods">保存</el-button>
            <el-button
              @click="handleCancle">返回</el-button>
          </el-form-item>
        </div>
      </el-form>  
    </el-tab-pane>
  </el-tabs>
</template>

<script>
import {fetchGet, fetchPost} from '@/api'
import {URIS} from '@/config'
import kindeditor from '@/components/kindeditor' //富文本编辑器

export default {
  name: 'goods-add',
  components: {
    kindeditor
  },
  data() {
    return {
      processing: false,
      goods: {
        id: '',
        goodsName: '',
        goodsSn: '',
        categoryId: null,
        keywords: '',
        isOnSale: 1,
        isHot: 2,
        isFirst: 2,
      },
      options: {
        categorys: [],
        sales: [
          {value: 1, label: '上架'},
          {value: 2, label: '下架'}
        ],
        firsts: [
          {value: 1, label: '是'},
          {value: 2, label: '否'}
        ],
        hots: [
          {value: 1, label: '是'},
          {value: 2, label: '否'}
        ],
        units: [
          {value: 1, label: 'xx枝/扎'},
          {value: 2, label: '10枝/扎'},
          {value: 3, label: '20枝/扎'},
          {value: 4, label: '30枝/扎'}
        ],
        editor: {
          height:'500',
          uploadJson: URIS.Kindeditor
        }
      },
      //验证规则
      formRules: {
        goodsName: [{
          required: true,
          message: "请输入商品名称",
          trigger: 'blur'
        }, {
          min: 2,
          max: 50,
          message: "必须输入2-50个字符",
          trigger: 'blur'
        }],
        goodsSn: [{
          required: true,
          message: "请输入商品编码",
          trigger: 'blur'
        }, {
          min: 8,
          max: 20,
          message: "必须输入8-20个字符",
          trigger: 'blur'
        }],
        categoryId: [{
          required: true,
          message: "请选择",
          trigger: 'blur'
        }],
        isOnSale: [{
          required: true,
          message: "请选择",
          trigger: 'blur'
        }],
        isFirst: [{
          required: true,
          message: "请选择",
          trigger: 'blur'
        }],
        isHot: [{
          required: true,
          message: "请选择",
          trigger: 'blur'
        }],
        keywords: [{
          required: true,
          message: "请输入关键字",
          trigger: 'blur'
        }, {
          min: 2,
          max: 255,
          message: "必须输入2-255个字符",
          trigger: 'blur'
        }]
      }
    }
  },
  //实时计算属性
  computed: {
    //判断是否处于编辑状态
    isEdit() {
      this.goods.id = this.$route.query.id ? this.$route.query.id : null
      return this.goods.id
    }
  },
  methods: {
    getCategorys() {
      const self = this
      fetchGet('/goods/category/list', {pid:0}).then(response => {  
        const status = response.data.state
        const res = response.data.data
        const message = response.data.msg
        if (status) {
          self.options.categorys = [{id:0, categoryName:'全部'}, ...res]
        }
      }).catch(ex => {
        self.$notify({
          title: '请求错误',
          message: ex,
          type: 'error'
        })
      })
    },
    // 基本信息
    handleGoods(){
      const self = this
      this.$refs.postForm.validate(valid => {
        if (valid) {
          self.processing = true

          fetchPost('/goods/product/save', self.goods).then(response => {
            const status = response.data.state
            const res = response.data.data
            const msg = response.data.message
            if (status) {
              self.$notify({
                title: '成功',
                message: msg,
                type: 'success',
                duration: 2000
              })
            } else {
              self.$notify({
                title: '失败',
                message: msg,
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
      this.$router.push('/goods/list')
    }
  },
  created() {
    this.getCategorys();
  }
}
</script>