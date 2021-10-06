package main

// MyCircularDeque Leetcode: https://leetcode-cn.com/problems/design-circular-deque/
type MyCircularDeque struct {
	queue    []int
	capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue:    make([]int, 0, k),
		capacity: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if !this.IsFull() {
		this.queue = append([]int{value}, this.queue...)
		return true
	}
	return false
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if !this.IsFull() {
		this.queue = append(this.queue, value)
		return true
	}
	return false
}

func (this *MyCircularDeque) DeleteFront() bool {
	if !this.IsEmpty() {
		this.queue = this.queue[1:]
		return true
	}
	return false
}

func (this *MyCircularDeque) DeleteLast() bool {
	if !this.IsEmpty() {
		this.queue = this.queue[:len(this.queue)-1]
		return true
	}
	return false
}

func (this *MyCircularDeque) GetFront() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}

func (this *MyCircularDeque) GetRear() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[len(this.queue)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	if len(this.queue) == 0 {
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	length := len(this.queue)
	capacity := this.capacity
	if length == capacity {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
