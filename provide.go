/*
 * Copyright (c) 2023-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package objcache

import (
	"fmt"
	"objcache/internal/floatdrop"
	"objcache/internal/hashicorp"
	"objcache/internal/imcache"
	"objcache/internal/theine"
)

// Creates and return new LRU object cache with K key type and V value type.
//
// Maximum cache size is limited by size param. Optional onEvicted cb is called then some value evicted from cache.
func New[K comparable, V any](size int, onEvicted func(K, V)) ICache[K, V] {
	return hashicorp.New[K, V](size, onEvicted)
}

type CacheProvider uint8

const (
	Hashicorp CacheProvider = iota
	Theine
	Floatdrop
	Imcache

	CacheProvider_Count
)

func NewProvider[K comparable, V any](p CacheProvider, size int, onEvicted func(K, V)) ICache[K, V] {
	switch p {
	case Hashicorp:
		return hashicorp.New[K, V](size, onEvicted)
	case Theine:
		return theine.New[K, V](size, onEvicted)
	case Floatdrop:
		return floatdrop.New[K, V](size, onEvicted)
	case Imcache:
		return imcache.New[K, V](size, onEvicted)
	}
	panic(fmt.Errorf("unknown cache provider specified %v", p))
}
