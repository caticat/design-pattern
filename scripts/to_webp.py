"""把生图模型输出的 PNG 转为 webp，写入 docs/assets/images/。

用法：python scripts/to_webp.py <源图片> <模式名>
示例：python scripts/to_webp.py C:/.../xxx.png singleton
"""

import sys
from pathlib import Path

from PIL import Image

QUALITY = 85
OUT_DIR = Path(__file__).resolve().parent.parent / "docs" / "assets" / "images"


def main() -> None:
    if len(sys.argv) != 3:
        sys.exit(__doc__)
    src, name = Path(sys.argv[1]), sys.argv[2]
    OUT_DIR.mkdir(parents=True, exist_ok=True)
    dst = OUT_DIR / f"{name}.webp"
    with Image.open(src) as im:
        im.convert("RGB").save(dst, "WEBP", quality=QUALITY, method=6)
    print(f"{dst}  {src.stat().st_size // 1024}KB -> {dst.stat().st_size // 1024}KB")


if __name__ == "__main__":
    main()
