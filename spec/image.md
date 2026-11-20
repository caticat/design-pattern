# 图示规范

## 分工

| 图类型 | 生成方式 | 位置 |
|---|---|---|
| 类比插图（每章 1 张，章首） | 生图模型 | `docs/assets/images/<模式>.webp` |
| 类图 / 时序图 / 前后对比 | Mermaid，写在 Markdown 中 | 正文 |

## 生图模型（强制）

- 模型：`gpt-image-2.6-sunburst-reverse`（provider `work`，source_agent `claude-code`）
- 流程：`prepare` → `request`；每张图单独 prepare
- 参数：`size: "1536x1024"`、`quality: "high"`、`output_format: "png"`、`n: 1`
- 不要使用 `gemini-nano-banana-2.1`（上游 404）

## Prompt 模板

固定前缀（一字不改）：

```text
Flat cartoon illustration, clean vector style, soft rounded shapes, limited palette of teal (#00ADD8), warm yellow, coral and off-white, subtle flat shadows, plain light cream background, no gradients, no text, no letters, no numbers, no words anywhere in the image.

Mascot: an original cute chubby gopher-like character with big round eyes, two small front teeth, light teal fur, wearing a small yellow engineer hard hat and a tool belt. Friendly, curious expression.
```

接着写场景：

```text
Scene (design pattern '<Pattern>' metaphor): <三格叙事：问题 → 引入模式 → 解决，左到右用珊瑚色箭头连接>. Composition with generous whitespace, wide framing, suitable as a book chapter header.
```

## 规则

- 图内**不放文字/数字**，说明写在图注。
- 构图优先三格叙事（问题 → 模式 → 解决）。
- 风格参考：`docs/assets/images/adapter.webp`（首张测试图，后续适配器章节直接复用）。
- 网关会忽略 `output_format`，实际返回 PNG（约 1.6MB）；必须用 `scripts/to_webp.py` 转为 webp（约 100–150KB）后再入库，PNG 不入库。
- 生成后必须**本人查看**：吉祥物特征（青色、黄安全帽、工具腰带）、无文字、隐喻正确；不合格重新生成。
- 吉祥物已定稿，不再更换或增减配件；README 注明 Go Gopher 原作者 Renée French（CC BY）。
- 站点 logo / favicon：`docs/assets/images/mascot.webp`（320px）、`favicon.png`（64px），由吉祥物定稿图裁成透明圆角；小结 `takeaway` 框复用 mascot.webp。
