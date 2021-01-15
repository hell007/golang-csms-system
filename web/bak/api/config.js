let CSMSKEY = {
	Expire: 1 * 60 * 60,//单位秒：设置为1小时
	Token: 'csms-web-token',
}

let CRYPTO = {
	Secret: 'csms-web-jie-secret',
}

let URIS = {
	Goods: "http://127.0.0.1:7000/uploads/goods/",
	Kindeditor: "http://127.0.0.1:7000/goods/product/upload",
}

export {
	CSMSKEY,
	CRYPTO,
	URIS
}