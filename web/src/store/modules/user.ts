import { MutationTree, ActionTree } from 'vuex';
import { fetchGet } from '@/utils/api';
import { URIS } from '@/config';

interface UserState {
  token: string;
  user: any;
  addressList: any;
}

const state: UserState = {
  token: '',
  user: {},
  addressList: []
};

const mutations: MutationTree<UserState> = {
  SET_TOKEN: (state, token) => {
    state.token = token;
  },
  SET_USER: (state, user) => {
    state.user = user;
  },
  SET_ADDRESSLIST: (state, list) => {
    state.addressList = list;
  }
};

const actions: ActionTree<UserState, {}> = {
  //获取会员收货地址
  getUserAddressList({ commit }, query) {
    return new Promise((resolve, reject) => {
      (async () => {
        try {
          let response = await fetchGet(URIS.user.address, query);
          const ok = response.data;
          const list = response.data.data;
          if (ok) {
            commit('SET_ADDRESSLIST', list);
          }
          resolve(response);
        } catch (ex) {
          reject(ex);
        }
      })();
    });
  },
  //单个收货地址
  getUserAddress({}, query) {
    return new Promise((resolve, reject) => {
      try {
        for (let item of state.addressList) {
          if (item.id == query.aid) {
            resolve(item);
          }
        }
      } catch (ex) {
        reject(ex);
      }
    });
  },
  //获取城市列表
  getCityList({}, query) {
    return new Promise((resolve, reject) => {
      (async () => {
        try {
          let response = await fetchGet(URIS.user.city, query);
          resolve(response);
        } catch (ex) {
          reject(ex);
        }
      })();
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
