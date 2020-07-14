package day2

import (
	"sync"
	"testing"
)

/**
就是hashmap
 */
func TestMap(t *testing.T) {
	cache := make(map[string]int)
	cache["a"] = 1
	cache["b"] = 2
	t.Log(cache)

	mapping := map[string]string{"k1":"v1"}
	mapping["k2"] = "v2"
	mapping["k3"] = "v3"
	mapping["k1"] = "1v" // 覆盖

	delete(mapping, "k3") // 删除
	t.Log(len(mapping), mapping["k2"])
	for key,value := range mapping{
		t.Log(key, "-", value)
	}

	// 线程安全的map
	var safeMap sync.Map
	safeMap.Store("k1", "v1")
	safeMap.Store("k2", "v2")

	if value, ok := safeMap.Load("k3"); ok {
		t.Log("loaded: ", value)
	}else {
		t.Log("load failed")
	}

}
