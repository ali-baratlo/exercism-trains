package isogram

import "strings"

func IsIsogram(word1 string) bool {
    
    word := strings.ToLower(word1)

    for i , c := range word {
        if c == ' ' || c == '-'{
            continue
        }

        for j:=i+1 ; j < len(word) ; j++{
            if word[i]==word[j] {
                return false
            }
        }
    }
    return true
}
