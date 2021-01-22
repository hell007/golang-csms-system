// import CryptoJS from 'crypto-js';
// import { CRYPTO } from '@/config';

// export function encryptString(val: string) {
//   let str = CryptoJS.AES.encrypt(val, CRYPTO.Secret).toString();
//   return str;
// }

// export function decryptString(val: string) {
//   let str = CryptoJS.AES.decrypt(val, CRYPTO.Secret).toString(CryptoJS.enc.Utf8);
//   return str;
// }

// function encryptObj(obj: object) {
//   let o = CryptoJS.AES.encrypt(JSON.stringify(obj), CRYPTO.Secret);
//   let str = o.toString();
//   return str;
// }

// function decryptObj(obj: object) {
//   let str: any = CryptoJS.AES.decrypt(JSON.stringify(obj), CRYPTO.Secret);
//   let o: object = JSON.parse(o.toString(CryptoJS.enc.Utf8));
//   return o;
// }

// export function encrypt(data: any) {
//   if (!data) {
//     throw new Error('encryption data is not string or object!');
//     return;
//   }

//   if (typeof data === 'string') {
//     return encryptString(data);
//   } else {
//     return encryptObj(data);
//   }
// }

// export function decrypt(data: any) {
//   if (!data) {
//     throw new Error('decryption data need string!');
//     return;
//   }

//   if (typeof data === 'string') {
//     return decryptString(data);
//   } else {
//     return decryptObj(data);
//   }
// }
