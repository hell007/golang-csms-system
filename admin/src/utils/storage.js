import store from 'storejs'
import {
	CSMSKEY
} from '@/config'

export function set(key, val) {
	let expire = Date.now() + CSMSKEY.Expire * 1000
	store.set(key, {
		value: val,
		expire: expire
	})
}

export function get(key) {
	const data = store.get(key)

	if (!data) return null

	if (data.expire == -1) return data

	if (Date.now() >= data.expire) {
		console.log('过期==', key)
		store.remove(key)
		store.clear()
		return null
	} else {
		return data.value
	}
}

export function remove(key) {
	return store.remove(key)
}

export function clear() {
	return store.clear()
}

export function hasToken() {
	return store.has(CSMSKEY.Token)
}

export function getToken() {
  return get(CSMSKEY.Token)
}

export function setPermission(permission) {
	return store.set(CSMSKEY.Permission, permission)
}

export function getPermission() {
	return store.get(CSMSKEY.Permission)
}
