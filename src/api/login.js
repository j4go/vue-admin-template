import request from '@/utils/request'

export function login(username, password) {
  return request({
    url: '/api/user/login',
    method: 'post',
    data: {
      username,
      password
    }
  })
}

// {"code":20000,"data":{"roles":["admin"],"name":"admin",
// "avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"}}
export function getInfo(token) {
  return request({
    url: '/api/user/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/api/user/logout',
    method: 'post'
  })
}
