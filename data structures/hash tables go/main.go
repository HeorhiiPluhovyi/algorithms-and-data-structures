package main

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 8

// HashTable structure will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the HashTable array
type bucket struct {
	head *buckedNode
}

// bucket
type buckedNode struct {
	key  string
	next *buckedNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert will take in a key,
// create a node with the key and insert the node in the bucket
func (b *bucket) insert(key string) {
	if !b.search(key) {
		newBuckedNode := &buckedNode{
			key:  key,
			next: b.head,
		}

		b.head = newBuckedNode
	} else {
		fmt.Println("bucket already exist")
	}

}

// search will take in a key,
// create a node with the key and insert the node in the bucket
func (b *bucket) search(key string) bool {
	curr := b.head
	for curr != nil {
		if curr.key == key {
			return true
		}
	}

	return false
}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previousNode := b.head

	for previousNode.next != nil {
		if previousNode.key == key {
			previousNode.next = previousNode.next.next
			return
		}
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	newHashTable := &HashTable{}

	for i := 0; i < ArraySize; i++ {
		newHashTable.array[i] = &bucket{}
	}

	return newHashTable
}

func main() {
	hm := Init()
	fmt.Println(hm)
	fmt.Println(hash("RANDY"))
	hm.Insert("RANDY")
	fmt.Println(hm.Search("RANDY"))
}
