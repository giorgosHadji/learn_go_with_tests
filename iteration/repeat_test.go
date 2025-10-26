package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

// to run benchmarks do :$ go test -bench="."
func BenchmarkRepeat(b *testing.B) {
	// b.loop didn't work for my current (maybe old? go version  version go1.23.3 windows/amd64)
	// for b.Loop() {
	// 	Repeat("a")
	// }
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
func BenchmarkRepeatFaster(b *testing.B) {

	for i := 0; i < b.N; i++ {
		RepeatFaster("a")
	}
}
