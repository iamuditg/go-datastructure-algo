package psday01

import (
	"fmt"
	"strings"
	"testing"
)

var maxWordsInput map[string]int
var maxWordsOutput map[string][]string

func getMaxWordsInput() map[string]int {
	maxWordsInput = make(map[string]int)
	maxWordsInput["The quick brown fox jumps right over the little lazy dog."] = 2
	return maxWordsInput
}

func getMaxWordsOutput() map[string][]string {
	maxWordsOutput = make(map[string][]string)
	maxWordsOutput["The quick brown fox jumps right over the little lazy dog."] = []string{"quick", "over", "little"}
	return maxWordsOutput
}

func difference(a, b []string) bool {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	if diff != nil {
		return true
	}
	return false
}

func TestFindMaximumNumberOfWords(t *testing.T) {
	maxWordsInput = getMaxWordsInput()
	maxWordsOutput = getMaxWordsOutput()
	for k, v := range maxWordsInput {
		actualResult := FindMaximumNumberOfWords(k, v)
		expectedOutput := maxWordsOutput[k]
		if difference(actualResult, expectedOutput) {
			t.Errorf("error in result: %v with input of: %s", actualResult, k)
		} else {
			t.Logf("expected result passed: %s", k)
		}
	}
}

func getGreatestElementInput() []int {
	return []int{15, 10, 16, 20, 8, 9, 7, 50}
}

func getGreatestElementOutput() map[int]int {
	output := make(map[int]int)
	output[15] = 16
	output[10] = 16
	output[16] = 20
	output[20] = 50
	output[8] = 9
	output[9] = 50
	output[7] = 50
	output[50] = -1
	return output
}

func TestFindGreatestElementArray(t *testing.T) {
	input := getGreatestElementInput()
	actualResult := FindGreatestElementArray(input)
	expectedOutput := getGreatestElementOutput()
	for k, actV := range actualResult {
		if expV, ok := expectedOutput[k]; !ok {
			t.Errorf("error in result: %v and output: %v", actualResult, expectedOutput)
		} else {
			if expV != actV {
				t.Errorf("error in result value: %v and output: %v", actualResult, expectedOutput)
			} else {
				t.Logf("result %v result passed", actualResult)
			}
		}
	}
	fmt.Println(actualResult)
}

func getValidParenthesisInput(i int) string {
	switch i {
	case 0:
		return "{}{}{}"
	case 1:
		return "{}"
	case 2:
		return "{}{}{}"
	}
	return ""
}

func TestCheckValidParenthesis(t *testing.T) {
	for i := 0; i < 3; i++ {
		input := getValidParenthesisInput(i)
		actualResult := CheckValidParenthesis(input)
		if actualResult {
			t.Logf("result is passed: %v with input: %s", actualResult, input)
		} else {
			t.Errorf("error in result with input: %s and output: %v", input, actualResult)
		}
	}
}

func getPalindromeInput() string {
	return "ababd"
}

func TestCheckPalindromeString(t *testing.T) {
	input := getPalindromeInput()
	palindromeString := CheckPalindromeString(input)
	if strings.Contains(palindromeString, "not palindrome") {
		t.Errorf("error in result: %v output: %v", input, palindromeString)
	}
}
