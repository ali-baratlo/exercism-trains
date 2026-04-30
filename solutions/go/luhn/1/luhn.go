package luhn


func Valid(id string) bool {
    var sum int
    var count int
    var double bool = false
    
	for i:=len(id) - 1 ; i>=0 ; i-- {
        if id[i] == ' ' {
            continue
        }
        if id[i] < '0' || id[i] > '9' || id[i] == '-' {
            return false
        }
        number := int(id[i]-'0')
        count ++

        if double{
            number *= 2 
            if number > 9 {
                number -= 9 
            }
        }
        sum += number
        double = !double
    }
    if count <= 1 {
        return false
    }
    return sum%10 ==0
}
