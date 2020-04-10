module.exports = {
  plugins: {
    autoprefixer: {},
    'postcss-px-to-viewport': ({
			unitToConvert: "px", // 默认值`px`，需要转换的单位
			unitPrecision: 3, // 指定`px`转换为视窗单位值的小数位数
			viewportWidth: 375, // 视窗的宽度，设计稿的宽度
			viewportHeight: 667, //视窗的高度，也可以不配置
			viewportUnit: "vw", //指定需要转换成的视窗单位，建议使用vw
			selectorBlackList: ['.ignore-'], // 指定不转换为视窗单位的类，可以自定义，可以无限添加,建议定义一至两个通用的类名
			minPixelValue: 1, // 小于或等于`1px`不转换为视窗单位，你也可以设置为你想要的值
			mediaQuery: false, // 允许在媒体查询中转换`px`
			exclude: ['src/style/base','src/style/uitls'], // 忽略某些文件夹下的文件或特定文件
			landscape: false, // 是否添加根据landscapeWidth生成的媒体查询条件
			landscapeUnit: "vw", // 横屏时使用的单位
  	 landscapeWidth: 1134 // 横屏时使用的视窗宽度
		})
  }
}
