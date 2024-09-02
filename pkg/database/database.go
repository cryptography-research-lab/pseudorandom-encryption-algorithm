package database

import (
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/models"
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
	"github.com/syndtr/goleveldb/leveldb"
)

const DatabasePath = "pseudorandom-word-database/"

var database *leveldb.DB

func init() {
	path, err := project_root_directory.GetRootFilePath(DatabasePath)
	if err != nil {
		panic(err)
	}
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		panic(err)
	}
	database = db
}

// Save 保存伪随机序列上的一个单词
func Save(word *models.PseudorandomWord) error {
	return database.Put(word.GetDatabaseKey(), word.GetDatabaseValue(), nil)
}

// Read 根据单词读取其在伪随机序列上的位置
func Read(word string) (*models.PseudorandomWord, error) {
	value, err := database.Get([]byte(word), nil)
	if err != nil {
		return nil, err
	}
	pseudorandomWord := &models.PseudorandomWord{
		Text: word,
	}
	err = pseudorandomWord.SetDatabaseValue(value)
	if err != nil {
		return nil, err
	}
	return pseudorandomWord, nil
}

func ListAll() []*models.PseudorandomWord {
	list := make([]*models.PseudorandomWord, 0)
	iter := database.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		list = append(list, models.NewPseudorandomWordFromDatabaseKeyValue(key, value))
	}
	return list

}
