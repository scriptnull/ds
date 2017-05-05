package primitive

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringLength(t *testing.T) {

	// normalStr has 3 bytes
	if len(normalStr) != 3 {
		t.Error("Expected 3, got ", len(normalStr))
	}

	// specialStr has 3 bytes for abc and another 3 bytes for the unicode character at the end
	if len(specialStr) != 6 {
		t.Error("Expected 6, got ", len(specialStr))
	}
}

func TestStringLooping(t *testing.T) {

	for i := 0; i < len(normalStr); i++ {
		fmt.Printf("%d is present at index %d\n", normalStr[i], i)
		// type will be uint8 - byte is an alias for uint8
		fmt.Printf("type of element is %s\n", reflect.TypeOf(normalStr[i]))
	}

	var commonIterate = 0

	for i := 0; i < len(specialStr); i++ {
		fmt.Printf("%d is present at index %d\n", specialStr[i], i)
		commonIterate++
	}

	for index, runeValue := range normalStr {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
		// type will be int32 - rune is an alias for int32
		// range keyword actually iterates the string, giving out a code point ( rune ) on each iteration
		// So, it might be good to use range when we want to iterate over the actual characters, instead of bytes
		fmt.Printf("type of element is %s\n", reflect.TypeOf(runeValue))
	}

	var rangeIterate = 0

	for index, runeValue := range specialStr {
		// iterating using range will iterate over runes ( codepoints )
		// whereas iterating with normal for loop will iterate over bytes
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
		rangeIterate++
	}

	if commonIterate != 6 {
		t.Error("Expected 6, got ", commonIterate)
	}

	if rangeIterate != 4 {
		t.Error("Expected 4, got ", rangeIterate)
	}
}
