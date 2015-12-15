package dict

/*
	average insert: O(n)
	average seek: O(n)
	average delete: O(n)
*/

// Entry represents the value that is being saved.
type Entry struct {
	Key   string
	Value string
}

// ArrayDict is an implementation of a dictionary using an array.
type ArrayDict []Entry

func NewArray() *ArrayDict {
	result := ArrayDict([]Entry{})
	return &result
}

func (a *ArrayDict) find(key string) int {
	for i, e := range *a {
		if e.Key == key {
			return i
		}
	}
	return -1
}

// Insert will insert a value into the dictionary.
func (a *ArrayDict) Insert(key, value string) {
	index := a.find(key)
	if index < 0 {
		index = len(*a)
	}
	entry := Entry{Key: key, Value: value}
	*a = append(*a, entry)
	copy((*a)[index+1:], (*a)[index:])
	(*a)[index] = entry
}

// Delete will delete a key from the dictionary.
func (a *ArrayDict) Delete(key string) {
	index := a.find(key)
	if index < 0 {
		return
	}
	copy((*a)[index:], (*a)[index+1:])
	*a = (*a)[:len(*a)-1]
}

// Get will retrieve a value for a given key. An empty string will be returned for non-extant keys.
func (a *ArrayDict) Get(key string) string {
	index := a.find(key)
	if index < 0 {
		return ""
	}
	return (*a)[index].Value
}

// Has will return true if the given key is found in the dictionary.
func (a *ArrayDict) Has(key string) bool {
	return a.find(key) >= 0
}
