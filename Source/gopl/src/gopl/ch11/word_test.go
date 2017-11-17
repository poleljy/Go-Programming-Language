package ch11

import (
	"testing"
	"time"
	"math/rand"
)

// 表格测试
func TestIsPalindrome(t *testing.T) {
	tests := []struct{
		input string
		want bool
	}{
		  {"", true},
        {"a", true},
        {"aa", true},
        {"ab", false},
        {"kayak", true},
        {"detartrated", true},
        {"A man, a plan, a canal: Panama", true},
        {"Evil I did dwell; lewd did I live.", true},
        {"Able was I ere I saw Elba", true},
        {"été", true},
        {"Et se resservir, ivresse reste.", true},
        {"palindrome", false}, // non-palindrome
        {"desserts", false},   // semi-palindrom
	}
	
	for _, test := range tests {
		if ret:= IsPalindrome(test.input); ret != test.want {
			t.Errorf(`IsPalindrome(%q) = %v,want %v`, test.input, ret, test.want)
		}
	}
}

// 随即测试
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	
	rng := rand.New(rand.NewSource(seed))
	for i:= 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
            t.Errorf("IsPalindrome(%q) = false", p)
        }
	}
}
