package url_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/colduction/url"
)

func benchPerCoreConfigs(b *testing.B, f func(b *testing.B)) {
	b.Helper()
	coreConfigs := []int{1, 2, 4, 8, 12}
	for _, n := range coreConfigs {
		name := fmt.Sprintf("%d cores", n)
		b.Run(name, func(b *testing.B) {
			runtime.GOMAXPROCS(n)
			f(b)
		})
	}
}

func BenchmarkUrl(b *testing.B) {
	benchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				url_temp := url.New()
				url_temp.Host = "www.google.com"
				url_temp.Scheme = "https"
				_ = url_temp.String()
				url_temp.Reset(true)
			}
		})
	})
}
