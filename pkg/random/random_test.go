package random

import (
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/database"
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/models"
	"testing"
)

func TestPseudorandomWordSequence_NextWord(t *testing.T) {
	r := New(0)
	for {
		word, i := r.NextWord()
		println(word, i)
		err := database.Save(&models.PseudorandomWord{0, i, word})
		if err != nil {
			panic(err)
		}
	}
}
