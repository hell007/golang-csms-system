<template>
  <section>
    <!-- 基本信息 -->
    <el-card class="xa-block" shadow="hover">
      <el-form
        class="p-form"
        ref="postForm"
        label-suffix="："
        label-width="120px">

        <div class="p-form__section">
          <div class="p-form__s-head">
            <span>基本信息</span>
            <span class="p-form__s-icon"
              @click="handleToggle(cards[0])">
              <i :class="cards[0].active ? 'el-icon-arrow-up' : 'el-icon-arrow-down'"></i>
            </span>
          </div>
          <div class="p-form__s-body"
            v-if="cards[0].active">
            <el-form-item label="百科标题">
              这是一个百科的标题题目
            </el-form-item>

            <el-form-item label="关键字">
              百科1、百科2
            </el-form-item>

            <el-row :gutter="20">
              <el-col :span="8">
                <el-form-item label="有效期">
                  长期有效
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="创建日期">
                  2020-05-06 13:38:09
                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="8">
                <el-form-item label="创建人">
                  xwSysAdmin
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="组织角色">
                  云南公司 超级管理员
                </el-form-item>
              </el-col>
            </el-row>

            <el-form-item label="下载展示">
              <el-row :gutter="20">
                <el-col :span="4">
                  问题汇总.xlsx
                </el-col>
                <el-col :span="4">
                  上传时间：2020-04-28
                </el-col>
                <el-col :span="4">
                  <a class="p-form__link" 
                    href="#" 
                    download="问题汇总.xlsx">点击下载</a>
                </el-col>
              </el-row>
            </el-form-item>
            
            <el-form-item label="发布情况展示">
              <el-row :gutter="20">
                <el-col :span="4">
                  xwSysAdmin
                </el-col>
                <el-col :span="4">
                  创建百科
                </el-col>
                <el-col :span="4">
                  2020-05-06 13:55:09
                </el-col>
              </el-row>
              <div class="p-form__tip" style="color:red;">:span=24 宽度可以通过修改数字值来改变</div>
            </el-form-item>

            <el-form-item label="文字描述">
              <div class="p-form__rich">
                这是一个百科的文字描述，这是一个百科的文字描述。这是一个百科的文字描述，这是一个百科的文字描述。这是一个百科的文字描述，这是一个百科的文字描述。
              </div>
            </el-form-item>

            <el-form-item label="图文描述">
              <div class="p-form__rich">
                <kindeditor
                  :options="options.editor"
                  v-model="form.content" />
              </div>
            </el-form-item>
          </div>
        </div>
      </el-form>
    </el-card>

    <!--  图片展示 -->
    <el-card class="xa-block" shadow="hover">
      <el-form
        class="p-form"
        ref="postForm"
        label-suffix="："
        label-width="120px">
        
        <div class="p-form__section">
          <div class="p-form__s-head">
            <span>图片展示</span>
            <span class="p-form__s-icon"
              @click="handleToggle(cards[1])">
              <i :class="cards[1].active ? 'el-icon-arrow-up' : 'el-icon-arrow-down'"></i>
            </span>
          </div>
          <div class="p-form__s-body"
            v-if="cards[1].active">
            <el-form-item label="头像">
              <el-upload
                class="p-uploader"
                style="width:80px;height:80px"
                :action="''"
                :disabled="true"
                :show-file-list="false">
                <img class="p-uploader__img" src="../../assets/project/images/avatar.jpg">
              </el-upload>
            </el-form-item>

            <el-form-item label="身份证照片">
              <div class="p-uploader-inline">
                <el-upload
                  class="p-uploader"
                  style="width: 300px; height: 200px;"
                  :action="''"
                  :disabled="true"
                  :show-file-list="false">
                  <img class="p-uploader__img" src="../../assets/project/images/id-1.jpeg">
                </el-upload>
                <div>正面</div>
              </div>
              <div class="p-uploader-inline">
                <el-upload
                  class="p-uploader"
                  style="width: 300px; height: 200px;"
                  :action="''"
                  :disabled="true"
                  :show-file-list="false">
                  <img class="p-uploader__img" src="../../assets/project/images/id-2.jpeg">
                </el-upload>
                <div>背面</div>
              </div>
            </el-form-item>

            <el-form-item label="图片上传">
              <div class="p-form__avatar"
                style="width:80px;height:80px">
                <img src="../../assets/project/images/avatar.jpg">
              </div>
            </el-form-item>

            <el-form-item label="图片相册">
              <ul class="p-form__pic">
                <li class="p-form__pic-wrap"
                  style="width:120px;height:120px"
                  v-for="item,index in form.picList"
                  :key="index">
                  <img :src="item.url">
                </li>
              </ul>
            </el-form-item>
          </div>    
        </div>
      </el-form>
    </el-card>

    <el-button
      type="info"
      round
      @click="handleCancle">返回</el-button>
  </section>
</template>

<script>

import kindeditor from '@/components/kindeditor'
export default {
  name: 'unit-detail',
  components: {
    kindeditor
  },
  data() {
    return {
      // 卡片
      cards: [{
        name: '基本信息',
        active: true
      }, {
        name: '图片展示',
        active: true
      }],
      form: {
        content: '',
        picList: [{
          name: 'food.jpg', 
          url: 'http://183.224.28.69:9080/csmp/static/upload/images/d57c26262d0a4229a4e7b2d9fa4cfb1c.jpg'
        }, {
          name: 'food.jpg', 
          url: 'http://183.224.28.69:9080/csmp/static/upload/images/d57c26262d0a4229a4e7b2d9fa4cfb1c.jpg'
        }],
      },
      options: {
        editor: {
          height:300,
          readonlyMode: true
        }
      },
    }
  },
  methods: {
    handleToggle(card) {
      card.active = !card.active
    },
    handleGo() {
      
    },
    handleCancle() {
      this.$router.push('/unit1/list')
    }
  },
  created() {
    
  },
  mounted() {
    this.form.content = '<p>这是一个百科的文字描述，这是一个百科的文字描述。这是一个百科的文字描述，这是一个百科的文字描述。这是一个百科的文字描述，这是一个百科的文字描述。</p><p><img style="width:200px;" src="http://183.224.28.69:9080/csmp/static/upload/images/d57c26262d0a4229a4e7b2d9fa4cfb1c.jpg" /></p>'
  }
}
</script>
