package main

import "testing"

type TestCase struct {
	//InputData string // то что будут на даваться вход
	Answer   Word // то что вернёт тестируемая функция
	Expected Word // то что ожидаем получить
}

var InputData = "newWord.txt"

var r = Word{
	english: "main",
	russian: "Главный",
}
var d = Word{
	english: "wait",
	russian: "Белый",
}

var cases []TestCase = []TestCase{
	{
		//InputData: "newWord.txt",
		Expected: r,
	},
	{
		//InputData: "newWord.txt",
		Expected: d,
	},
}

func TestTakeTXT(t *testing.T) {
	answer := TakeTXT(InputData)
	for id, test := range cases {
		if test.Expected != *answer[id] {
			t.Errorf("test case %d failed: input %v! result %v expected %v", id, InputData, test.Answer, test.Expected)
		}
	}
}
