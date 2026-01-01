package main

import (
	"fmt"
	"sort"
)

// map
var baseMap map[string]bool

func InsertKey(key string, value bool) {
	baseMap[key] = value
}

func DeleteKey(key string) {
	delete(baseMap, key)
}

// Get
func GetValue(key string) (bool, bool) {
	value, ok := baseMap[key]
	return value, ok
}

// Loop across keys
func iterateKeys() []string {
	keys := make([]string, len(baseMap))
	for k := range baseMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k)
	}
	return keys
}

// loop across values

// loop across range
