package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthClient_ValidateToken(t *testing.T) {
	// 创建模拟认证服务
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "valid_token" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))
	defer server.Close()

	client := NewAuthClient(server.URL)

	// 测试有效令牌
	valid, err := client.ValidateToken("valid_token")
	assert.NoError(t, err)
	assert.True(t, valid)

	// 测试无效令牌
	valid, err = client.ValidateToken("invalid_token")
	assert.NoError(t, err)
	assert.False(t, valid)
}

type AuthClient struct {
	baseURL string
}

func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{baseURL: baseURL}
}

func (c *AuthClient) ValidateToken(token string) (bool, error) {
	req, err := http.NewRequest("GET", c.baseURL+"/validate", nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK, nil
}

func (c *AuthClient) CheckPermission(token string, role string) (bool, error) {
	req, err := http.NewRequest("GET", c.baseURL+"/permission?role="+role, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK, nil
}

func TestAuthClient_CheckPermission(t *testing.T) {
	// 创建模拟认证服务
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		role := r.URL.Query().Get("role")

		if token == "admin_token" && role == "admin" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"has_permission": true}`))
			return
		}

		w.WriteHeader(http.StatusForbidden)
	}))
	defer server.Close()

	client := NewAuthClient(server.URL)

	// 测试有权限的请求
	hasPermission, err := client.CheckPermission("admin_token", "admin")
	assert.NoError(t, err)
	assert.True(t, hasPermission)

	// 测试无权限的请求
	hasPermission, err = client.CheckPermission("user_token", "admin")
	assert.NoError(t, err)
	assert.False(t, hasPermission)
}
