package chapter11

import "testing"

func TestHello(t *testing.T) {
	// t.Run 用于多 case 测试
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

// 辅助函数
func assertCorrectMessage(t testing.TB, got, want string) {
	// tell test suite  that this method is a helper
	t.Helper() // 可以输出错误调用的代码行号，而不是当前方法的行号
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
