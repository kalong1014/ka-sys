import DOMPurify from 'dompurify';
// HTML净化
export function sanitizeHTML(html) {
    return DOMPurify.sanitize(html);
}
// 安全地设置HTML内容
export function setSafeHTML(element, html) {
    element.innerHTML = sanitizeHTML(html);
}
// 编码特殊字符
export function encodeHTML(str) {
    if (!str)
        return '';
    return str
        .replace(/&/g, '&amp;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/"/g, '&quot;')
        .replace(/'/g, '&#039;');
}
//# sourceMappingURL=security.js.map