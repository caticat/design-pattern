package interpreter

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// --8<-- [start:leaves]
// Parcel 保存包裹的标签，缺少的标签视为 false。
type Parcel map[string]bool

// Expr 是可对包裹求值的规则表达式。
type Expr interface {
	Eval(Parcel) bool
}

type tagExpr string

func (e tagExpr) Eval(p Parcel) bool { return p[string(e)] }

// --8<-- [end:leaves]

// --8<-- [start:operators]
type notExpr struct{ inner Expr }

func (e notExpr) Eval(p Parcel) bool { return !e.inner.Eval(p) }

type andExpr struct{ left, right Expr }

func (e andExpr) Eval(p Parcel) bool { return e.left.Eval(p) && e.right.Eval(p) }

type orExpr struct{ left, right Expr }

func (e orExpr) Eval(p Parcel) bool { return e.left.Eval(p) || e.right.Eval(p) }

// --8<-- [end:operators]

// --8<-- [start:parse]
// Parse 把标签规则解析为表达式树，只支持标签、括号、!、&& 和 ||。
func Parse(rule string) (Expr, error) {
	node, err := parser.ParseExpr(rule)
	if err != nil {
		return nil, err
	}
	return build(node)
}

func build(node ast.Expr) (Expr, error) {
	switch n := node.(type) {
	case *ast.Ident:
		return tagExpr(n.Name), nil
	case *ast.ParenExpr:
		return build(n.X)
	case *ast.UnaryExpr:
		if n.Op == token.NOT {
			inner, err := build(n.X)
			if err != nil {
				return nil, err
			}
			return notExpr{inner: inner}, nil
		}
	case *ast.BinaryExpr:
		return buildBinary(n)
	}
	return nil, fmt.Errorf("不支持的规则节点：%T", node)
}

// --8<-- [end:parse]

// --8<-- [start:binary]
func buildBinary(n *ast.BinaryExpr) (Expr, error) {
	if n.Op != token.LAND && n.Op != token.LOR {
		return nil, fmt.Errorf("不支持的运算符：%s", n.Op)
	}
	left, err := build(n.X)
	if err != nil {
		return nil, err
	}
	right, err := build(n.Y)
	if err != nil {
		return nil, err
	}
	if n.Op == token.LAND {
		return andExpr{left: left, right: right}, nil
	}
	return orExpr{left: left, right: right}, nil
}

// --8<-- [end:binary]
