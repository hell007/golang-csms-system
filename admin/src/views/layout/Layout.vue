<template>
  <div class="xa-main xa-frame" 
    v-cloak
    v-bind:class="{'xa-frame--collapsed': menu.collapsed}">
    
    <!-- logo -->
    <div class="xa-frame__platform">
      <div class="xa-frame__platform-icon"></div>
      <div class="xa-frame__platform-name">系统管理</div>
    </div>
    
    <!-- 头部 -->
    <div class="xa-frame__header">
      <el-button
        class="xa-sidebar-toggle"
        type="text"
        @click="menu.collapsed = !menu.collapsed">
        <i
          class="xa-icon"
          :class="{'xa-icon-menu-collapse': !menu.collapsed, 'xa-icon-menu-expand': menu.collapsed}"></i>
      </el-button>

      <ul class="xa-header-nav">
        <li class="xa-header-nav__item xa-header-nav__message">
          <a href="###">
            <el-badge is-dot v-if="message.new">
              <i class="xa-icon xa-icon-message-o"></i>
            </el-badge>
            <i v-else class="xa-icon xa-icon-message-o"></i>
          </a>
        </li>
        <li class="xa-header-nav__item xa-header-nav__qrcode">
          <el-dropdown>
            <span class="el-dropdown-link">
              <i class="xa-icon xa-icon-qrcode"></i>
            </span>
            <el-dropdown-menu slot="dropdown" class="xa-qrcode-drop">
              <el-dropdown-item>
                <img class="xa-qrcode-drop__img" :src="qrcode" />
                扫描公众号二维码访问web前台
              </el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </li>
        <li class="xa-header-nav__item xa-header-nav__user">
          <el-dropdown @command="userCommand">
            <span class="el-dropdown-link">
              <img
                class="xa-header-nav__user-avatar"
                src="../../assets/images/avatar.jpg" />
              <span>{{ user.username }}</span>
              <i class="el-icon-caret-bottom el-icon--right"></i>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item>修改密码</el-dropdown-item>
              <el-dropdown-item divided command="signout">退出</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </li>
        <li class="xa-header-nav__item xa-header-nav__switch-platform">
          <el-dropdown trigger="click">
            <el-button class="xa-header-nav__switch-platform-btn" plain>
              切换平台
              <i class="el-icon-menu"></i>
            </el-button>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item>web</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </li>
      </ul>
    </div>

    <!-- 菜单 -->
    <div class="xa-frame__sidebar xa-sidebar">
      <el-scrollbar
        class="xa-sidebar__scrollbar"
        wrap-class="xa-sidebar__wrap">
        <el-menu 
          :default-active="menu.active" 
          :collapse="menu.collapsed" 
          :unique-opened="true"
          background-color="#324157"
          text-color="#bfcbd9"
          active-text-color="#20a0ff">

          <template v-for="menu, index in permission_routers">
            <el-submenu
              :key="index"
              v-if="!menu.hidden && menu.noDropdown"
              :index="`menu-${index}`">
              <template slot="title">
                <i :class="menu.icon"></i>
                <span slot="title">{{menu.name}}</span>
              </template>
              <el-menu-item
                v-for="subMenu, subIndex in menu.children"
                :key="subIndex"
                :index="`submenu-${subIndex}`">
                <router-link
                  class="xa-sidebar__link"
                  :to="`${subMenu.path}`" >
                  <i :class="subMenu.icon"></i>
                  {{subMenu.name}}
                </router-link>
              </el-menu-item>
            </el-submenu>
            
            <!-- 首页 -->
            <el-menu-item v-if="!menu.noDropdown && !menu.hidden" 
              :key="index"
              :index="`menu-${index}`">
              <router-link
                class="xa-sidebar__link"
                :to="menu.path">
                <i :class="menu.icon"></i>
                <span>{{menu.name}}</span>
              </router-link>
              <span class="xa-sidebar__title" slot="title">
                {{menu.name}}
              </span>
            </el-menu-item>
          </template>

        </el-menu>
      </el-scrollbar>
    </div>
    
    <!-- 主内容 -->
    <div class="xa-frame__main" element-loading-text="加载中">
      <div class="xa-content">
        <!-- 面包屑 -->
        <Levelbar></Levelbar>
        <!-- vue -->
        <transition name="fade" mode="out-in">
          <keep-alive>
            <router-view v-if="$route.meta.keepAlive" :key="`page-${key}`">
            </router-view>
          </keep-alive>
        </transition>
        <transition name="fade" mode="out-in">
          <router-view v-if="!$route.meta.keepAlive" :key="`page-${key}`">
          </router-view>
        </transition>
      </div>
    </div>
    
  </div>
</template>

<script>
import {
  mapState,
  mapActions,
  mapGetters
} from 'vuex'

import { Levelbar } from './components' //./index.js

import  qrcode from '@/assets/images/qrcode.png'

export default {
  name: 'layout',
  components: { 
    Levelbar
  },
  data() {
    return {
      menus: [],
      menu: {
        active: '0',
        collapsed: window.innerWidth < 1440
      },
      message: {
        new: true
      },
      platformDialogOpen: false,
      qrcode: qrcode
    }
  },
  methods: {
    ...mapActions(['LoginOut']),
    userCommand(command) {
      const self = this
      if (command === 'signout') {
        self.handleSignout()
      }
    },
    handleSignout() {
      const self = this 
      self.LoginOut().then(response => {
        const status = response.data.state
        const message = response.data.msg
        if (status) {
          this.$router.push({
            path: '/login'
          })
        }
      }).catch(e => {
        self.$notify({
          title: '失败',
          message: message,
          type: 'error'
        })
      })
    }
  },
  //实时计算属性
  computed: {
    //...mapState(['user'])
    ...mapGetters([
      'permission_routers',
      'user'
    ]),
    key() {
      return this.$route.name !== undefined ? this.$route.name + +new Date() : this.$route + +new Date()
    }
  },
  created: function() {
    //console.log('mapGetters===', this.user)
    //console.log('mapGetters', this.permission_routers)
  }
}
</script>
