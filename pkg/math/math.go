package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Ceil 函数用于计算a除以b的向上取整的结果
// 参数a和b的类型为T，满足T为浮点型或者整型
// 返回值为T类型的向上取整的结果
func Ceil[T constraints.Float | constraints.Integer](a, b T) T {
	if b == 0 {
		return 0
	}

	return T(math.Ceil(float64(a) / float64(b)))
}

// Abs 函数用于计算a的绝对值
func Abs[T constraints.Float | constraints.Integer](a T) T {
	return T(math.Abs(float64(a)))
}

// Floor 函数用于计算a除以b的向下取整的结果
func Floor[T constraints.Float | constraints.Integer](a, b T) T {
	if b == 0 {
		return 0
	}

	return T(math.Floor(float64(a) / float64(b)))
}

// Mod 函数用于计算a除以b的余数
func Mod[T constraints.Float | constraints.Integer](a, b T) T {
	if b == 0 {
		return 0
	}

	return T(math.Mod(float64(a), float64(b)))
}

// Pow 函数用于计算a的b次方
func Pow[T constraints.Float | constraints.Integer](a, b T) T {
	return T(math.Pow(float64(a), float64(b)))
}

// Sqrt 函数用于计算a的平方根
func Sqrt[T constraints.Float | constraints.Integer](a T) T {
	return T(math.Sqrt(float64(a)))
}

// Cbrt 函数用于计算a的立方根
func Cbrt[T constraints.Float | constraints.Integer](a T) T {
	return T(math.Cbrt(float64(a)))
}

// Round 函数用于计算a的四舍五入的结果
func Round[T constraints.Float | constraints.Integer](a T) T {
	return T(math.Round(float64(a)))
}

// IsNaN 函数用于判断a是否为NaN
// 参数a的类型为T，满足T为浮点型或者整型
// 返回值为bool类型，表示a是否为NaN
func IsNaN[T constraints.Float | constraints.Integer](a T) bool {
	return math.IsNaN(float64(a))
}

// Trunc 函数用于计算a的向下取整的结果
func Trunc(x float64) float64 {
	return math.Trunc(x)
}

// Max 函数用于计算a和b的最大值
func Max[T constraints.Float | constraints.Integer](a, b T) T {
	return T(math.Max(float64(a), float64(b)))
}

// Min 函数用于计算a和b的最小值
func Min[T constraints.Float | constraints.Integer](a, b T) T {
	return T(math.Min(float64(a), float64(b)))
}

// Dim 函数用于计算a和b的最大值
func Dim[T constraints.Float | constraints.Integer](a, b T) T {
	return T(math.Dim(float64(a), float64(b)))
}

// Hypot 返回Sqrt (p * p q * q)，注意避免不必要的上溢和下溢
func Hypot[T constraints.Float | constraints.Integer](a, b T) T {
	return T(math.Hypot(float64(a), float64(b)))
}
