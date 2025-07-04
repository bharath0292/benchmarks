# 🧪 Go Map Benchmark Results

> Comparing performance of different concurrent map implementations in Go.

---

## 🔧 System Info

- **GOOS**: linux
- **GOARCH**: amd64
- **CPU**: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz

---

## 📊 Benchmark Summary

| 🗂️ Map Type       | ⚡ Ops/sec        | ⏱️ Time/op     | 📦 Bytes/op | 📌 Allocs/op | 🧠 Verdict                                  |
|-------------------|------------------|----------------|-------------|--------------|---------------------------------------------|
| `sync.Map`      | ~5.41 M           | 246.3 ns/op    | 145 B       | 3            | Decent for reads, high allocs               |
| `map + RWMutex` | ~3.27 M           | 324.7 ns/op    | 7 B         | 0            | Lowest memory usage, but slower             |
| `cmap`          | ~5.93 M           | 220.0 ns/op    | 81 B        | 1            | Good balance, sharded                       |
| `xsync.Map`     | ~16.56 M          | 88.14 ns/op    | 96 B        | 1            | ✅ Fastest overall, low latency              |

---

## ✅ Conclusion

- `xsync.Map` is the **clear winner**, delivering:
  - 🔼 ~273% faster performance than `cmap`
  - 🔽 Lowest latency per op (~88 ns)
  - 👌 Reasonable memory footprint
- `cmap` still offers a good tradeoff of performance and simplicity.
- `sync.Map` is suitable for read-heavy use cases, but has more overhead.
- `map + RWMutex` is memory efficient but slower due to locking.

---
