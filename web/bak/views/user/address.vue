<template>
<div class="page">
  <div class="page-header">
    <header class="header">
      <div class="header-left">
        <a href="javascript:history.go(-1);" 
          class="header-btn header-btn-back">
          返回
        </a>
      </div>
      <div class="header-title">
        收货地址管理
      </div>
    </header>
  </div>

  <div class="page-main">
    <div class="page-scroller">
      <!-- 展示 -->
      <section v-if="!edit">
        <section class="address-list">
          <div class="address-item" 
            v-for="item,index in list">
            <div class="address-item-bd">
              <div class="address-item-name">
                <span>{{item.consignee}}</span>
                <span>{{item.phone}}</span>
              </div>
              <p>{{item.province}}{{item.city}}{{item.district}}{{item.address}}</p>
            </div>
            <div class="address-item-ft">
              <label class="u-check-radio-check" v-if="item.state==1">
                <input
                  class="u-check-native-control"
                  type="radio"
                  name="address"
                  checked/>
                <i class="u-check-indicator"></i>
                <span>默认地址</span>
              </label>
              <div class="address-item-opt">
                <em class="o-edit" 
                  @click="handleEdit(item)">编辑</em>
                <em class="o-del"
                  @click="handleDel(item)">删除</em>
              </div>
            </div>
          </div>
        </section>
        <div class="no-entry" v-if="list.length == 0">暂时没有收货地址</div>
        <div class="address-add"
          @click="handleAdd">新增收货地址</div>
      </section>
      
      <!-- 编辑 -->
      <section v-if="edit">
        <section class="address-form">
          <dl class="address-form-item">
            <dt>收货人</dt>
            <dd>
              <input
                class="address-form-input"
                type="text"
                v-model="form.consignee"
                placeholder="填写收货人姓名"
              />
            </dd>
          </dl>
          <dl class="address-form-item">
            <dt>手机号</dt>
            <dd>
              <input
                class="address-form-input"
                type="number"
                v-model="form.phone"
                placeholder="填写收货人手机号"
              />
            </dd>
          </dl>
          <dl class="address-form-item">
            <dt>所在地区</dt>
            <dd>
              <span class="address-form-input"
                @click="handleShow">
                <span v-if="picker.zone">{{picker.zone}}</span>
                <span v-if="!picker.zone">{{form.province}}{{form.city}}{{form.district}}</span>
              </span>
              <city-picker 
                ref="picker" 
                :selected-index="picker.index" 
                :data="picker.list"
                @select="handleSelect">
              </city-picker>
              <!-- <span class="address-form-location">定位</span> -->
            </dd>
          </dl>
          <dl class="address-form-item">
            <dt>详细地址</dt>
            <dd>
              <input
                class="address-form-input"
                type="text"
                v-model="form.address"
                placeholder="街道、楼牌号等"
              />
            </dd>
          </dl>
          <dl class="address-form-item">
            <dt style="flex:2.5;">设置默认地址</dt>
            <dd style="justify-content: flex-end;">
              <label class="u-check-switch">
                <input
                  class="u-check-native-control"
                  type="checkbox"
                  ref="checkbox"
                  @change="handleState()"
                  :checked="form.state==1 ? true : false"/>
                <i class="u-check-indicator"></i>
              </label>
            </dd>
          </dl>
          <br/>
          <p>
            提醒：每次下单时会使用该地址，实际下单地址需要您在提交订单时确认！
          </p>
        </section>
        <div class="address-btn">
          <div class="address-cancle"
            @click="edit = false">取消</div>
          <div class="address-add"
            @click="handleSave">保存</div>
        </div>
      </section>
    </div>
  </div>


</div>
</template>
<script>
import {
  mapActions
} from 'vuex'

import cityPicker from '@/components/citypicker'

export default {
  name: 'user-address',
  components: {
    cityPicker
  },
  data() {
    return {
      listQuery: {
        uid: ''
      },
      edit: false,
      list: [],
      picker: {
        list: [],
        index: [24, 0, 0],
        selected: [],
        zone: '',
      },
      form: {}
    }
  },
  computed: {
    isEdit() {
      this.listQuery.uid = this.$route.query.uid ? this.$route.query.uid : null
      return this.listQuery.uid
    }
  },
  methods: {
    ...mapActions(['getAddress', 'getCity', 'saveAddress', 'delAddress']),
    get() {
      const self = this
      self.getAddress(self.listQuery).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if (state) {
          self.list = res
        } else {
          self.$toast(message, 'error')
        }
      })
    },
    handleDel(row) {
      const self = this
      self.delAddress({id:row.id}).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if (state) {
          self.$toast(message, 'success')
        } else {
          self.$toast(message, 'error')
        }
      })
    },
    handleState() {
      const self = this
      self.form.state = self.$refs['checkbox'].checked ? 1 : 2
    },
    handleAdd() {
      const self = this
      self.edit = true
      self.form = {}
      self.picker.zone = "请选择"
      self.handleCity()
    },
    handleEdit(row) {
      const self = this
      self.edit = true
      self.form = row
      self.handleCity()
    },
    handleCity() {
      const self = this
      self.getCity({}).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if (state) {
          self.picker.list = [res.province, res.city, res.area]
        } else {
          console.log(message)
        }
      })
    },
    handleShow() {
      this.$refs['picker'].show()
    },
    handleSelect() {
      console.log('select',arguments)
      const self = this
      // index
      self.picker.index[0] = arguments[1][0] 
      self.picker.index[1] = arguments[1][1] 
      self.picker.index[2] = arguments[1][2]
      // zone
      self.picker.zone = arguments[2][0] + arguments[2][1] + arguments[2][2]
      // form
      self.picker.selected = arguments[0]
    },
    handleSave() {
      const self = this
      let form = self.form
      let selected = self.picker.selected
      let uid = self.listQuery.uid

      form.province = selected[0]
      form.city = selected[1]
      form.district = selected[2]
      
      if(!uid || !form.city || !form.phone, !form.consignee) return self.$toast('请正确输入收货信息', 'error')

      let o = {
        address: form,
        token: uid
      }

      self.saveAddress(o).then(response => {
        const res = response.data.data
        const state = response.data.state
        const message = response.data.msg
        if (state) {
          self.$toast(message, 'success')
          this.get()
        } else {
          self.$toast(message, 'error')
        }
      })
    }
  },
  created() {
    if (this.isEdit) {
      this.get()
    }
  }
}
</script>
<style type="text/css">
@import '../../styles/address.css'
</style>
