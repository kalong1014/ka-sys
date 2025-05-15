package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// AuthClient 认证服务客户端
type AuthClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// 验证令牌
func (c *AuthClient) ValidateToken(token string) (bool, error) {
	url := c.baseURL + "/api/v1/auth/validate-token"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("验证令牌失败: %s", string(body))
		return false, nil
	}

	return true, nil
}

// 检查权限
func (c *AuthClient) CheckPermission(token, role string) (bool, error) {
	url := c.baseURL + "/api/v1/auth/check-permission?role=" + role

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("权限检查失败: %s", string(body))
		return false, nil
	}

	// 解析响应
	var result struct {
		HasPermission bool `json:"has_permission"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false, err
	}

	return result.HasPermission, nil
}
