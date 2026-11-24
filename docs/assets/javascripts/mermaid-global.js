// 把 Mermaid 的 esbuild 构建暴露为全局 mermaid，供主题的图表渲染器使用。
// dist/mermaid.min.js 的 API 在模块对象的 default 上。
window.mermaid = window.__esbuild_esm_mermaid_nm.mermaid.default;
