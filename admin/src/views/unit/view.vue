<template>
  <el-card class="xa-block" shadow="hover">
    <el-row :gutter="20">
      <el-col :span="12" 
        v-for="(item,index) in list"
        :key="index">
        <section class="menu-panel">
          <div class="menu-panel-hd">{{item.title}}</div>
          <ul class="menu-panel-bd">
            <li v-for="(it,i) in item.children" :key="i">
              <div class="picture">
                <img :src="it.icon"/>
              </div>
              <span class="title">{{it.name}}</span>
            </li>
          </ul>
        </section>
      </el-col>
      <!-- <el-col :span="12">
        <section class="menu-panel">
          <div class="menu-panel-hd">综合型事务支撑</div>
          <ul class="menu-panel-bd">
            <li v-for="(item,index) in list" :key="index">
              <div class="picture">
                <img :src="item.icon"/>
              </div>
              <span class="title">{{item.name}}</span>
            </li>
          </ul>
        </section>
      </el-col>
      <el-col :span="12">
        <section class="menu-panel">
          <div class="menu-panel-hd">综合型事务支撑</div>
          <ul class="menu-panel-bd">
            <li v-for="(item,index) in list" :key="index">
              <div class="picture">
                <img :src="item.icon"/>
              </div>
              <span class="title">{{item.name}}</span>
            </li>
          </ul>
        </section>
      </el-col> -->
    </el-row>  
</el-card>
</template>
<script>
export default {
  name: 'userlist',
  components: {
  },
  data() {
    return {
      list: []
    }
  },
  methods: {
    getList() {
    },
  },
  created() {
    this.list = [{
      title: '综合型事务支撑1',
      children: [{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      }]
    },{
      title: '综合型事务支撑2',
      children: [{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      }]
    },{
      title: '综合型事务支撑3',
      children: [{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      },{
        name: '我的应用',
        icon: 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2108859129,1331866552&fm=26&gp=0.jpg',
      }]
    }]
  }
}
</script>

<style scoped lang="scss">
 .menu {

    &-panel {
      margin-bottom:20px;

      &-hd {
        color:#323232;
        font-size:16px;
        font-weight:bold;
        margin-bottom:16px;
      }

      &-bd {
        display:flex;
        flex-direction:row;
        flex-wrap:wrap;
        align-items:center;
        margin-bottom:0;
      }

      li {
        margin-bottom:20px;
        margin-right:20px;
        text-align:center;

        .picture {
          width:60px;
          height:60px;
          margin:0 auto;
          overflow:hidden;
          //border-radius:999px;
        }

        .title {
          font-size:14px;
          color:#666;
        }
      }

    }
 } 
</style>
