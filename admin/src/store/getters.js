const getters = {
	//token: state => state.user.token,
  user: state => state.user.admin,
  access: state => state.user.access,
  roles: state => state.user.roles,
  permission_routers: state => state.permission.routers,
  addRouters: state => state.permission.addRouters
}
export default getters