const CSMSKEY = {
	Expire: 1 * 60 * 60,//秒为单位：设置为1小时
	Token: 'csms-system-token',
	User: 'csms-system-user',
	Access: 'csms-system-access',
	Permission: 'csms-system-permission',
}

const CRYPTO = {
	Secret: 'csms-system-jie-secret',
}

const URIS = {
	Goods: "http://127.0.0.1:9000/uploads/goods/",
	Kindeditor: "http://127.0.0.1:9000/goods/product/upload",
}

const Theme = {
	primary: '#F39800',
} 

export {
	CSMSKEY,
	CRYPTO,
	URIS,
	Theme
}