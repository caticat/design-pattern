"""生成并校验各章的 Go Playground「在线运行」链接。

每章正文中放一行按钮（ID 由本脚本维护）：
    [:material-play: 在 Go Playground 运行](https://go.dev/play/p/<ID>){ .md-button .md-button--primary target=_blank }

上传的程序只有一个文件，与正文「使用示例」一致：
    example_test.go 改成 package main，末尾追加 main() 依次调用各 Example。
模式包本身由 Playground 通过 Go 模块代理从公共仓库拉取（import 路径不变），
因此 --update 前须先把代码推送到远程。

用法：
    python -I scripts/playground.py           # 离线校验链接是否与当前示例一致（自审用）
    python -I scripts/playground.py --update  # 上传、在 Playground 上编译运行验证、改写正文中的链接
"""

import base64
import hashlib
import json
import pathlib
import re
import sys
import time
import urllib.parse
import urllib.request

ROOT = pathlib.Path(__file__).resolve().parent.parent
SHARE_URL = "https://go.dev/_/share"
COMPILE_URL = "https://go.dev/_/compile"
# 与 golang.org/x/playground 的 snippet ID 算法一致：sha256(salt + body) 的 base64url 前 11 位
SALT = b"Go playground salt\n"

LINK_RE = re.compile(r"https://go\.dev/play/p/[\w-]+")
SNIPPET_RE = re.compile(r'--8<-- "(codes/[^/"]+/[^/"]+)/')
EXAMPLE_RE = re.compile(r"^func (Example\w*)\(\) \{", re.M)
MARKER_RE = re.compile(r"^\s*// --8<-- .*\n", re.M)


def build(pkg: str) -> bytes:
    """把某个模式包的 example_test.go 改写成可直接运行的 main 程序。"""
    src = (ROOT / pkg / "example_test.go").read_text(encoding="utf-8")
    src = MARKER_RE.sub("", src)
    if '"fmt"' not in src:
        sys.exit(f"{pkg}/example_test.go 必须导入 fmt")
    names = EXAMPLE_RE.findall(src)
    if not names:
        sys.exit(f"{pkg}/example_test.go 中没有 Example 函数")
    calls = []
    for name in names:
        if len(names) > 1:
            calls.append(f'\tfmt.Println("=== {name} ===")')
        calls.append(f"\t{name}()")
    src = re.sub(r"^package \w+", "package main", src, count=1, flags=re.M)
    src = src.rstrip() + "\n\nfunc main() {\n" + "\n".join(calls) + "\n}\n"
    return src.encode("utf-8")


def snippet_id(body: bytes) -> str:
    b = base64.urlsafe_b64encode(hashlib.sha256(SALT + body).digest()).decode()
    n = 11
    while n <= len(b) and b[n - 1] == "_":
        n += 1
    return b[:n]


def share(body: bytes) -> str:
    req = urllib.request.Request(SHARE_URL, data=body, method="POST")
    # 必须声明纯文本；默认的表单类型会让服务端把代码按 URL 编码存储
    req.add_header("Content-Type", "text/plain; charset=utf-8")
    with urllib.request.urlopen(req, timeout=30) as resp:
        return resp.read().decode().strip()


def run(body: bytes) -> str:
    """在 Playground 上编译运行；首次拉取模块可能超时，重试几次。"""
    data = urllib.parse.urlencode({"body": body.decode(), "version": "2", "withVet": "true"}).encode()
    err = ""
    for _ in range(4):
        with urllib.request.urlopen(COMPILE_URL, data, timeout=90) as resp:
            r = json.load(resp)
        err = r.get("Errors") or r.get("VetErrors") or ""
        if not err:
            return "".join(e["Message"] for e in r.get("Events") or [])
        time.sleep(5)
    sys.exit(f"Playground 运行失败（代码是否已推送？）：\n{err}")


def main() -> None:
    update = "--update" in sys.argv[1:]
    stale = 0
    for md in sorted((ROOT / "docs").rglob("*.md")):
        text = md.read_text(encoding="utf-8")
        links = LINK_RE.findall(text)
        if not links:
            continue
        pkgs = set(SNIPPET_RE.findall(text))
        if len(pkgs) != 1:
            sys.exit(f"{md}: 需恰好引用一个代码包，实际 {sorted(pkgs)}")
        body = build(pkgs.pop())
        want = f"https://go.dev/play/p/{snippet_id(body)}"
        rel = md.relative_to(ROOT).as_posix()
        if not update:
            if all(link == want for link in links):
                print(f"ok      {rel}")
            else:
                print(f"STALE   {rel}  (运行 --update 刷新)")
                stale += 1
            continue
        out = run(body)
        got = f"https://go.dev/play/p/{share(body)}"
        if got != want:
            sys.exit(f"{rel}: 服务端 ID {got} 与本地算法 {want} 不一致")
        md.write_text(LINK_RE.sub(want, text), encoding="utf-8", newline="\n")
        print(f"updated {rel}  {want}\n{out}")
    sys.exit(1 if stale else 0)


if __name__ == "__main__":
    main()
