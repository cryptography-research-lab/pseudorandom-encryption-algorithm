package models

import (
	"fmt"
	"strconv"
	"strings"
)

const DatabaseValueDelimiter = ","

// PseudorandomWord 伪随机数序列的位置
type PseudorandomWord struct {

	// 随机数序列的种子
	Seed int64 `json:"seed"`

	// 是第多少个单词
	Index int64 `json:"index"`

	// 这个单词是啥
	Text string `json:"text"`
}

// ------------------------------------------------ ---------------------------------------------------------------------

func NewPseudorandomWordFromDatabaseKeyValue(key, value []byte) (*PseudorandomWord, error) {
	x := &PseudorandomWord{}
	x.SetDatabaseKey(key)
	err := x.SetDatabaseValue(value)
	return x, err
}

func (x *PseudorandomWord) GetDatabaseKey() []byte {
	return []byte(x.Text)
}

func (x *PseudorandomWord) SetDatabaseKey(key []byte) {
	x.Text = string(key)
}

func (x *PseudorandomWord) GetDatabaseValue() []byte {
	s := strconv.FormatInt(x.Seed, 10) + DatabaseValueDelimiter + strconv.FormatInt(x.Index, 10)
	return []byte(s)
}

func (x *PseudorandomWord) SetDatabaseValue(value []byte) error {
	split := strings.Split(string(value), DatabaseValueDelimiter)
	if len(split) != 2 {
		return fmt.Errorf("%s invalid", value)
	}
	seed, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		return err
	}
	index, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return err
	}
	x.Seed = seed
	x.Index = index
	return nil
}

func (x *PseudorandomWord) String() string {
	return strconv.FormatInt(x.Seed, 10) + DatabaseValueDelimiter + strconv.FormatInt(x.Index, 10) + DatabaseValueDelimiter + x.Text
}

// ------------------------------------------------ ---------------------------------------------------------------------
