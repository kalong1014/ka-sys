#!/bin/bash

# 测试环境配置
AUTH_SERVICE_URL="http://localhost:8090"
GATEWAY_URL="http://localhost:8089"

# 1. 测试auth-service注册和登录
echo "=== 测试auth-service ==="

# 注册用户
echo "注册用户..."
register_resp=$(curl -s -X POST "$AUTH_SERVICE_URL/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpassword",
    "email": "test@example.com",
    "role": "user"
  }')

echo "注册响应: $register_resp"

# 用户登录
echo "用户登录..."
login_resp=$(curl -s -X POST "$AUTH_SERVICE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpassword"
  }')

echo "登录响应: $login_resp"

# 提取令牌
token=$(echo "$login_resp" | grep -o '"token":"[^"]*"' | grep -o '"[^"]*"$' | tr -d '"')
echo "获取的令牌: $token"

# 验证令牌
echo "验证令牌..."
validate_resp=$(curl -s -X POST "$AUTH_SERVICE_URL/api/v1/auth/validate-token" \
  -H "Authorization: $token")

echo "令牌验证响应: $validate_resp"


# 2. 测试gateway-service
echo "=== 测试gateway-service ==="

# 配置网关路由（需要先启动网关）
echo "配置网关路由..."
route_resp=$(curl -s -X POST "$GATEWAY_URL/api/admin/gateway/routes" \
  -H "Content-Type: application/json" \
  -d '{
    "path": "/api/test",
    "method": "GET",
    "service_name": "echo-service",
    "service_addr": "https://httpbin.org",
    "auth_required": true
  }')

echo "路由配置响应: $route_resp"

# 测试带有效令牌的请求
echo "测试带有效令牌的请求..."
valid_token_resp=$(curl -s -i -X GET "$GATEWAY_URL/api/test/get" \
  -H "Authorization: $token")

echo "有效令牌请求响应: $valid_token_resp"

# 测试不带令牌的请求
echo "测试不带令牌的请求..."
no_token_resp=$(curl -s -i -X GET "$GATEWAY_URL/api/test/get")

echo "无令牌请求响应: $no_token_resp"

# 验证结果
if echo "$no_token_resp" | grep -q "401 Unauthorized"; then
  echo "无令牌请求测试通过"
else
  echo "无令牌请求测试失败"
fi

if echo "$valid_token_resp" | grep -q "200 OK"; then
  echo "有效令牌请求测试通过"
else
  echo "有效令牌请求测试失败"
fi