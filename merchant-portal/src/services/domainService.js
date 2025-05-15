import api from './api'

// 获取域名列表
export const getDomains = async () => {
  return api.get('/domains')
}

// 添加域名
export const addDomain = async (domainData) => {
  return api.post('/domains', domainData)
}

// 更新域名
export const updateDomain = async (id, domainData) => {
  return api.put(`/domains/${id}`, domainData)
}

// 删除域名
export const deleteDomain = async (id) => {
  return api.delete(`/domains/${id}`)
}