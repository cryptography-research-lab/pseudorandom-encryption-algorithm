package random

import (
	"testing"

	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/database"
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/models"
)

func TestPseudorandomWordSequence_NextWord(t *testing.T) {
	r := New(0)
	for {
		word, i := r.NextWord(10)
		println(word, i)
		err := database.Save(&models.PseudorandomWord{Index: i, Text: word})
		if err != nil {
			panic(err)
		}
	}
}
