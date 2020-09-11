<template>
	<div class="wf-dragging-proxy" v-bind:style="cc">
		<label>
			{{componentText}}
		</label>
		<i class="wf-dragging-icon" 
			:class="componentName"></i>
	</div>
</template>
<script>
	export default{
		name: 'wf-dragging',
		data: function () {
			return {
				componentName: '',
				componentText: '',
				cc: {},
				isStart: false
			}
		},
		created: function () {
			let self = this

			drag.$on('moveStart', function (obj) {
				self.cc = Object.assign({}, self.cc, {
					display: 'none',
					top: obj.clientY + 'px',
					left: obj.clientX + 'px'
				})
				self.componentName = obj.componentName
				self.componentText = obj.componentText
				self.isStart = true
			})

			drag.$on('move', function (obj) {
				if (!self.isStart) {
					return
				}
				let clientX = obj.clientX
				let clientY = obj.clientY
				let startX = parseInt(self.cc.left);
				let startY = parseInt(self.cc.top);
				let moveX = clientX - startX + 'px'
				let moveY = clientY - startY + 'px'
				self.cc = Object.assign({}, self.cc, {
					display: 'block',
					top: startY + 'px',
					left: startX + 'px',
					transform: 'translate3d(' + moveX + ',' + moveY + ',0)'
				})
				obj.preventDefault()
			})

			drag.$on('dragend', function (obj) {
				self.isStart = false
				let startX = self.cc.left;
				let startY = self.cc.top;
				self.cc = Object.assign({}, self.cc, {
					display: 'none',
					top: startY + 'px',
					left: startX + 'px'
				});
				document.querySelector('html').classList.remove('wf-cursor-move');
			})
		}
	}
</script>
<style lang="scss" scoped>
.wf-dragging {
	
	&-proxy {
		width: 158px;
    height: 32px;
    line-height: 32px;
    font-size: 12px;
    text-align: left;
    cursor: move;
    color: #191f25;
    background: #fff;
    border: 1px solid #d9d9d9;
    border-radius: 4px;
    position: fixed;
  	pointer-events: none;
    overflow: hidden;
    display: none;
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
</style>