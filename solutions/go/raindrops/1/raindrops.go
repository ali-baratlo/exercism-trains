package raindrops

import "fmt"

func Convert(number int) string {
	var result string
    
    if number % 3 == 0{
    	result = result + "Pling"
    }
    if number % 5 == 0{
    	result = result + "Plang"
    }
    if number % 7 == 0{
    	result = result + "Plong"
    }
    if result == ""{
    	 return fmt.Sprintf("%d", number)
    }
    
    return result
}
    
