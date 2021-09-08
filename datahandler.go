package main

import "sync"

//This stores the hashed passwords in int-string key value.
type PasswordMap struct {
	sync.RWMutex
	m map[int]string
}

//Read hash value based on int key
func (r *PasswordMap) Get(key int) string {
	r.RLock()
	defer r.RUnlock()
	return r.m[key]
}

//Add new hash value and return new key
func (r *PasswordMap) Set(val string) int {
	r.Lock()
	defer r.Unlock()
	key := len(r.m) + 1
	r.m[key] = val
	return key
}
