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
            <el-radio-group 
              v-model="goods.isOnSale"
              size="small">
              <el-radio-button
                v-for="item, index in options.sales"
                :key="`sales-${item.value}`"
                :label="item.value">{{item.label}}</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="主推" prop="isFirst" >
            <el-radio-group 
              v-model="goods.isFirst"
              size="small">
              <el-radio-button
                v-for="item, index in options.firsts"
                :key="`first-${item.value}`"
                :label="item.value">{{item.label}}</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="热门" prop="isHot" >
            <el-radio-group 
              v-model="goods.isHot"
              size="small">
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

    <!-- 商品规格 -->
    <el-tab-pane label="商品规格">
      <div class="p-form__section">
        <div class="p-form__s-head">商品规格</div>
        <div class="xa-block p-form__rich">
          <el-table
            class="p-table p-table--no-hover"
            size="mini"
            :data="skuValList"
            border
            max-height="600">
            <el-table-column
              type="index"
              width="40"
              align="center">
            </el-table-column>
            <el-table-column 
              label="SKU规格" 
              prop="value" 
              header-align="center">
            </el-table-column>
            <el-table-column
              label="价格"
              width="200"
              header-align="center">
              <template slot-scope="scope">
                <el-input
                  v-model="scope.row.price"
                  size="mini"
                  type="number">
                  <span slot="prefix">&yen;</span>
                </el-input>
              </template>
            </el-table-column>
            <el-table-column
              label="数量"
              width="200"
              header-align="center">
              <template slot-scope="scope">
                <el-input
                  v-model="scope.row.stock"
                  size="mini"
                  type="number">
                  <span slot="prefix">扎</span>
                </el-input>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <el-button type="primary"
          :loading="processing"
          @click="saveSkuList">保存</el-button>
        <el-button
          @click="handleCancle">返回</el-button>
      </div>
    </el-tab-pane>
    
    <!-- 商品图片 -->
    <el-tab-pane label="商品图片">
      <div class="p-form__section">
        <div class="p-form__s-head">商品图片
          <span class="xa-text-gray">最多上传10张</span>
        </div>
        <div class="xa-block">
          <el-upload
            ref="upload"
            action=""
            list-type="picture-card"
            :file-list="galleryList"
            :on-remove="removeGallery"
            accept="image/gif, image/png, image/jpeg"
            :http-request="uploadGallery">
            <i class="el-icon-plus"></i>
          </el-upload>
        </div>
      </div> 
    </el-tab-pane>
  </el-tabs>
</template>

<script>
import {fetchGet, fetchPost} from '@/api'
import {uuid} from '@/utils'
import {URIS} from '@/config'
import kindeditor from '@/components/kindeditor'

export default {
  name: 'goods-form',
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
      },
      galleryList: [],
      skuValList: [],
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
          height:'300',
          uploadJson: URIS.Kindeditor
        },
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
    //根据id获取数据
    getItem() {
      const self = this
       fetchGet('/goods/product/item', {id: self.goods.id}).then(response => { 
        const status = response.data.state
        const res = response.data.data
        const msg = response.data.message
        if (status) {
          //gallery
          let galleryList = res.galleryList
          if(galleryList.length > 0) {
            for(const item of galleryList) {
              item.url = URIS.Goods + item.medium
            }
          }

          self.goods = res.goods
          self.galleryList = galleryList
          self.skuValList = res.skuValList
        } else {
          self.$notify({
            title: '失败',
            message: msg,
            type: 'error'
          })
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
    //sku
    saveSkuList() {
      const self = this
      self.skuValList.map(function(item) {
        return item.stock = parseInt(item.stock );
      })

      fetchPost('/goods/sku/val', self.skuValList).then(response => {
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
    },
    //图册
    removeGallery(file) {
      const self = this

      fetchPost('/goods/gallery/delete', file).then(response => {
        const status = response.data.state
        const msg = response.data.message
        if (status) {
          let list = self.galleryList
          for (let i=0,len=list.length; i<len; i++) {
            if (file.uid == list[i].uid) {
              self.galleryList.splice(i, 1)
            }
          }
        } else {
          self.$notify({
            title: '失败',
            message: msg,
            type: 'error'
          })
        } 
      })
    },
    uploadGallery(item) {
      const self = this
      let len = self.$refs.upload.uploadFiles.length
      let file = item.file
      let fileType = file.type
      let isImage = fileType.indexOf('image') != -1
      let isLt2M = file.size / 1024 / 1024 < 2;
      if (!isLt2M) {
        self.$notify({
          title: '失败',
          message: '上传图片大小不能超过2MB!',
          type: 'error'
        })
        return false;
      }
      if(!isImage){
        self.$notify({
          title: '失败',
          message: '请选择图片!',
          type: 'error'
        })
        return false;
      }
      //上传
      let formData = new FormData()
      formData.append('file', item.file)
      formData.append('gid', self.goods.id)
      formData.append('id', uuid())

      fetchPost('/goods/gallery/upload', formData).then(response => {
        const status = response.data.state
        const res = response.data.data
        const msg = response.data.message
        if (status) {
          let gallery = res
          gallery.url = URIS.Goods + gallery.medium
          self.galleryList.push(gallery)
        } else {
          self.$notify({
            title: '失败',
            message: msg,
            type: 'error'
          })
        } 
      })
      //防止跳动两次
      self.$refs.upload.uploadFiles[len-1] = null
    },
    handleCancle() {
      this.$router.push('/goods/list')
    }
  },
  created() {
    this.getCategorys();
    if (this.isEdit) {
      this.getItem();
    }
  }
}
</script>