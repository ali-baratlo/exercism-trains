package isogram

import "strings"

func IsIsogram(word string) bool {
    
    word1 := strings.ToLower(word)
    word2 := strings.ReplaceAll(word1, " ", "")
	word3 := strings.ReplaceAll(word2, "-", "")
    
    for i:=0 ; i<len(word3) ; i++{
    	for j:=i+1 ; j<len(word3) ; j++{
    		if word3[i]==word3[j]{
                return false
            }
        }
    }
    return true
}
