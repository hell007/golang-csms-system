<template>
	<section class="wf-form">
		<el-form
			class="wf-form-pannel"
      ref="postForm"
      :rules="formRules"
      :model="options"
      label-position="top">
      <el-form-item class="wf-field" prop="title">
      	<div class="fieldname" name="label">
					<b>*</b>
					<span>审批名称</span>
					<span class="fieldinfo">最多50个字</span>
				</div>
        <el-input
        	class="fieldval"
        	size="small"
          placeholder="请输入"
          clearable
          v-model="form.title"
          :maxlength="formRules.title[1].max" >
        </el-input>
      </el-form-item>

      <el-form-item class="wf-field" prop="group">
      	<div class="fieldname" name="label">
					<b>*</b>
					<span>选择分组</span>
					<span class="fieldinfo">最多50个字</span>
				</div>
        <el-select 
        	class="fieldval"
					v-model="options.group" 
					filterable 
					size="small"
					required
					placeholder="请选择">
			    <el-option
			    	style="width:100%"
			      v-for="item in items"
			      :key="item.value"
			      :label="item.label"
			      :value="item.value">
			    </el-option>
			  </el-select>
      </el-form-item>
    </el-form>
	</section>
</template>
<script>
export default{
	name: "wf-form",
	data() {
		return {
			options:{},
			items: [],
			//验证规则
      formRules: {
        title: [{
          required: true,
          message: "请输入审批名称",
          trigger: 'blur'
        }, {
          min: 2,
          max: 10,
          message: "最多输入10个字符",
          trigger: 'blur'
        }]
      }
		}
	},
	props: {
    form: {
      type: Object,
      default: null
    }
  },
  watch: {
    form() {
      this.options = this.form
    }
  },
	methods: {
		changeInfo() {
			let obj = this.options
			drag.$emit("changeInfo", obj)
		}
	},
	created() {
	},
	mounted() {
		this.options = this.form
	}
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

.wf {

	&-form {
		position:relative;
    width: 500px;
    margin:15px auto;
    box-shadow: 2px 3px 5px rgba(0,0,0,.2);

    &::before {
    	position:absolute;
    	left:8px;
    	top:-12px;
    	@include triangle(up, 2em);
    	color:#fff;
    }

    &-pannel {
    	padding:10px 20px;
    	border-radius:8px;
    	background: #fff;
    }

    /deep/ .el-form-item__label {
    	padding:0;
    }
  }

  &-field {
	  margin-bottom:15px;
	
		.fieldname {
		  display: block;
		  margin-bottom: -10px;
		  font-size: 14px;
		  color: #666;

		  b {
		  	color:$color-danger;
		  	margin-right:5px;
		  }
		}	

		.fieldval {
			width:100%;
		}

		.fieldinfo {
		  color: #ccc;
		  font-size: 12px;
		  margin-left: 8px;
		}
	}
}	
</style>
