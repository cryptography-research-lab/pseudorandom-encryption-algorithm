package find

import (
	"fmt"
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/data"
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/database"
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/models"
	"github.com/cryptography-research-lab/pseudorandom-encryption-algorithm/pkg/random"
	"strings"
)

const SeedIndexLimit = 1000000000

func ReadWordSet() map[string]struct{} {
	lines := strings.Split(data.Words, "\n")
	dict := make(map[string]struct{})
	for _, line := range lines {
		dict[strings.TrimSpace(line)] = struct{}{}
	}
	return dict
}

func Run() {
	for i := 0; i < 10000; i++ {
		RunForSeed(int64(i))
	}
}

func RunForSeed(seed int64) {
	set := ReadWordSet()
	r := random.New(seed)
	for {
		word, index := r.NextWord()

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
		err = database.Save(&models.PseudorandomWord{seed, index, word})
		fmt.Printf("save text %s, seed = %d, index = %d\n", word, seed, index)
		if err != nil {
			panic(err)
		}
	}
}
