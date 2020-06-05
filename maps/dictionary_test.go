package maps

import (
	"testing"
)

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("this word already exists")
	ErrWordDoesNotExist = DictionaryErr("this word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// dont need pointer to change map，map是引用类型，不能初始化一个空的map
// 但是可以通过以下两种方式创建空的map  dic=map[string]string{}   dic=make(map([string]string)

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
func (d Dictionary) Delete(word string) {
	delete(d, word) //map内置函数
}
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})

}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}
func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}
func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected '%s' to be deleted", word)
	}
}

//summary
// 创建map：var dic=map[string]string{}   dic:=make(map([string]string)
// 添加/更新新值 dic[key]=value
// 查询 value,found=dic[key]  found is bool
// 删除 delete(dic,word)
