package random

import "math/rand"

type PseudorandomWordSequence struct {
	R     *rand.Rand
	Seed  int64
	Index int64
}

func New(seed int64) *PseudorandomWordSequence {
	return &PseudorandomWordSequence{
		Seed:  seed,
		R:     rand.New(rand.NewSource(seed)),
		Index: 0,
	}
}

func (x *PseudorandomWordSequence) NextWord() (string, int64) {

	runes := make([]rune, 0)

	// 清空掉前置的0
	for {
		c := rune(x.R.Uint32() % 26)
		if c != 0 {
			runes = append(runes, 'a'+c)
			break
		}
	}

	// 然后开始读取直到遇到0
	for {
		c := rune(x.R.Uint32() % 26)
		if c == 0 {
			break
		}
		// TODO 设置一个最大长度限制
		runes = append(runes, 'a' + c)
	}

	index := x.Index
	x.Index++

	return string(runes), index
}

// ------------------------------------------------ ---------------------------------------------------------------------

func FindWordByIndex(seed, index int64) string {
	r := New(seed)
	var word string
	for i := int64(0); i < index; i++ {
		word, _ = r.NextWord()
	}
	return word
}

// ------------------------------------------------ ---------------------------------------------------------------------



