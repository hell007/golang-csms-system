<template>
  <div class="kindeditor">
    <textarea class='textarea' 
      ref="kindeditor" 
      v-model="localValue" 
      name="content">
    </textarea>
  </div>
</template>

<script>
import '../../../node_modules/kindeditor/kindeditor-all.js'
import '../../../node_modules/kindeditor/lang/zh-CN.js'
import '../../../node_modules/kindeditor/themes/default/default.css'

//跨域上传解决方案
//https://blog.csdn.net/lanseguhui/article/details/101375903
export default {
  name: 'kindeditor',
  props: ['options', 'value'],
  data() {
    return {
      editor: null,
      hasInit: false,
      localValue: '',
      items: [
        'source', '|', 
        'undo', 'redo', '|', 
        //'preview', 
        //'print', 
        'template', 'code', 'cut', 'copy', 'paste', 'plainpaste', 'wordpaste', '|', 
        'justifyleft', 'justifycenter', 'justifyright', 'justifyfull', 'insertorderedlist', 'insertunorderedlist', 'indent', 'outdent', 'subscript', 'superscript', 'clearhtml', 'quickformat', 'selectall', '|', 
        //'fullscreen', 
        '/',
        'formatblock', 'fontname', 'fontsize', '|', 
        'forecolor', 'hilitecolor', 'bold', 'italic', 'underline', 'strikethrough', 'lineheight', 'removeformat', '|', 
        'image', 'multiimage', 'flash', 'media', 'insertfile', 'table', 'hr', 'emoticons', 
        //'baidumap', 
        'pagebreak',
        'anchor', 'link', 'unlink', '|',
        'about'
      ]
    }
  },
  watch: {
    localValue: 'updateValue',
    value: 'setDefault'
  },
  mounted() {
    this.initKindeditor()
  },
  methods: {
    initKindeditor() {
      var self = this
        // 默认参数
      var options = {
          uploadJson: '',
          width: '100%',
          height: '200',
          items: this.items,
          emoticonsPath: './img/emoticons/',//表情路径
          afterCreate() {
            self.hasInit = true
          },
          afterChange() {
            self.hasInit = false
            self.localValue = this.html()
          }
        }
        // 调用外部参数
      if (self.options) {
        for (var i in self.options) {
          options[i] = self.options[i]
        }
      }
      self.editor = window.KindEditor.create(self.$refs.kindeditor, options)
    },
    // 设置初始值
    setDefault() {
      if (this.hasInit) {
        this.editor.html(this.value)
      }
    },
    // 修改父件的值
    updateValue() {
      this.$emit('input', this.localValue)
    }
  }
}
</script>
