package hashmap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const alphas = "abcdefghijklmnopqrstuvwxyz"

var hashMap *HashMap

var randKeys, randVals []string
var size = 100

func init() {
	hashMap = NewHashMap()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		key := ""
		value := ""
		for j := random(); j >= 0; j-- {
			index := random()
			index1 := len(alphas) - 1 - index
			key += fmt.Sprint(alphas[index:index+1], i, j)
			value += fmt.Sprint(alphas[index1:index1+1], i, j)
		}
		randKeys = append(randKeys, key)
		randVals = append(randVals, value)
	}
}

func random() int {
	return rand.Intn(len(alphas) - 1)
}

func TestAll(t *testing.T) {
	//test put
	for i := 0; i < size; i++ {
		hashMap.Put(randKeys[i], randVals[i])
	}
	if hashMap.Size() != size {
		t.Error("map size is", hashMap.Size(), "expected", size)
	}
	//test get
	for i := 0; i < size; i++ {
		if hashMap.Get(randKeys[i]) != randVals[i] {
			t.Error("got", hashMap.Get(randKeys[i]), "as value for", randKeys[i], "expected", randVals[i])
		}
	}
	//test overwrite, replace all entries with the values of the keys
	for i := 0; i < size; i++ {
		hashMap.Put(randKeys[i], randKeys[i])
	}
	if hashMap.Size() != size {
		t.Error("map size is", hashMap.Size(), "expected", size)
	}
	for i := 0; i < size; i++ {
		if hashMap.Get(randKeys[i]) != randKeys[i] {
			t.Error("got", hashMap.Get(randKeys[i]), "as value for", randKeys[i], "expected", randKeys[i])
		}
	}
	//test delete
	for i := 0; i < size; i++ {
		hashMap.Delete(randKeys[i])
		if hashMap.Get(randKeys[i]) != nil {
			t.Error("expected nil, value for", randKeys[i], "has been deleted")
		}
	}
	if hashMap.Size() != 0 {
		t.Error("Hashmap not empty after deleting every elements", hashMap)
	}
}
