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

var passwordChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789&*-+.?!,;:")

func (x *PseudorandomWordSequence) NextWord(maxLen int) (string, int64) {

	runes := make([]rune, 0)

	for {
		// 获取一个随机数
		// 如果这个数是passwordChars的长度,并且runes为空,那么就继续循环,直到runes不为空结束
		c := x.R.Intn(len(passwordChars) + 1)
		if c == len(passwordChars) {
			if len(runes) == 0 {
				continue
			}
			break
		}

		runes = append(runes, passwordChars[c])

		if len(runes) >= maxLen {
			break
		}
	}

	index := x.Index
	x.Index++

	return string(runes), index
}

// ------------------------------------------------ ---------------------------------------------------------------------

func FindWordByIndex(maxLen int, seed, index int64) string {
	r := New(seed)
	var word string
	for i := int64(0); i < index; i++ {
		word, _ = r.NextWord(maxLen)
	}
	return word
}

// ------------------------------------------------ ---------------------------------------------------------------------
