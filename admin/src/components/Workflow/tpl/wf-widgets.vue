<template>
	<section class="wf-panel wf-widget">
		<div class="wf-panel-body">
			<div v-for="(item,index) in components"
				 class="wf-widget-item" 
				 :data-index="index" 
				 :data-type="item.componentName"
				 @mousedown="start">
				<label>
					{{item.name}}
				</label>
				<i class="wf-widget-icon" 
					:class="item.componentName"></i>
			</div>
		</div>
	</section>
</template>
<script>
export default {
	name: 'wf-widgets',
	data(){
		return {
			components: [{
				name: '单行输入框',
				defaultLable: '单行输入框',
				defaultProps: '请输入',
				defaultRequired: false,
				defaultPrint: false,
				componentName: 'textfield',
				supports: ['label', 'placeholder', 'required', 'print']
			},
			{
				name: '多行输入框',
				defaultLable: '多行输入框',
				defaultProps: '请输入',
				defaultRequired: false,
				defaultPrint: false,
				componentName: 'textareafield',
				supports: ['label', 'placeholder', 'required', 'print']
			},
			{
				name: '数字输入框',
				defaultLable: '数字输入框',
				defaultProps: '请输入',
				defaultRequired: false,
				defaultPrint: false,
				componentName: 'numberfield',
				supports: ['label', 'placeholder', 'required', 'print', 'uint']
			},
			{
				name: '单选框',
				defaultLable: '单选框',
				defaultProps: '请选择',
				defaultRequired: false,
				defaultSync: false,
				defaultPrint: false,
				defaultOptions: [
					{idx: 1, text: '选项1'},
					{idx: 2, 'text': '选项2'},
					{idx: 3, text: '选项3'}
				],
				componentName: 'selectfield',
				supports: ['label', 'placeholder', 'options', 'sync', 'required', 'print']
			},
			{
				name: '多选框',
				defaultLable: '多选框',
				defaultProps: '请选择',
				defaultRequired: false,
				defaultPrint: false,
				defaultOptions: [
					{idx: 1, text: '选项1'},
					{idx: 2, 'text': '选项2'},
					{idx: 3, text: '选项3'}
				],
				componentName: 'multiselectfield',
				supports: ['label', 'placeholder', 'options', 'required', 'print']
			},
			{
				name: '日期',
				defaultLable: '日期',
				defaultProps: '请选择',
				defaultRequired: false,
				defaultPrint: false,
				defaultFormat: 'yyyy-MM-dd',
				componentName: 'datefield',
				supports: ['label', 'placeholder', 'dateformat', 'required', 'print']
			},
			{
				name: '日期区间',
				defaultLable: '开始时间',
				defaultLable2: '结束时间',
				defaultProps: '请选择',
				defaultProps2: '请选择',
				defaultRequired: false,
				defaultPrint: false,
				defaultAutorekonTime: false,
				defaultFormat: 'yyyy-MM-dd',
				defaultSubtitle: '时长',
				componentName: 'daterangefield',
				supports: [
					'label',
					'label2',
					'placeholder',
					'placeholder2',
					'dateformat',
					'required',
					'autorekonTime',
					'subtitle',
					'print']
			},
			{
				name: '图片',
				defaultLable: '图片',
				defaultRequired: false,
				defaultPrint: false,
				componentName: 'photofield',
				supports: ['label', 'required', 'print']
			},
			{
				name: '说明文字',
				defaultLable: '说明文字',
				defaultRequired: false,
				defaultProps: '请输入说明文字',
				defaultShow: false,
				defaultPrint: false,
				defaultHref: '',
				defaultShow: false,
				componentName: 'textnote',
				supports: ['textnote', 'show', 'href', 'print']
			},
			{
				name: '附件',
				defaultLable: '附件',
				defaultRequired: false,
				defaultPrint: false,
				componentName: 'attachment',
				supports: ['label', 'required', 'print']
			},
			{
				name: '身份证',
				defaultLable: '身份证',
				defaultProps: '请输入',
				defaultRequired: false,
				defaultPrint: false,
				componentName: 'identityfield',
				supports: ['identity', 'label', 'placeholder', 'required', 'print']
			},
			{
				name: '手机',
				defaultLable: '手机',
				defaultProps: '请输入',
				defaultRequired: false,
				defaultPrint: false,
				componentName: 'phonefield',
				supports: ['label', 'placeholder', 'required', 'print']
			},
			{
				name: '地点',
				defaultLable: '地点',
				defaultProps: '例如：昆明市东来大厦',
				defaultRequired: false,
				defaultPrint: false,
				componentName: 'locationfield',
				supports: ['label', 'required', 'print']
			},
			{
				name: '评分',
				defaultLable: '评分',
				defaultProps: '',
				defaultRate: '5',
				defaultRequired: false,
				defaultPrint: false,
				componentName: 'ratefield',
				supports: ['rate', 'label', 'required', 'print']
			},
			{
				name: '表格/明细',
				defaultLable: '表格',
				defaultType: 'table',//table/detail
				defaultAction: '增加',
				components: [],
				selected:null,
				defaultPrint: false,
				InTableCanvas:null,
				componentName: 'tablefield',
				supports: ['type', 'label', 'action']
			}]
		}
	},
	methods: {
		start(e) {
			let obj = {}
			let dom = e.currentTarget
			let index = dom.getAttribute('data-index')
			let actualLeft = dom.offsetLeft;
			let actualTop = dom.offsetTop;
			let current = dom.offsetParent;
			while (current !== null) {
				actualLeft += current.offsetLeft;
				actualTop += current.offsetTop;
				current = current.offsetParent;
			}
			obj.componentName = dom.getAttribute("data-type")
			obj.componentText = dom.querySelector('label').innerText
			obj.clientX = e.clientX
			obj.clientY = e.clientY
			obj.isStart = true
			obj.componentView = this.components[index]
			//console.log('widgets==', obj)
			drag.$emit("moveStart", obj)
		}
	}
}
</script>
<style lang="scss" scoped>
@import "../../../styles/_global.scss";

.wf {

	&-widget {
		width:350px;
		border-top:1px solid $color-border-gray;
		padding-top:20px;
		background:#fff;
		position:absolute;
		left:0;
		top:0;
		bottom:0;

		&-item {
			width: 158px;
	    height: 32px;
	    margin: 0 0 10px 9px;
	    line-height: 32px;
	    font-size: 12px;
	    text-align: left;
	    padding-left: 15px;
	    cursor: move;
	    color: #191f25;
	    background: #fff;
	    border: 1px solid #d9d9d9;
	    border-radius: 4px;
	    position: relative;
	    overflow: hidden;
	    float:left;

		  &:hover {
			  background: hsla(0,0%,100%,.2);
			}

			label {
				cursor: move;
			}
		}

		&-icon {
			display: inline-block;
		  position: absolute;
		  right: 10px;
		  top: 3px;
		  background-image: url('../../../assets/images/workflow/widget.png');
		  background-repeat:no-repeat;
		  height: 24px;
		  width: 24px;
		  

		  &.textfield {
			  background-position: 0 -24px;
			}

			&.textareafield {
			  background-position: 0 0;
			}

			&.selectfield {
			  background-position: 0 -48px;
			}

			&.datefield {
			  background-position: 0 -72px;
			}

			&.daterangefield {
			  background-position: 0 -96px;
			}

			&.numberfield {
			  background-position: 0 -120px;
			}

			&.photofield {
			  background-position: 0 -144px;
			}

			&.tablefield {
			  background-position: 0 -168px;
			}

			&.externalcontactfield {
			  background: url('../../../assets/images/workflow/widget-contact.png');
			}

			&.attachment {
			  background-position: 0 -192px;
			}

			&.textnote {
			  background-position: 0 -216px;
			}

			&.moneyfield {
			  background-position: 0 -240px;
			}

			&.multiselectfield {
			  background-position: 0 -264px;
			}
		}
	}
}
</style>
