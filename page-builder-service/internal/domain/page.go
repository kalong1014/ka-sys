package domain

import "time"

// 页面元素类型
const (
	ElementTypeTitle  = "title"
	ElementTypeText   = "text"
	ElementTypeImage  = "image"
	ElementTypeButton = "button"
	ElementTypeList   = "list"
)

// 基础元素
type BaseElement struct {
	ID       string            `json:"id"`
	Type     string            `json:"type"`
	Position int               `json:"position"`        // 元素排序位置
	Style    map[string]string `json:"style,omitempty"` // 样式属性（如字体、颜色、边距）
}

// 标题元素
type TitleElement struct {
	BaseElement
	Content string `json:"content"` // 标题内容
	Level   int    `json:"level"`   // 标题级别（1-6级标题）
}

// 文本元素
type TextElement struct {
	BaseElement
	Content string `json:"content"` // 文本内容
}

// 图片元素
type ImageElement struct {
	BaseElement
	URL    string `json:"url"`    // 图片URL
	Alt    string `json:"alt"`    // 替代文本
	Width  int    `json:"width"`  // 宽度（像素）
	Height int    `json:"height"` // 高度（像素）
}

// 按钮元素
type ButtonElement struct {
	BaseElement
	Content string `json:"content"` // 按钮文本
	Link    string `json:"link"`    // 跳转链接
}

// 列表元素
type ListElement struct {
	BaseElement
	Items []string `json:"items"` // 列表项内容
	Type  string   `json:"type"`  // "ol"（有序列表）或 "ul"（无序列表）
}

// 页面结构
type Page struct {
	ID          string        `json:"id"`
	MerchantID  string        `json:"merchant_id"` // 所属商户ID
	Name        string        `json:"name"`        // 页面名称
	Description string        `json:"description"` // 页面描述
	Elements    []interface{} `json:"elements"`    // 元素列表（多态类型）
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

// 创建页面请求
type CreatePageRequest struct {
	MerchantID  string        `json:"merchant_id" validate:"required"`
	Name        string        `json:"name" validate:"required,min=2,max=50"`
	Description string        `json:"description" validate:"max=200"`
	Elements    []interface{} `json:"elements"` // 元素列表（需符合元素类型规范）
}

// 更新页面请求
type UpdatePageRequest struct {
	Elements []interface{} `json:"elements"` // 新的元素列表
}
