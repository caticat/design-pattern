//go:build race && race_demo

package concurrency

import "testing"

// TestIntentionalRace 只在显式启用竞态演示时运行，预期失败。
func TestIntentionalRace(t *testing.T) {
	Example_race()
}
