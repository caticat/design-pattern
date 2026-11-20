package decorator_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/caticat/design-pattern/codes/structural/decorator"
)

// --8<-- [start:example]
func Example() {
	// 一层层包装：美式 → 加牛奶 → 再加奶油
	var b decorator.Beverage = decorator.Coffee{}
	b = decorator.Milk{B: b}
	b = decorator.Cream{B: b}

	fmt.Println(b.Description(), "=", b.Cost(), "元")
	// Output:
	// 美式 + 牛奶 + 奶油 = 23 元
}

// --8<-- [end:example]

// --8<-- [start:example-chain]
func ExampleChain() {
	hello := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("   处理请求")
	})
	h := decorator.Chain(hello, decorator.Logging(os.Stdout), decorator.RequireToken)

	req := httptest.NewRequest("GET", "/hi", nil)
	req.Header.Set("X-Token", "demo")
	h.ServeHTTP(httptest.NewRecorder(), req)
	// Output:
	// -> GET /hi
	//    处理请求
	// <- /hi
}

// --8<-- [end:example-chain]
