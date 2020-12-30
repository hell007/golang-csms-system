<template>
  <el-card class="xa-block" shadow="hover">
    <!-- 数据 -->
    <div
      class="xa-block"
      v-loading="loading"
      element-loading-text="正在玩命加载中..."
    >
      <el-table
        class="p-table"
        ref="listTable"
        row-key="id"
        :key="tableKey"
        :data="list"
        lazy
        fit
        :load="onLoad"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
      >
        <el-table-column prop="name" label="组织名称" align="left">
        </el-table-column>
        <el-table-column
          fixed="right"
          label="操作"
          width="500"
          align="center"
          header-align="center"
        >
          <template slot-scope="scope">
            <el-button
              type="text"
              icon="el-icon-folder-add"
              class="organization-btn"
              @click="addOrganization(scope.row)"
              >新增</el-button
            >
            <el-button
              type="text"
              icon="el-icon-edit-outline"
              class="organization-btn organization-btn-primary"
              @click="editOrganization(scope.row)"
              >修改</el-button
            >
            <el-button
              type="text"
              icon="el-icon-delete"
              class="organization-btn"
              @click="deleteOrganization(scope.row)"
              >删除</el-button
            >
            <el-button
              type="text"
              icon="el-icon-paperclip"
              class="organization-btn organization-btn-primary"
              @click="bindOrganization(scope.row)"
              >绑定角色</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 新增、修改组织 -->
    <el-dialog
      :show-close="false"
      :title="organizationTitle"
      :visible.sync="dialog.organization"
      center
      width="40%"
    >
      <el-form
        class="p-form"
        :model="organizationForm"
        :rules="organizationRule"
        ref="organizationForm"
        label-suffix="："
        label-width="120px"
      >
        <el-form-item label="上级组织名称">
          <el-input
            v-model="organizationForm.parentName"
            class="p-form__input"
            readonly
            disabled
            style="width:300px"
          ></el-input>
        </el-form-item>
        <el-form-item label="组织名称" prop="name">
          <el-input
            v-model="organizationForm.name"
            class="p-form__input"
            placeholder="请输入组织名称"
            style="width:300px"
          ></el-input>
        </el-form-item>
      </el-form>
      <div class="permissionDialog-btn">
        <el-button @click="organizationFormClose">取 消</el-button>
        <el-button type="primary" @click="organizationFormSubmit" :disabled="addUpdateBtnFlg"
          >确 定</el-button
        >
      </div>
    </el-dialog>

    <!-- 绑定角色 -->
    <el-dialog
      :show-close="false"
      title="绑定角色"
      class="dialog-box permissionDialog-box"
      :visible.sync="dialog.role"
      width="30%"
      center
    >
      <div class="permissionDialog-explain">
        <div class="permissionDialog-explain-name">
          当前组织：<span>{{ currentOrganization }}</span>
        </div>
        <div class="permissionDialog-explain-text">
          当前组织角色配置如下，可按需修改
        </div>
      </div>
      <div class="permissionDialog-tree">
        <el-tree
          :data="role.list"
          show-checkbox
          node-key="id"
          ref="roleTree"
          highlight-current
          default-expand-all
          :default-checked-keys="role.check"
          :props="role.props"
          :empty-text="tips"
        >
        </el-tree>
      </div>
      <div class="permissionDialog-btn">
        <el-button @click="roleClose()">取 消</el-button>
        <el-button type="primary" @click="roleSubmit()" :disabled="bindRoleFlag">确 定</el-button>
      </div>
    </el-dialog>
  </el-card>
</template>

<script>
import { fetchGet, fetchPost } from "@/api";
export default {
  name: "organization-list",
  components: {},
  data() {
    return {
      loading: false,
      list: [],
      tableKey: 0,
      currentOrganization: "",//绑定角色
      currentCenterId: "",//绑定角色的centerId
      role: {
        list: [],
        check: [],
        props: {
          id: "id",
          label: "name",
          children: "children"
        }
      },
      dialog: {
        role: false,
        organization: false
      },
      //新增、修改
      organizationTitle: "",
      organizationForm: {
        currentCenter: {},
        parentName: "",//父组织名
        name: "",
        type: "",//组织新增还是修改
      },
      organizationRule: {
        name: [{ required: true, message: "请选输入组织名称", trigger: "blur" }]
      },
      parentMap: new Map(),
      tips: "",
      addUpdateBtnFlg: false,//新增修改确认按钮可点击标志
      bindRoleFlag: false,//绑定角色确认按钮可点击标志
      rolesIds: [],//能选择的roleIds
    };
  },
  methods: {
    //局部刷新数据
    reLoad(parentNodeId){
      let parentNode = this.parentMap.get(parentNodeId);
      if(parentNode) {
        this.$set(this.$refs.listTable.store.states.lazyTreeNodeMap, parentNodeId, []);
        this.onLoad(parentNode.tree, parentNode.treeNode, parentNode.resolve);
      }
    },
    onLoad(tree, treeNode, resolve) {
      this.parentMap.set(tree.id, {tree,treeNode,resolve});
      let self = this;
      let data = {};
      data.parentId = tree.id;
      fetchPost('/flep/web/api/center/queryChildCenter', data).then(function (res) {
        var data=res.data;
        if(data.status==200){
          let childCenterList = [];
          if(data.body != null && data.body.length > 0){
            data.body.forEach((item, index) => {
              item.id = item.centerId;
              item.name = item.centerName;
              item.hasChildren = true;
              childCenterList.push(item);
            });
          }
          resolve(childCenterList);
        }else{
          self.$message.error("查询下一级组织失败");
        }
      }).catch(() => {
        self.$message({type: 'error', message: '查询下一级组织网络异常'});
      });
    },
    //新增、修改
    addOrganization(row) {
      this.organizationForm.currentCenter = row;
      this.organizationForm.parentName = row.centerName;//新增时，当前选择行作为父组织
      this.organizationForm.name = "";
      this.organizationForm.type = "ADD";
      this.organizationTitle = "新增组织机构";
      this.dialog.organization = true;
    },
    editOrganization(row) {
      this.organizationForm.currentCenter = row;
      this.organizationForm.parentName = "";//先把父组织名清除，后面调接口查询父组织信息
      this.organizationForm.name = row.centerName;//编辑先反显当前行的名
      this.organizationForm.type = "UPDATE";
      this.organizationTitle = "编辑组织机构";
      this.dialog.organization = true;
      //查询父组织名
      this.queryCenter();
    },
    organizationFormClose() {
      this.dialog.organization = false;
    },
    //查询组织信息
    queryCenter(){
      let data = {};
      data.centerId = this.organizationForm.currentCenter.parentCenterId;
      let self = this;
      fetchPost('/flep/web/api/center/queryCenterInfo', data).then(function (res) {
        var data=res.data;
        if(data.status==200){
          if(data.body != null && data.body.length > 0){
            self.organizationForm.parentName = data.body[0].centerName;
          }
        }else{
          self.$message.error("查询父组织信息失败！");
        }
      }).catch(() => {
        self.$message({type: 'error', message: '查询父组织信息网络异常!'});
      });
    },
    //增加组织信息
    addCenter(){
      let data = {};
      data.centerName = this.organizationForm.name.trim();//新增的组织名
      data.provinceId = this.organizationForm.currentCenter.provinceId;//省id
      data.cityId = this.organizationForm.currentCenter.cityId;//地市id
      data.countyId = this.organizationForm.currentCenter.countyId;//区县id
      data.gridId = this.organizationForm.currentCenter.gridId;//网格id
      data.areaId = this.organizationForm.currentCenter.areaId;//区域id
      data.parentCenterId = this.organizationForm.currentCenter.centerId;//新增的组织的父id
      let self = this;
      self.addUpdateBtnFlg = true;
      fetchPost('/flep/web/api/center/addCenter', data).then(function (res) {
        var data=res.data;
        if(data.status==200){
          self.$message.success("新增组织信息成功！");
          self.reLoad(self.organizationForm.currentCenter.centerId);
        }else{
          self.$message.error(data.body);
        }
        self.addUpdateBtnFlg = false;
        self.dialog.organization = false;
      }).catch(() => {
        self.addUpdateBtnFlg = false;
        self.dialog.organization = false;
        self.$message({type: 'error', message: '新增组织信息网络异常!'});
      });
    },
    //更新组织信息
    updateCenter(){
      let data = {};
      data.centerId = this.organizationForm.currentCenter.centerId;
      data.centerName = this.organizationForm.name.trim();
      data.parentCenterId = this.organizationForm.currentCenter.parentCenterId;//修改的组织的父id，只做校验使用
      let self = this;
      self.addUpdateBtnFlg = true;
      fetchPost('/flep/web/api/center/updateCenter', data).then(function (res) {
        var data=res.data;
        if(data.status==200){
          self.$message.success("更新组织信息成功！");
          self.reLoad(self.organizationForm.currentCenter.parentCenterId);
        }else{
          self.$message.error(data.body);
        }
        self.dialog.organization = false;
        self.addUpdateBtnFlg = false;
      }).catch(() => {
        self.dialog.organization = false;
        self.addUpdateBtnFlg = false;
        self.$message({type: 'error', message: '更新组织信息网络异常!'});
      });
    },
    //校验当前行是否能删除
    checkCanBeDelete(row){
      let data = {};
      data.centerId = row.centerId;
      let self = this;
      fetchPost('/flep/web/api/center/checkCanBeDelete', data).then(function (res) {
        var data=res.data;
        if(data.status==200){
          self.$confirm(`确认删除组织"${row.name}"`, "提示", {
            confirmButtonText: "确定",
            cancelButtonText: "取消",
            type: "warning"
          }).then(() => {
            //删除
            self.deleteCenter(row);
          });
        }else{
          self.$message.error(data.body);
        }
      }).catch(() => {
        self.$message({type: 'error', message: '校验当前组织能否删除网络异常!'});
      });
    },
    //逻辑删除组织
    deleteCenter(row){
      let data = {};
      data.centerId = row.centerId;
      let self = this;
      fetchPost('/flep/web/api/center/deleteCenter', data).then(function (res) {
        var data=res.data;
        if(data.status==200){
          self.$message.success("组织"+row.name+"删除成功！");
          self.reLoad(row.parentCenterId);
        }else{
          self.$message.error("组织"+row.name+"删除失败！");
        }
      }).catch(() => {
        self.$message({type: 'error', message: '组织删除网络异常!'});
      });
    },
    queryRoles(row) {
      //获取当前级别角色
      let provinceId = row.provinceId;//省id
      let cityId = row.cityId;//地市id
      let countyId = row.countyId;//区县id
      let gridId = row.gridId;//网格id
      let roleLevel = 0;
      let roleTitle = "";
      if(provinceId != "" && provinceId != null && provinceId != undefined){
        roleLevel = 1;
        roleTitle = "省级角色";
      }
      if(cityId != "" && cityId != null && cityId != undefined){
        roleLevel = 2;
        roleTitle = "市级角色";
      }
      if(countyId != "" && countyId != null && countyId != undefined){
        roleLevel = 3;
        roleTitle = "区县角色";
      }
      if(gridId != "" && gridId != null && gridId != undefined){
        roleLevel = 4;
        roleTitle = "网格角色";
      }

      this.role.list = [];
      this.role.check = [];
      this.tips = "";
      let self = this;
      let data = {};
      data.roleLevel = roleLevel;
      data.centerId = row.centerId;
      fetchPost('/flep/web/api/center/queryRoles', data).then(function (res) {
        var data=res.data;
        if(data.status==200){
          let roleList = data.body.roleList;//当前组织能绑定的角色
          let bindRoleList = data.body.bindRoleList;//当前组织已绑定的角色

          if(roleList != null && roleList.length > 0){
            let children = [];
            let rolesIds = [];//角色id组成的数组，为了接下来判定组织已绑定的角色在能绑定的角色里
            roleList.forEach((item, index) => {
              item.id = item.roleId;
              item.name = item.roleName;
              children.push(item);
              rolesIds.push(item.roleId);
            });

            self.rolesIds = rolesIds;//把当前的能选择的rolesIds放在外面，选择角色时使用到

            //能绑定的角色展示
            self.role.list = [{id: 1, name: roleTitle, children: children}];

            //将已选择的角色进行勾选
            if(bindRoleList != null && bindRoleList.length > 0){
              bindRoleList.forEach((item, index) => {
                let roleId = item.roleId;
                //要勾选的角色id在能绑定的角色id里面，则勾选
                if(rolesIds.indexOf(roleId) != -1){
                  self.role.check.push(roleId);
                }
              });
            }
          }else{
            self.tips = "暂无数据！";
          }
        }else{
          self.$message.error("查询角色失败");
        }
      }).catch(() => {
        self.$message({type: 'error', message: '查询角色网络异常'});
      });
    },
    organizationFormSubmit() {
      if(this.organizationForm.name.trim() == ""){
        this.$message.info("组织名称必输！");
        return;
      }

      if(this.organizationForm.type == "ADD"){
        //新增
        this.addCenter();
      }else if(this.organizationForm.type == "UPDATE"){
        //修改
        this.updateCenter();
      }
    },
    //删除
    deleteOrganization(row) {
      this.checkCanBeDelete(row);
    },
    //绑定
    bindOrganization(row) {
      this.currentOrganization = row.name;
      this.currentCenterId = row.centerId;
      this.queryRoles(row);
      this.dialog.role = true;
    },
    roleClose() {
      this.dialog.role = false;
    },
    roleSubmit() {
      let self = this;
      let data = {};

      let roleList = [];//角色ids
      let checkedKeys = self.$refs.roleTree.getCheckedKeys();//选中的角色
      //目的是去除一级选中的1
      checkedKeys.forEach((item, index) => {
        if(self.rolesIds.indexOf(item) != -1){
          roleList.push(item);
        }
      });

      data.roleList = roleList;
      data.centerId = this.currentCenterId;
      self.bindRoleFlag = true;
      fetchPost('/flep/web/api/center/centerBindRoles', data).then(function (res) {
        var data=res.data;
        if(data.status==200){
          self.$message.success("绑定角色成功");
          self.dialog.role = false;
        }else{
          self.$message.error("绑定角色失败");
        }
        self.bindRoleFlag = false;
      }).catch(() => {
        self.$message({type: 'error', message: '绑定角色网络异常'});
        self.bindRoleFlag = false;
      });
    },
    //查询当前登录人组织信息
    queryLoginUserCenter(){
      let data = {};
      let self = this;
      fetchPost('/flep/web/api/center/queryCenterInfo', data).then(function (res) {
        var data=res.data;
        if(data.status==200){
          if(data.body != null && data.body.length > 0){
            let startNode = data.body[0];
            startNode.id = startNode.centerId;
            startNode.name = startNode.centerName;
            startNode.hasChildren = true;
            self.list = [startNode];
          }else{
            self.$message.error("当前登录用户组织信息为空！");
          }
        }else{
          self.$message.error("查询当前用户织信息失败！");
        }
      }).catch(() => {
        self.$message({type: 'error', message: '查询当前用户织信息网络异常!'});
      });
    },
  },
  created() {
    this.queryLoginUserCenter();
  }
};
</script>
<style scoped lang="scss">
@import "../../styles/_global.scss";

/deep/ .dialog-box > div:first-child {
  border-radius: 8px;
  overflow: hidden;
}

/deep/ .organization-btn {
  color: $color-32;
  margin-right: 50px;
  font-size: 14px;
  display: flex;
  align-items: center;
  float: left;

  i {
    font-size: 22px;
  }

  &-primary {
    i {
      color: $color-primary;
    }
  }

  &:last-child {
    margin-right: 0;
  }
}

/deep/ .permissionDialog {
  &-tree {
    width: 100%;
    max-height: 440px;
    overflow-y: scroll;
    padding: 0 25px 12px;
    margin-top: 16px;
  }

  &-box {
    .el-dialog__body {
      padding: 0 0 30px 0;
    }
  }

  &-tabs {
    display: flex;

    li {
      width: 50%;
      border-bottom: 1px solid $color-border-gray;
      background-color: #fff;
      color: $color-primary;
      font-size: 13px;
      font-weight: bold;
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 8px 0;
      cursor: pointer;
    }

    &-active {
      color: #fff !important;
      background-color: $color-primary !important;
    }
  }

  &-explain {
    padding: 10px 25px;
    border-bottom: 1px solid $color-border-gray;
    width: 100%;
    text-align:left;

    &-text {
      font-size: 12px;
      color: $color-96;
      margin-top: 4px;
    }

    &-name {
      font-size: 13px;
      color: $color-32;

      span {
        color: $color-primary;
      }
    }
  }

  &-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 50px;

    button {
      width: 150px;

      &:first-child {
        color: #f39800;
        border-color: #f39800;
      }
    }
  }
}
</style>
