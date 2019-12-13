import CryptoJS from 'crypto-js'
import {
	CRYPTO
} from '@/api/config'

export function encryptString(val) {
	let str = CryptoJS.AES.encrypt(val, CRYPTO.Secret).toString()
	return str
}

export function decryptString(val) {
	let str = CryptoJS.AES.decrypt(val, CRYPTO.Secret).toString(CryptoJS.enc.Utf8)
	return str
}

function encryptObj(obj) {
	let o = CryptoJS.AES.encrypt(JSON.stringify(obj), CRYPTO.Secret)
	let str = o.toString()
	return str
}

function decryptObj(obj) {
	let str = CryptoJS.AES.decrypt(JSON.stringify(obj), CRYPTO.Secret)
	let o = SON.parse(o.toString(CryptoJS.enc.Utf8))
	return o
}

export function encrypt(data) {
	if (!data) {
		console.Error('encryption data is not string or object!')
		return
	}

	if (typeof data === 'string') {
		return encryptString(data)
	} else {
		return encryptObj(data)
	}
}

export function decrypt(data) {
	if (!data) {
		console.Error('decryption data need string!')
		return
	}

	if (typeof data === 'string') {
		return decryptString(data)
	} else {
		return decryptObj(data)
	}
}