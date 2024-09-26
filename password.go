package upgen

import (
	"math"
	"math/rand"
	"strings"
)

func GetPassword(requirements map[int]interface{}, min_entropy int) string {

	H := 0.0 // entropy of password
	cs := make(map[string]interface{}, 0)
	pwd := make([]string, 0)

	// choose a character from each set specified
	for req := range requirements {
		switch req {
		case WordBIP39:
			for _, word := range Bip39List {
				cs[word] = nil
			}
			pwd = append(pwd, Bip39List[rand.Int()%len(Bip39List)])
		default:
			ncs := GetCharacters(req, true)
			for _, c := range ncs {
				cs[string(c)] = nil
			}
			pwd = append(pwd, string(ncs[rand.Int()%len(ncs)]))
			H += math.Log2(float64(len(ncs)))
		}
	}

	// add additional characters until entropy requirement is met
	keys := make([]string, 0, len(cs))
	for k := range cs {
		keys = append(keys, k)
	}
	for H < float64(min_entropy) {
		pwd = append(pwd, string(keys[rand.Int()%len(keys)]))
		H += math.Log2(float64(len(keys)))
	}

	rand.Shuffle(len(pwd), func(i, j int) {
		pwd[i], pwd[j] = pwd[j], pwd[i]
	})

	return strings.TrimSpace(strings.ReplaceAll(strings.Join(pwd, ""), "  ", " "))
}
