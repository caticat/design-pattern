// Package command 演示命令模式：把设备操作封装后交给队列。
package command

// --8<-- [start:command]
// Command 是一条可以执行的命令。
type Command interface {
	Execute()
}

// Lamp 是实际执行开灯动作的设备。
type Lamp struct{ On bool }

// TurnOn 打开台灯。
func (l *Lamp) TurnOn() { l.On = true }

// LightOn 把「打开哪盏灯」封装成一条命令。
type LightOn struct{ Lamp *Lamp }

func (c LightOn) Execute() { c.Lamp.TurnOn() }

// --8<-- [end:command]

// --8<-- [start:queue]
// Queue 暂存命令，并按提交顺序执行。
type Queue struct{ pending []Command }

// Add 把命令排到队尾。
func (q *Queue) Add(c Command) { q.pending = append(q.pending, c) }

// Run 执行开始时已经入队的命令；执行期间新入队的命令留到下次。
func (q *Queue) Run() {
	batch := q.pending
	q.pending = nil
	for _, c := range batch {
		c.Execute()
	}
}

// --8<-- [end:queue]

// Fan 是实际执行启动动作的设备。
type Fan struct{ Running bool }

// Start 启动风扇。
func (f *Fan) Start() { f.Running = true }

// --8<-- [start:func]
// Func 把无参函数适配成命令。
type Func func()

func (f Func) Execute() { f() }

// --8<-- [end:func]
