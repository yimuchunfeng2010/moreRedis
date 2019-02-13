package test

import ("moreRedis/routes/rpc"
"testing")

func Benchmark_RedisSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rpc.RedisSet("AAAA", "BBBB")
	}
}

func Benchmark_RedisGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rpc.RedisGet("AAAA")
	}
}

func Benchmark_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rpc.Get("AAAA")
	}
}
func Benchmark_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rpc.Set("AAAA","BBBB")
	}
}

func Benchmark_LigthGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rpc.LigthGet("AAAA")
	}
}