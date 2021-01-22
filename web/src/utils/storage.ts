import store from 'storejs';
import { CSMSKEY } from '@/config';

export function set(key: string, val: any) {
  let expire = Date.now() + CSMSKEY.expire * 1000;
  store.set(key, {
    value: val,
    expire: expire
  });
}

export function get(key: string) {
  const data = store.get(key);

  if (!data) return null;

  if (data.expire == -1) return data;

  if (Date.now() >= data.expire) {
    console.log('过期==', key);
    store.remove(key);
    store.clear();
    return null;
  } else {
    return data.value;
  }
}

export function remove(key: string) {
  return store.remove(key);
}

export function clear() {
  return store.clear();
}

export function hasToken() {
  return store.has(CSMSKEY.token);
}

export function getToken() {
  return get(CSMSKEY.token);
}
