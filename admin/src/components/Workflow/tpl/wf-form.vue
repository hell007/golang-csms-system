<template>
	<section class="wf-form">
		<div class="wf-form-pannel">
			<div class="wf-field">
				<div class="fieldname">
					<b>*</b>
					<span>审批名称</span>
					<span class="fieldinfo">最多50个字</span>
				</div>
				<div class="fieldblock">
					<el-input
					  placeholder="请输入"
					  size="small"
					  maxlength="50"
					  v-model="title"
					  required
					  clearable>
					</el-input>
				</div>
			</div>
			<div class="wf-field">
				<div class="fieldname">
					<b>*</b>
					<span>选择分组</span>
				</div>
				<div class="fieldblock">
					<el-select 
						v-model="group" 
						filterable 
						size="small"
						required
						placeholder="请选择">
				    <el-option
				      v-for="item in options"
				      :key="item.value"
				      :label="item.label"
				      :value="item.value">
				    </el-option>
				  </el-select>
				</div>
			</div>
			<div class="wf-field">
				<div class="fieldname">
					<span>审批说明</span>
					<span class="fieldinfo">最多100个字</span>
				</div>
				<div class="fieldblock">
					<el-input
					  type="textarea"
					  :autosize="{ minRows: 2, maxRows: 4}"
					  placeholder="请输入"
					  maxlength="100"
					  v-model="description">
					</el-input>
				</div>
			</div>
			<!-- <div class="wf-field">
				<div class="fieldname">
					图标
				</div>
			</div> -->
		</div>
	</section>
</template>
<script>
export default{
	name: "wf-form",
	data: function () {
		return {
			title: '',
			description: '',
			group:'',
			options:[]
		}
	},
	methods: {
		// selectIcon: function (e) {
		// 	this.selected = e.currentTarget.getAttribute('data-index')
		// 	let obj = {
		// 		title: this.title,
		// 		icon: e.currentTarget.getAttribute('data-icon'),
		// 		description: this.description
		// 	}
		// 	drag.$emit("changeInfo", obj)
		// },
		// changeInfo: function () {
		// 	let obj = {
		// 		title: this.title,
		// 		icon: this.icons[this.selected].value,
		// 		description: this.description
		// 	}
		// 	drag.$emit("changeInfo", obj)
		// }
	}
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

.wf {

	&-form {
    width: 500px;
    margin:0 auto;
    padding:15px 0;

    &-pannel {
    	padding:10px 20px;
    	border-radius:8px;
    	background: #fff;
    }
  }

  &-field {
	  display: block;
	  margin: 5px 0 30px;
	
		.fieldname {
		  display: block;
		  font-size: 14px;
		  color: #666;
		  margin-bottom: 12px;
		  padding-left: 2px;

		  b {
		  	color:$color-danger;
		  	margin-right:5px;
		  }
		
			span {
			  font-size: 14px;
			}
		}	

		.fieldinfo {
		  color: #ccc;
		  font-size: 12px !important;
		  margin-left: 8px;
		  font-weight: normal;
		  opacity: 0.9;
		}

		.fieldblock {
		  display: block;
		  margin: 5px 0 5px 12px;
		  line-height: normal;
		}
		
	}
}	
</style>
