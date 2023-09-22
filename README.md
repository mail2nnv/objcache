# `objcache` investigation

## Motivation

To justify the choice of implementation of the [Voedger object cache](https://github.com/voedger/voedger/issues/455), it is necessary to investigate the performance of various implementations of the cache.

## Cache implementation to investigate

1. [Ristretto, 4.8K stars](https://github.com/dgraph-io/ristretto)
1. [hashicorp/golang-lru, 3.6K starts](https://github.com/hashicorp/golang-lru)
1. [gcache, 2.4K stars](https://github.com/bluele/gcache)
1. [theine-go, 103 stars](https://github.com/Yiling-J/theine-go)
1. [imcache, 66 stars](https://github.com/erni27/imcache)
1. [floatdrop, 19 stars](https://github.com/floatdrop/lru)

- theine-go seems as a technical leader, but not popular (yet?)
- theine-go [claims](https://github.com/dgraph-io/ristretto/issues/336) that Ristretto hits ration is very low (???). Strange that the Ristretto Team has not answered it yet
- Ristretto does NOT provide generic interface
- theine-go and hashicorp/golang-lru DOES provide generic interface
- hashicorp/golang-lru seems as a preferable solution (generic + popular)
  - lru, since it is not clear how to handle evict event in mru version

## Proposal

An exploratory object cache prototype is created in a separate `voedger/inv-go` repository. A prototype, compared to a fully functional cache, should not have a mechanism for automating references and destroying objects, but should provide easy switching of implementation between different candidates.

The following quantitative and qualitative characteristics are subject to assessment:

- **support for generics**
- peak memory usage
- time of a single operation to store the value in the cache:
  - with sequential access
  - with parallel access
- time of a single operation to get the value from the cache
  - with sequential access
  - with parallel access.

## Technical Design

`objcache` interface

```golang
func New[Key comparable, Value any](size int, onEvict func (K, V)) ICache[K, V]

type ICache[Key comparable, Value any] interface {
  Put(K, V)
  Get(K) (value V, ok bool)
}
```

## Results

- [bench.md](results/bench.md)
- [General.pfd](results/General.pdf)
- [Parallel.pfd](results/Parallel.pdf)

## Conclusions

- Ristretto is not suitable as it does not support generics
- Theine-go is not suitable, as it shows unacceptable peak memory usage
- the three remaining candidates (hashicorp, imcache and floatdrop) show approximately the same results, but hashicorp should be preferred for the following reasons:

1. Great popularity (3.6K stars versus 66 and 19 respectively)
2. Slightly better cache stacking speed (~10-15%)

Choice: **[hashicorp/golang-lru, 3.6K starts](https://github.com/hashicorp/golang-lru)**
