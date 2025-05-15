import DOMPurify from 'dompurify'

// HTML净化
export function sanitizeHTML(html: string) {
  return DOMPurify.sanitize(html)
}

// 安全地设置HTML内容
export function setSafeHTML(element: HTMLElement, html: string) {
  element.innerHTML = sanitizeHTML(html)
}

// 编码特殊字符
export function encodeHTML(str: string) {
  if (!str) return ''
  
  return str
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&#039;')
}