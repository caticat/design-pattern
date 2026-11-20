# 图示规范

## 分工

| 图类型 | 生成方式 | 位置 |
|---|---|---|
| 封面大图（每章 1 张，章首页 `{ .cover }`） | 生图模型 | `docs/assets/images/<模式>.webp` |
| 类图 / 时序图 / 前后对比 | Mermaid，写在 Markdown 中 | 正文 |

## 生图模型（强制）

- 模型：`gpt-image-2.6-sunburst-reverse`（provider `work`，source_agent `claude-code`）
- 流程：`image.generate` 的 `prepare` → `request`；每张图单独 prepare
- 参数：`size: "1536x1024"`、`quality: "high"`、`n: 1`
- 不用 `image.edit`（图生图）：网关的 `/v1/images/edits` 只收 JSON，不收参考图上传。角色一致性靠下面的**固定角色描述**保证
- 不要使用 `gemini-nano-banana-2.1`（上游 404）

## 角色设定（定稿）

设定图：`spec/assets/mascot-sheet.webp`（正面 / 侧面 / 挥手 / 讲解四个姿势）。每张插图生成后都要和它对照。

| 部位 | 特征 |
|---|---|
| 质感 | **光滑无毛**，像软胶玩具；轮廓全是平滑曲线 |
| 身体 | 豆形，无脖子，浅青 `#6AD7E5`，奶油色 `#FFF3DC` 椭圆肚皮 |
| 头部 | 两只小圆耳（奶油色内圈）；两个紧挨的大白圆眼，黑瞳孔带一个高光；奶油色口鼻、深棕椭圆鼻、两颗白方门牙 |
| 配件 | 黄色 `#FDC93B` 安全帽（中脊 + 短檐）；棕色工具腰带、金色方扣、左袋银扳手、右袋红柄螺丝刀 |
| 线条 | 统一中等粗细的深青描边 `#1E4D5C`，平涂，最多一层浅阴影 |

## Prompt 模板

固定前缀（一字不改，角色一致性全靠它）：

```text
Flat cartoon vector illustration, clean smooth curves, limited palette of light teal, warm yellow, coral and off-white, at most one subtle flat shadow tone, no gradients, plain light cream background, no text, no letters, no numbers, no words anywhere in the image.

Mascot (keep exactly this design in every appearance): a cute chubby gopher-like character with a SMOOTH, HAIRLESS body like a soft vinyl toy. Absolutely no fur, no hair, no fuzzy or jagged edges, no whiskers, no tufts, no texture strokes.
- Rounded bean / pear shaped body, no neck, solid flat light teal (#6AD7E5), cream (#FFF3DC) oval belly patch.
- Two small round ears with cream inner circles.
- Two large perfect white circle eyes set close together, black round pupils with one small white highlight.
- Small cream muzzle, small dark brown oval nose, two white rectangular front teeth.
- Short stubby arms, tiny round feet.
- Glossy yellow (#FDC93B) engineer hard hat with a center ridge and short brim.
- Brown tool belt with a square gold buckle, silver wrench in left pouch, red-handled screwdriver in right pouch.
- Uniform medium-weight dark teal outline (#1E4D5C) on every shape in the whole image.
```

接着写场景（场景里只写道具和动作，**不要再描述角色外观**，避免与前缀冲突）：

```text
Scene (design pattern '<Pattern>' metaphor): ONE single unified scene, no panels, no borders, no frames, no dividing lines. <场景：一个画面里同时看出「模式的核心关系」>. The mascot appears only once. Composition with generous whitespace, wide framing, suitable as a book chapter cover.
```

## 规则

- 图内**不放文字/数字**，说明写在图注；钟面、路牌等道具注明 no numerals / blank。
- 构图为**单一场景**（参照 hello-algo 章首图）：不分格、不加边框，吉祥物只出现一次；「问题 → 解决」的过程交给正文 Mermaid 图。样例：单例 = 全城一座钟楼，各家窗户以虚线连向同一座钟。
- 网关会忽略 `output_format`，实际返回 PNG（约 1.5MB）；必须用 `scripts/to_webp.py` 转为 webp（约 100–150KB）后再入库，PNG 不入库。
- 生成后必须**本人查看**，不合格重新生成：
    - 与设定图对照：无毛发/毛刺/胡须、五官和配件齐全、描边粗细一致
    - 无文字、隐喻正确
- 吉祥物已定稿，不再更换或增减配件；README 注明 Go Gopher 原作者 Renée French（CC BY）。
- 站点 logo / favicon：`docs/assets/images/mascot.webp`（320px）、`favicon.png`（64px），用固定前缀生成「深青圆形徽章 + 竖大拇指」图，圆外裁透明；小结 `takeaway` 框复用 mascot.webp。
