package differenceofsquares

func SquareOfSum(n int) int {
    var sum int
    for i:=0 ; i<=n ; i++ {
        sum=sum+i
    }
	return sum * sum
}

func SumOfSquares(n int) (sum int) {
    for i:=0 ; i<=n ; i++ {
        sum=sum+(i*i)
    }
	return
}

func Difference(n int) int {
    if (SumOfSquares(n) - SquareOfSum(n)) <0 {
        return - (SumOfSquares(n) - SquareOfSum(n))
    }
	return SumOfSquares(n) - SquareOfSum(n) 
}
