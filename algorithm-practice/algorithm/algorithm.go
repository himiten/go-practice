package algorithm
import (
	"errors"
	//"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	//"strconv"
)

func PrimeDesnity(x complex128) float64 {
	v := cmplx.Log(x)
	return real(v) / real(x)
}

func Riemann(x complex128, ln int) (res complex128) {
	for i := 1; i <= ln; i++ {
		res += (1 / cmplx.Pow(complex(float64(i), 0), x))
	}
	return
}

func QuickSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	mid, i := 0, 0
	j := arrLen - 1
	k := arr[i]

	for i != j {
		for ; j > mid; j-- {
			if arr[j] < k {
				arr[j], arr[mid] = arr[mid], arr[j]
				mid = j
				break
			}
		}
		for ; i < mid; i++ {
			if arr[i] > k {
				arr[i], arr[mid] = arr[mid], arr[i]
				mid = i
				break
			}
		}
	}
	QuickSort(arr[:mid])
	QuickSort(arr[(mid + 1):])
}

func AbsVector(sl []float64) (res float64) {
	for _, v := range sl {
		res += (v * v)
	}
	res = math.Sqrt(res)
	return
}

func VecInnerProduct(sl1, sl2 []float64) (res float64, err error) {
	res = 0
	if len(sl1) != len(sl2) {
		err = errors.New("两个向量参数维度不一致")
		return
	}
	for i := 0; i < len(sl1); i++ {
		res += (sl1[i] * sl2[i])
	}
	err = nil
	return
}

func Fibonacci(n int) (sl []int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		sl = append(sl, a)
		a, b = b, a+b
	}
	return
}

func MergeSort(arr []int) []int {
	arrLen := len(arr)
	n := int(math.Ceil(math.Log2(float64(arrLen))))
	for i := 0; i < n; i++ {
		if i > 0 {
			itemLen := int(math.Pow(2, float64(i)))
			items := int(math.Ceil(float64(arrLen) / math.Pow(2, float64(i))))
			for k := 0; k < items; k++ {
				if k%2 == 1 {
					sl1 := arr[(k-1)*itemLen : k*itemLen]
					end := (k + 1) * itemLen
					if arrLen <= (k+1)*itemLen {
						end = arrLen
					}
					sl2 := arr[k*itemLen : end]
					mergeLen := len(sl1) + len(sl2)
					tmpSl := make([]int, mergeLen)
					m, j := 0, 0
					for p := 0; p < mergeLen; p++ {
						if j < len(sl2) && m < len(sl1) {
							if sl1[m] > sl2[j] {
								copy(tmpSl[p:p+1], sl2[j:j+1])
								j++
							} else {
								copy(tmpSl[p:p+1], sl1[m:m+1])
								m++
							}
						} else if j == len(sl2) && m < len(sl1) {
							copy(tmpSl[p:p+1], sl1[m:m+1])
							m++
						} else if j < len(sl2) && m == len(sl1) {
							copy(tmpSl[p:p+1], sl2[j:j+1])
							j++
						}
					}
					copy(arr[(k-1)*itemLen:(k-1)*itemLen+mergeLen], tmpSl)
				}
			}
		} else {
			for j := 0; j < arrLen; j++ {
				if j%2 == 1 {
					if arr[j-1] > arr[j] {
						arr[j-1], arr[j] = arr[j], arr[j-1]
					}
				}
			}
		}
	}
	return arr
}

func GenMatrix(rows, cols int) [][]float64 {
	matrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = float64(rand.Int31n(10))/10 + float64(rand.Int31n(10))
		}
	}
	return matrix
}

func MulMatrix(m1, m2 [][]float64) [][]float64 {
	if len(m1[0]) != len(m2) {
		return nil
	} else {
		m := make([][]float64, len(m1))
		for i := 0; i < len(m1); i++ {
			m[i] = make([]float64, len(m2[0]))
		}
		for j := 0; j < len(m1); j++ {
			for k := 0; k < len(m2[0]); k++ {
				for p := 0; p < len(m2); p++ {
					m[j][k] += m1[j][p] * m2[p][k]
				}
			}
		}
		return m
	}
}

func Transpose(m [][]float64) [][]float64 {
	matrix := make([][]float64, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		matrix[i] = make([]float64, len(m))
		for j := 0; j < len(m); j++ {
			matrix[i][j] = m[j][i]
		}
	}
	return matrix
}