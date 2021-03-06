package trie

import (
	"fmt"
	"io/ioutil"
	"os"
	"seth/common"
	"seth/database"
	"seth/database/leveldb"
	"testing"
)

func newTestTrieDB() (database.Database, func()) {
	dir, err := ioutil.TempDir("", "trietest")
	if err != nil {
		panic(err)
	}
	db, err := leveldb.NewLevelDB(dir, 0, 0)
	if err != nil {
		panic(err)
	}
	return db, func() {
		db.Close()
		os.RemoveAll(dir)
	}
}

func Test_trie_Update(t *testing.T) {
	db, remove := newTestTrieDB()
	defer remove()
	trie, err := NewTrie(common.Hash{}, []byte("trietest"), db)
	if err != nil {
		panic(err)
	}
	trie.Put([]byte("12345678"), []byte("test"))
	trie.Put([]byte("12345678"), []byte("testnew"))
	trie.Put([]byte("12345557"), []byte("test1"))
	trie.Put([]byte("12375879"), []byte("test2"))
	trie.Put([]byte("02375879"), []byte("test3"))
	trie.Put([]byte("04375879"), []byte("test4"))
	trie.Put([]byte("24375879"), []byte("test5"))
	trie.Put([]byte("24375878"), []byte("test6"))
	trie.Put([]byte("24355879"), []byte("test7"))
	value := trie.Get([]byte("12345678"))
	fmt.Println(string(value))
	value = trie.Get([]byte("12345557"))
	fmt.Println(string(value))
	value = trie.Get([]byte("12375879"))
	fmt.Println(string(value))
	value = trie.Get([]byte("02375879"))
	fmt.Println(string(value))
	value = trie.Get([]byte("04375879"))
	fmt.Println(string(value))
	value = trie.Get([]byte("24375879"))
	fmt.Println(string(value))
	value = trie.Get([]byte("24375878"))
	fmt.Println(string(value))
	value = trie.Get([]byte("24355879"))
	fmt.Println(string(value))
	batch := db.NewBatch()
	hash, _ := trie.Commit(batch)
	batch.Commit()
	fmt.Println(hash)
	fmt.Println(trie.Hash())

}

func Test_trie_Delete(t *testing.T) {
	db, remove := newTestTrieDB()
	defer remove()
	trie, err := NewTrie(common.Hash{}, []byte("trietest"), db)
	if err != nil {
		panic(err)
	}
	trie.Put([]byte("12345678"), []byte("test"))
	trie.Put([]byte("12345678"), []byte("testnew"))
	trie.Put([]byte("12345557"), []byte("test1"))
	trie.Put([]byte("12375879"), []byte("test2"))
	trie.Put([]byte("02375879"), []byte("test3"))
	trie.Put([]byte("04375879"), []byte("test4"))
	trie.Put([]byte("24375879"), []byte("test5"))
	trie.Put([]byte("24375878"), []byte("test6"))
	trie.Put([]byte("24355879"), []byte("test7"))
	match := trie.Delete([]byte("12345678"))
	fmt.Println(match)
	match = trie.Delete([]byte("12375879"))
	fmt.Println(match)
	match = trie.Delete([]byte("24375879"))
	fmt.Println(match)
	match = trie.Delete([]byte("24375889"))
	fmt.Println(match)
	value := trie.Get([]byte("12345678"))
	fmt.Println(string(value))
	value = trie.Get([]byte("12375879"))
	fmt.Println(string(value))
	value = trie.Get([]byte("02375879"))
	fmt.Println(string(value))
	value = trie.Get([]byte("04375879"))
	fmt.Println(string(value))
	value = trie.Get([]byte("24375879"))
	fmt.Println(string(value))
	value = trie.Get([]byte("24375878"))
	fmt.Println(string(value))
	value = trie.Get([]byte("24355879"))
	fmt.Println(string(value))
	fmt.Println(trie.Hash())
}

func Test_trie_Commit(t *testing.T) {
	db, remove := newTestTrieDB()
	defer remove()
	trie, err := NewTrie(common.Hash{}, []byte("trietest"), db)
	if err != nil {
		panic(err)
	}
	trie.Put([]byte("12345678"), []byte("test"))
	trie.Put([]byte("12345557"), []byte("test1"))
	trie.Put([]byte("12375879"), []byte("test2"))
	trie.Put([]byte("02375879"), []byte("test3"))
	trie.Put([]byte("04375879"), []byte("test4"))
	trie.Put([]byte("24375879"), []byte("test5"))
	trie.Put([]byte("24375878"), []byte("test6"))
	trie.Put([]byte("24355879"), []byte("test7"))

	batch := db.NewBatch()
	hash, _ := trie.Commit(batch)
	batch.Commit()
	fmt.Println(hash)

	fmt.Println(string("----------------------------------"))
	trienew, err := NewTrie(hash, []byte("trietest"), db)

	trienew.Delete([]byte("24355879"))
	trienew.Put([]byte("243558790"), []byte("test8"))
	trienew.Put([]byte("043758790"), []byte("test9"))

	value := trienew.Get([]byte("12345678"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("12345557"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("12375879"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("02375879"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("04375879"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("24375879"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("24375878"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("243558790"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("043758790"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("12345557"))
	fmt.Println(string(value))
	value = trienew.Get([]byte("12375879"))
	fmt.Println(string(value))
}
