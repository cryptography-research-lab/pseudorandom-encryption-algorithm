package find

import (
	"fmt"
	"strings"

	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/data"
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/database"
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/models"
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/random"
)

const SeedIndexLimit = 1000000000

var MaxWordLength = 0

func ReadWordSet() map[string]struct{} {
	lines := strings.Split(data.Words, "\n")
	dict := make(map[string]struct{})
	for _, line := range lines {
		dict[strings.TrimSpace(line)] = struct{}{}
		MaxWordLength = max(MaxWordLength, len(line))
	}
	return dict
}

func Run() {
	set := ReadWordSet()

	for i := 0; i < 10000; i++ {
		RunForSeed(set, int64(i))
	}
}

func RunForSeed(set map[string]struct{}, seed int64) {

	r := random.New(seed)
	for {
		word, index := r.NextWord(MaxWordLength)

		if index > SeedIndexLimit {
			return
		}

		// 不在白名单中的就直接忽略了
		if _, exists := set[word]; !exists {
			continue
		}

		read, err := database.Read(word)
		if err == nil && read.Index <= index {
			continue
		}
		err = database.Save(&models.PseudorandomWord{Seed: seed, Index: index, Text: word})
		fmt.Printf("save text %s, seed = %d, index = %d\n", word, seed, index)
		if err != nil {
			panic(err)
		}
	}
}
