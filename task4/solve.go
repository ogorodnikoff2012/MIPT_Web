package main

import (
    "strings"
    "unicode"
)

func RemoveEven(slice []int) (ret []int) {
    for i := range slice {
        if slice[i] % 2 != 0 {
            ret = append(ret, slice[i])
        }
    }
    return
}

func PowerGenerator(base int) (func() int) {
    val := 1
    return func() (ret int) {
        val *= base
        ret = val
        return
    }
}

func DifferentWordsCount(text string) int {
    runes := append([]rune(strings.ToLower(text)), rune(' '))
    words := make(map[string]int)
    lastIdx := -1
    for idx, r := range runes {
        if !unicode.IsLetter(r) {
            if idx - lastIdx > 1 {
                words[string(runes[lastIdx + 1:idx])] = idx
            }
            lastIdx = idx
        }
    }
    return len(words)
}
