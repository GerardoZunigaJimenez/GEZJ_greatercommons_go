package word

import (
"testing"
"fmt"
"strings"
)

type test struct {
	paragraph   string
	wordsNumber int
	wordCounter map[string]int
}

var sTest = []test{
	test{"They are so fancy, and they need to be killed", 10, map[string]int{"they": 2, "are": 1, "so": 1, "fancy": 1, "and": 1, "need": 1, "to": 1, "be": 1, "killed": 1,}},
	test{"Bla bla bla bla", 4, map[string]int{"bla": 4,}},
	test{"Sometimes, someTimes, SoMeTiMeS; soMEtiMES. sometimes", 5,  map[string]int{"sometimes":5}},
}

func TestCountWordsByString(t *testing.T) {
	for i := 0; i < len(sTest); i++ {
		count := CountWordsByString(sTest[i].paragraph)
		if count != sTest[i].wordsNumber {
			t.Error("Expedted", sTest[i].wordsNumber, "Got", count)
		}
	}
}

func ExampleCountWordsByString() {
	s :=CountWordsByString("Bla BLA bLa BLA")
	fmt.Println(s)
	// Output:
	// 4
}

func BenchmarkCountWordsByString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(sTest); i++ {
			count := CountWordsByString(sTest[i].paragraph)
			if count != sTest[i].wordsNumber {
				b.Error("Expedted", sTest[i].wordsNumber, "Got", count)
			}
		}
	}
}

func TestUseCountContent(t *testing.T) {
	for i := 0; i < len(sTest); i++ {
		m := UseCount(sTest[i].paragraph)

		//Checking len
		if len(m) != len(sTest[i].wordCounter) {
			t.Error("Procesed Words -> Expected:", len(sTest[i].wordCounter), ", Got:", len(m))
		}

		//Checking Key and values
		for k, v := range sTest[i].wordCounter {
			k = strings.ToLower(k)
			if m[k] != v {
				t.Errorf("Expected %v[%v], Got %v[%v]", k, v, k, m[k])
			}
		}
	}
}

func ExampleUseCount() {
	m := UseCount("Bla bla BLA bla")
	fmt.Println(m)
	// Output:
	// map[bla:4]
}

func BenchmarkUseCount(b *testing.B) {
	for i:=0; i<b.N; i++{
		for i := 0; i < len(sTest); i++ {
			m := UseCount(sTest[i].paragraph)

			//Checking len
			if len(m) != len(sTest[i].wordCounter) {
				b.Error("Procesed Words -> Expected:", len(sTest[i].wordCounter), ", Got:", len(m))
			}

			//Checking Key and values
			for k, v := range sTest[i].wordCounter {
				k = strings.ToLower(k)
				if m[k] != v {
					b.Errorf("Expected %v[%v], Got %v[%v]", k, v, k, m[k])
				}
			}
		}
	}
}

