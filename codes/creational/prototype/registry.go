package prototype

import "fmt"

// --8<-- [start:registry]
// Registry 保存配好的原型，按名字复制出新对象。
type Registry struct {
	protos map[string]*Monster
}

func NewRegistry() *Registry {
	return &Registry{protos: map[string]*Monster{}}
}

// Register 存一份副本：调用方之后再改 m，不影响已登记的原型。
func (r *Registry) Register(name string, m *Monster) {
	r.protos[name] = m.Clone()
}

// Spawn 交出的永远是副本，原型本身从不交出去。
func (r *Registry) Spawn(name string) (*Monster, error) {
	p, ok := r.protos[name]
	if !ok {
		return nil, fmt.Errorf("没有名为 %q 的原型", name)
	}
	return p.Clone(), nil
}

// --8<-- [end:registry]
