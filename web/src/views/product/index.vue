<template>
  <div class="page">
    <div class="page-header">
      <header class="header">
        <div class="header-left">
          <a href="javascript:history.go(-1);" 
            class="header-btn header-btn-back">返回</a>
        </div>
        <div class="header-title">商品详情</div>
        <div class="header-right">
          <a href="###" 
            class="header-btn header-btn-share">分享</a>
        </div>
      </header>
    </div>
    <div class="page-main">
      <div class="page-scroller">
        <!--业务-->
        <div class="js-panel-1">
          <!-- gallery -->
          <section class="product-gallery swiper-container">
            <div class="swiper-wrapper">
              <div class="swiper-slide" 
                v-for="item,index in gallerys">
                <img :src="link + item.source" />
              </div>
            </div>
            <!-- Add Pagination -->
            <div class="swiper-pagination"></div>
          </section>
          <section class="product-name">
            <div>
              <span class="product-name-title">
                {{goods.goodsName}}
              </span>
              <span class="product-name-price">
                ￥{{goods.goodsPrice}}
              </span>
            </div>
            <p class="product-name-desc">
              {{goods.promoteWord}}
            </p>
          </section>
          <!--规格-->
          <section class="product-item">
            <div class="product-item-key">已&nbsp;&nbsp;&nbsp;&nbsp;选</div>
            <div class="product-item-val arrow">
              <div class="v-panel">
                <span v-if="!checked">请选择</span>
                <div v-if="checked" class="v-msg">
                  <span v-for="item,index in form.list">{{item.value}}</span>
                </div>
              </div>
            </div>
          </section>
          <!--提醒-->
          <section class="product-item">
            <div class="product-item-key">限购提醒</div>
            <div class="product-item-val">
              <div class="v-panel">
                <div class="v-msg">
                  该商家门店提供业务仅限归属地为昆明市移动用户办理
                </div>
              </div>
            </div>
          </section>
          <!--送至-->
          <section class="product-item">
            <div class="product-item-key">
              送&nbsp;&nbsp;&nbsp;&nbsp;至
            </div>
            <div class="product-item-val arrow">
              <div class="v-panel">
                <div class="v-address">
                  云南省昆明市官渡区广福路2号云南移动...
                </div>
              </div>
            </div>
          </section>
        </div>

        <!--详情-->
        <div class="product-detail">
          <ul class="tab product-detail-tab">
            <li :class="`tab-item ${tabs.index==0 ? 'tab-item-active' : '' }`"
              @click="handleTabs(0)">
              <span class="tab-inner">商品介绍</span>
            </li>
            <li :class="`tab-item ${tabs.index==1 ? 'tab-item-active' : '' }`"
              @click="handleTabs(1)">
              <span class="tab-inner">售后说明</span>
            </li>
          </ul>
          <div class="panel product-detail-panel" v-if="tabs.index==0">
            {{goods.contents | html2Text}}
          </div>
          <div class="panel product-detail-panel" v-if="tabs.index==1">
            {{goods.description}}
          </div>
        </div>
      </div>
    </div>
    <div class="page-footer">
      <div class="bar product-bar">
        <div class="bar-main">
          <div class="btn-block btn-opt">
            <span class="colection on">收藏</span>
            <a class="tell" hre="tel:10086">电话</a>
          </div>
          <button class="btn btn-block btn-orange" 
            type="button" 
            @click="handleAction(2)">
            加入订购单
          </button>
          <button class="btn btn-block btn-primary" 
            type="button" 
            @click="handleAction(1)">
            立即订购
          </button>
        </div>
      </div>
    </div>

    <!--规格-->
    <Dialog
      id="dialog-sku"
      :visible="dialog.sku.visible"
      @submit="handleSubmit"
      @cancle="handleCancle">
      <div class=" product-specs">
        <div class="product-picture">
          <div class="p-pic">
            <img :src="link + picture" />
          </div>
          <div class="p-text">
            <p>商品编号：{{goods.goodsSn}}</p>
            <p>规格：{{goods.unit}}</p>
            <p>颜色：{{goods.color}}</p>
          </div>
        </div>
        <!--级别-->
        <div class="product-params" v-for="item,index in skus">
          <p>{{item.name}}</p>
          <div class="product-params-list" 
            v-for="it,i in item.list">
            <div class="product-params-key">
              <span class="k-name">{{it.value}}</span>
              <span class="k-price">￥{{it.price}}</span>
            </div>
            <div class="product-params-val">
              <div class="u-number-input u-number-input-inited">
                <input 
                  type="number" 
                  min="1" 
                  :max="it.stock" 
                  step="1" 
                  pattern="\d*" 
                  :value="it.num" 
                  readonly 
                  class="u-number-input-input">
                <button 
                  @click="handleNumber('minus',i, it)" 
                  :disabled="it.disabledMin" 
                  class="u-number-input-button u-number-input-button-decrease" 
                  type="button">-</button>
                <button 
                  @click="handleNumber('plus', i, it)" 
                  :disabled="it.disabledMax" 
                  class="u-number-input-button u-number-input-button-increase" 
                  type="button">+</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Dialog>
    
    <!-- 地址 -->
    <Dialog
      id="dialog-addr"
      :visible="dialog.addr.visible">
      <div class="product-address">
        <div class="product-address-list">
          <div class="product-address-item">
            <label class="u-check-radio-check">
              <input class="u-check-native-control" type="radio" name="address" checked />
              <i class="u-check-indicator"></i>
              <span>云南省昆明市广福路云南移动有限公司</span>
            </label>
          </div>
          <div class="product-address-item">
            <label class="u-check-radio-check">
              <input class="u-check-native-control" 
                type="radio" name="address" />
              <i class="u-check-indicator"></i>
              <span>
                云南省昆明市广福路云南移动有限公司销售分公司是鸣笛致
              </span>
            </label>
          </div>
        </div>
        <div class="product-address-select">选择其他地址</div>
      </div>
    </Dialog>

  </div>
</template>
<script>
import {
  mapActions
} from 'vuex'

import {
  URIS
} from '@/api/config'

import {
  param2Obj,
  html2Text
} from '@/utils'

import Dialog from '@/components/dialog/index'

import Swiper from 'swiper'
import 'swiper/css/swiper.min.css'

import * as storage from '@/utils/storage'
import {
  CSMSKEY
} from '@/api/config'

export default {
  name: 'product',
  components: {
    Dialog
  },
  data() {
    return {
      link: URIS.Goods,
      checked: false,
      goods: {},
      gallerys: [],
      skus: [],
      picture: "",
      action: null,
      form: {
        code: '',
        token: storage.get(CSMSKEY.Token),
        order: {},
        goods: null
      },
      dialog: {
        sku: {
          visible: false
        },
        addr: {
          visible: false
        }
      },
      toast: {
        visible: false,
        icon: '',
        text: ''
      },
      tabs: {
        index: 0
      }
    }
  },
  methods: {
    ...mapActions(['getGoods', 'postGoods']),
    get() {
      const self = this
      let query = param2Obj()
      self.getGoods(query).then(response => {
        const state = response.data.state
        const res = response.data.data
        const msg = response.data.msg
        if (state) {
          self.goods = res.goods
          self.gallerys = res.gallery
          self.skus = res.sku

          // 组装、处理
          //self.goods.contents = html2Text(self.goods.contents)

          if(self.gallerys.length > 0) {
            self.picture = self.gallerys[0].small
          }

          if(self.skus.length > 0) {
            for (const item of self.skus[0].list) {
              item.num = 0
              item.disabledMin = true
              item.disabledMax = false
            }
          }
        } else {
          self.$toast(msg,'error')
        }
      });
    },
    handleNumber(opt, index, item) {
      const self = this
      if (opt == 'plus') {
        if (item.num >= item.stock) {
          item.disabledMax = true
        } else {
          item.disabledMin = false
          item.num = item.num + 1
        }

      } else if (opt == 'minus') {
        if (item.num == 0) {
          item.disabledMin = true
        } else {
          item.disabledMax = false
          item.num = item.num - 1
        }
      }
      self.$set(self.skus[0].list, index, item)
    },
    handleAction(action) {
      const self = this
      self.action = action || 2
      self.dialog.sku.visible = true
    },
    handleCancle() {
      const self = this
      self.dialog.sku.visible = false
      self.action = null
      self.form = {}
      let skus = self.skus[0].list
      for (let i = 0, len = skus.length; i < len; i++) {
        skus[i].num = 0
        skus[i].disabledMin = true
        self.$set(self.skus[0].list, i, skus[i])
      }
    },
    handleSubmit() {
      const self = this

      if (!self.form.token) {
        self.$router.push('/login')
        return
      }

      self.form.code = self.action
      self.form.goods = []

      let skus = self.skus[0].list
      for (const item of skus) {
        if (item.num > 0) {
          let temp = {
            goodsId: self.goods.id,
            goodsSku: item.value,
            goodsPrice: item.price,
            goodsNum: item.num
          }
          self.form.goods.push(temp)
        }
      }

      self.checked = true
      if (self.form.goods.length == 0) {
        self.$toast('请选择鲜花规格！','error')
        return
      }
      // if (!self.form.city) {
      //     self.$toast('请选择邮寄地址','error')
      //     return
      // }
      
      self.form.order = {
        consignee: '笑笑',
        phone: '18725100616',
        province: '530000',
        city: '530100',
        district: '530113',
        address: '南亚风情园',
      }

      self.postGoods(self.form).then(response => {
        const state = response.data.state
        const res = response.data.data
        const msg = response.data.msg
        if (state) {
          self.dialog.sku.visible = false
          self.$toast(msg,'success')
        } else {
          self.$toast(msg,'error')
        }
      });
    },
    handleTabs(index) {
      this.tabs.index = index
    }
  },
  created() {
    this.get()
  },
  mounted() {
    // swiper
    window.onload = function() {
      setTimeout(function(){
        var myswiper = new Swiper('.swiper-container', {
          pagination: {
            el: '.swiper-pagination',
            type: 'fraction',
          }
        })
      },200)
    }
  },
}
</script>
<style type="text/css">
@import '../../styles/product.css'
</style>
