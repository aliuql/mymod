package advanced

func Multiply(a, b int) int {
    return a * b
}

func Power(base, exponent int) int {
    result := 1
    for i := 0; i < exponent; i++ {
        result *= base
    }
    return result
}
