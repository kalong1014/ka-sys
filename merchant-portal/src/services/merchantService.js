import api from './api'

// 获取商户信息
export const getMerchantInfo = async () => {
  return api.get('/merchant/info')
}

// 更新商户信息
export const updateMerchantInfo = async (merchantData) => {
  return api.put('/merchant/info', merchantData)
}