<template>
	<section class="wf-setting" v-show="isSetting">
		<h2 class="wf-setting-title">{{settings.name}}</h2>
		<div class="wf-setting-panel">
			
			<div v-if="settings.identity" class="wf-field-warn">
				<i class="el-icon-warning"></i>
				<span>将自动验证身份证号码正确性</span>
			</div>

			<div v-if="settings.label" class="wf-field">
				<div class="fieldname">
					<span>标题</span>
					<span class="fieldinfo">最多10个字</span>
				</div>
				<div class="fieldblock">
					<input type="text" 
						@input="changeComponent" 
						maxlength="10" 
						v-model="settings.defaultLable">
				</div>
			</div>

			<div v-if="settings.label2" class="wf-field">
	      <div class="fieldname">
	        <span>标题2</span>
	        <span class="fieldinfo">最多10个字</span>
	      </div>
	      <div class="fieldblock">
	        <input type="text" 
	        	@input="changeComponent" 
	        	maxlength="10" 
	        	v-model="settings.defaultLable2">
	      </div>
	  	</div>

			<div v-if="settings.action" class="wf-field">
				<div class="fieldname">
					<span>动作名称</span>
					<span class="fieldinfo">最多10个字</span>
				</div>
				<div class="fieldblock">
					<input type="text" 
						@input="changeComponent" 
						maxlength="10" 
						v-model="settings.defaultAction">
				</div>
			</div>

			<div v-if="settings.type" class="wf-field">
				<div class="fieldname">
					<span>填写方式</span>
				</div>
				<div class="fieldblock">
					<el-radio-group 
						v-model="settings.defaultType"
						@change="changeComponent">
				    <div style="margin-bottom:8px;">
				    	<el-radio label="detail">列表</el-radio>
				    </div>
				    <div>
					    <el-radio label="table">表格</el-radio>
					  </div>
				  </el-radio-group>
				</div>
				<div class="fieldblock fieldblock-pic">
					<img src="../../../assets/images/workflow/detail.png"
						v-if="settings.defaultType=='detail'"/>
					<img src="../../../assets/images/workflow/table.png"
						v-if="settings.defaultType=='table'"/>
				</div>
			</div>

			<div v-if="settings.textnote" class="wf-field">
				<div class="fieldname">
					<span>说明文字</span>
					<span class="fieldinfo">最多5000个字</span>
				</div>
				<div class="fieldblock">
					<textarea type="text" 
						@input="changeComponent" 
						maxlength="5000"
						v-model="settings.defaultProps">
						请输入说明文字
					</textarea>
				</div>
			</div>

			<div v-if="settings.placeholder" class="wf-field">
				<div class="fieldname">
					<span>提示文字</span>
					<span class="fieldinfo">最多20个字</span>
				</div>
				<div class="fieldblock">
					<input type="text" 
						@input="changeComponent" 
						maxlength="20" 
						v-model="settings.defaultProps">
				</div>
			</div>

			<div v-if="settings.placeholder2" class="wf-field">
	      <div class="fieldname">
	        <span>提示文字</span>
	        <span class="fieldinfo">最多20个字</span>
	      </div>
	      <div class="fieldblock">
	        <input type="text" 
	        	@input="changeComponent" 
	        	maxlength="20" 
	        	v-model="settings.defaultProps2">
	      </div>
	  	</div>

			<div v-if="settings.href" class="wf-field">
				<div class="fieldname">
					<span>可以输入链接跳转地址</span>
				</div>
				<div class="fieldblock">
					<textarea type="text" 
						@input="changeComponent"  
						v-model="settings.defaultHref">
					</textarea>
				</div>
			</div>

			<div v-if="settings.dateformat" class="wf-field">
				<div class="fieldname">
					<span>日期类型</span>
				</div>
				<div class="fieldblock">
					<el-radio-group 
						v-model="settings.defaultFormat"
						@change="changeComponent">
				    <div style="margin-bottom:8px;">
				    	<el-radio label="yyyy-MM-dd HH:mm">年-月-日 时:分</el-radio>
				    </div>
				    <div>
					    <el-radio label="yyyy-MM-dd">年-月-日</el-radio>
					  </div>
				  </el-radio-group>	
				</div>
			</div>

			<div v-if="settings.rate" class="wf-field">
				<div class="fieldname">
					<span>评分制</span>
				</div>
				<div class="fieldblock">
					<el-radio-group 
						v-model="settings.defaultRate"
						@change="changeComponent">
				    <div style="margin-bottom:8px;">
				    	<el-radio label="5">5分制</el-radio>
				    </div>
				    <div>
					    <el-radio label="10">10分制</el-radio>
					  </div>
				  </el-radio-group>	
				</div>
			</div>

			<div v-if="settings.uint" class="wf-field">
				<div class="fieldname">
					<span>单位</span>
					<span class="fieldinfo">最多20个字</span>
				</div>
				<div class="fieldblock">
					<input type="text" 
						@input="changeComponent" 
						v-model="settings.defaulUint">
				</div>
			</div>

			<div v-if="settings.defaultOptions" class="wf-field">
				<div class="fieldname">
					<span>选项</span>
					<span class="fieldinfo">最多200项,每项最多20个字</span>
				</div>
				<div
					v-bind:class="{limitdel:settings.defaultOptions.length<=1,limitadd:settings.defaultOptions.length>=200}">
					<div v-for="(n,index) in settings.defaultOptions" 
						class="fieldblock fieldblock-options">
						<input type="text" maxlength="20" v-model="n.text">
						<div class="fieldblock-options-btn">
							<span @click="del" 
								v-bind:data-index="index" 
								class="del">
								<i class="el-icon-remove"></i>
							</span>
							<span @click="add" 
								v-bind:data-index="index" 
								class="add">
								<i class="el-icon-circle-plus"></i>
							</span>
						</div>
					</div>
				</div>
			</div>

			<div v-if="settings.required" class="wf-field">
				<div class="fieldname">验证</div>
				<div class="fieldblock">
					<el-checkbox 
						v-model="settings.defaultRequired"
						@change="changeComponent">必填</el-checkbox>
				</div>
			</div>

			<div v-if="settings.autorekonTime">
				<div class="wf-field">
					<div class="fieldname">自动计算时长</div>
				</div>
				<div class="wf-field">
					<div class="fieldblock">
						<el-checkbox 
							v-model="settings.defaultAutorekonTime"
							@change="changeComponent">开启
						</el-checkbox>
					</div>
				</div>
				<div v-if="settings.defaultAutorekonTime">
					<div class="wf-field">
						<div class="fieldname">
							<span>标题</span>
							<span class="fieldinfo">最多10个字</span>
						</div>
						<div class="fieldblock">
							<input type="text" 
								maxlength="10" 
								v-model="settings.defaultSubtitle">
						</div>
					</div>
				</div>
			</div>

			<div v-if="settings.sync" class="wf-field">
				<div class="fieldname">同步到考勤</div>
				<div class="fieldblock">
					<el-checkbox 
						v-model="settings.defaultSync"
						@change="changeComponent">开启
					</el-checkbox>
				</div>
			</div>

			<div v-if="settings.textnote" class="wf-field">
				<div class="fieldname">显示</div>
				<div class="fieldblock">
					<el-checkbox 
						v-model="settings.defaultShow"
						@change="changeComponent">不在审批页显示
					</el-checkbox>
				</div>
			</div>

			<div v-if="settings.print" class="wf-field">
				<div class="fieldname">打印</div>
				<div class="fieldblock">
					<el-checkbox 
						v-model="settings.defaultPrint"
						@change="changeComponent">参与打印
						<span class="fieldinfo">（如不勾选打印时，不显示此项）</span>
					</el-checkbox>
				</div>
			</div>
		</div>	
	</section>
</template>
<script>
export default{
	name: 'wf-settings',
	data() {
		return {
			isSetting: false,
			settings: {}
		}
	},
	methods: {
		add(e) {
			e.stopPropagation()
			e.preventDefault()
			let index = e.currentTarget.getAttribute('data-index');
			for (let i = 0, l = this.settings.defaultOptions.length; i < l; i++) {
				let has = false;
				for (let item in this.settings.defaultOptions) {
					if (this.settings.defaultOptions[item].idx == (i + 1)) {
						has = true
					}
				}
				if (!has) {
					this.settings.defaultOptions.splice((+index + 1), 0, {idx: i + 1, text: '选项' + (i + 1)})
					return
				}
			}
			if (index == this.settings.defaultOptions.length - 1) {
				this.settings.defaultOptions.push({
					idx: (this.settings.defaultOptions.length + 1),
					text: '选项' + (this.settings.defaultOptions.length + 1)
				})
			} else {
				this.settings.defaultOptions.splice((+index + 1), 0, {
					idx: (this.settings.defaultOptions.length + 1),
					text: '选项' + (this.settings.defaultOptions.length + 1)
				})
			}
		},
		del(e) {
			e.stopPropagation()
			e.preventDefault()
			let index = e.currentTarget.getAttribute('data-index');
			this.settings.defaultOptions.splice(index, 1)
		},
		changeComponent() {
			drag.$emit("changeComponent", this.settings);
		},
		showSetting() {
			drag.$emit("showSetting", this.isSetting);
		}
	},
	watch:{
	},
	created() {
		let self = this
		drag.$on("selectComponent", function (obj) {
			self.settings = {}
			for (let i = 0; i < obj.supports.length; i++) {
				self.settings[obj.supports[i]] = true
			}
			self.settings = Object.assign({}, self.settings, obj)
			self.isSetting = true
			console.log('settings==', self.settings)
		})

		drag.$on("showSetting", function(bool){
			self.isSetting = bool
		})
	}
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

.wf {

	&-setting {
    width: 330px;
    background: #fff;
    border-top:1px solid $color-border-gray;
    position:absolute;
    right:0;
    top:0;
    bottom:0;

    &-title {
    	padding:8px 20px;
    	font-size:16px;
    	color:$color-32;
    	border-bottom:1px solid $color-border-gray;
    }

    &-panel {
    	padding: 10px 20px 20px;
    }
  
	  input[type="text"],
	  textarea {
	    background: rgba(255,255,255,0.1);
	    border: 1px solid #d9d9d9;
	    color: #333;
	    padding: 6px 10px;
	    width: 290px;
	    border-radius: 2px;
	    outline: none;
	    font-size: 12px;
	  }

	  input[type="text"]:focus,
	  textarea:focus {
	    border-color: $color-primary;
	    background: transparent;
	  }
  }

  &-field {
	  display: block;
	  margin: 5px 0 30px;

	  &-warn {
	  	padding: 9px 16px;
    	background: #deeffd;
    	border: 1px solid #38adff;
    	border-radius: 4px;

    	i {
    		color:#38adff;
    		margin-right:5px;
    	}
	  }
	
		.fieldname {
		  display: block;
		  font-weight: bold;
		  font-size: 14px;
		  color: #666;
		  margin-bottom: 12px;
		  padding-left: 2px;
		
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
		  margin: 5px 0;
		  line-height: normal;

	    textarea {
		    min-height: 4.5em;
		  }

		  &-pic {
		  	overflow:hidden;
		  	margin-top:15px;

		  	img {
		  		width:100%;
		  	}
		  }

		  &-options {
		  	@include flex-row();
		  	align-items:center;

		  	input[type="text"] {
			    width: 200px;
			  }

		  	&-btn {
					flex:1;
					text-align:right;

					.add,
					.del {
						padding:2px 5px;
						font-size:18px;
						color:$color-primary;
						cursor:pointer;
					}
		  	}
		  }
		}
		
	}
}	
</style>
