package arr

import (
	"math/rand"
	"testing"
)

func TestInsertDelRandom(t *testing.T) {
	r := ConstructorRandomizedSet()
	r.Insert(1)
	r.Insert(2)
	r.Insert(3)
	r.Remove(2)
}

type RandomizedSet struct {
	dict map[int]int
	s    []int
}

func ConstructorRandomizedSet() RandomizedSet {
	d := make(map[int]int)
	s := make([]int, 0)
	return RandomizedSet{dict: d, s: s}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dict[val]
	if ok {
		return false
	}
	this.s = append(this.s, val)
	this.dict[val] = len(this.s) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.dict[val]
	if !ok {
		return false
	}
	//找到切片最后的值
	lastIndex := len(this.s) - 1
	lastValue := this.s[lastIndex]
	//找到val对应的index，把最后的元素进行替换
	this.s[index] = lastValue
	//交换dict数据
	this.dict[lastValue] = index
	delete(this.dict, val)
	this.s = this.s[:lastIndex]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	v := rand.Intn(len(this.s))
	return this.s[v]
}
