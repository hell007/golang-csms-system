<template>
  <el-card class="xa-block" shadow="hover">
    <workflow
      :data.sync="options"
      @ok="ok"
    />
  </el-card>
</template>

<script>
import workflow from '@/components/Workflow/ui/index'
export default {
  name: 'unit-flow',
  components: {
    workflow
  },
  data () {
    return {
      options: {}
    }
  },
  methods: {
    ok (data) {
      console.log('ok===', data)
    }
  },
  created() {
    this.options = {
      title: '员工请假审批流程',
      node: {
        name: '发起人',
        type: 'start',
        nodeId: 'sid-startevent',
        childNode: {}
      },
      components: []
      // components: [{
      //   InTableCanvas: null,
      //   componentName: "textfield",
      //   defaultLable: "姓名",
      //   defaultPrint: false,
      //   defaultProps: "请输入姓名",
      //   defaultRequired: true,
      //   idx: 0,
      //   name: "单行输入框",
      //   supports: ["label", "placeholder", "required", "print"]
      // }, {
      //   InTableCanvas: null,
      //   componentName: "tablefield",
      //   components: [{
      //     componentName: "textfield",
      //     defaultLable: "类型",
      //     defaultPrint: false,
      //     defaultProps: "请输入类型",
      //     defaultRequired: true,
      //     idx: 2,
      //     name: "单行输入框",
      //     supports: ["label", "placeholder", "required", "print"]
      //   }],
      //   defaultAction: "增加",
      //   defaultLable: "表格",
      //   defaultPrint: false,
      //   idx: 1,
      //   name: "明细/表格",
      //   selected: "0",
      //   supports: ["label", "action"]
      // }]
    }

    //console.log('test created===', this.options)
  }
}
</script>