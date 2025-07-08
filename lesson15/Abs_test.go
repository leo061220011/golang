// package main

// import "testing"

//	func TestSum(t *testing.T) {
//		if sum := Sum(1, 2); sum != 3 {
//			t.Errorf("Я ожидал %v,  а получил %v", 3, sum)
//		}
//	}

// package main

// import "testing"

// func TestSum(t *testing.T) {
// 	if sum := Sum(2, 2); sum != 3 {
// 		t.Logf("Привет, я провален")
// 		t.Fatalf("Я ожидал %v,  а получил %v", 3, sum)
// 	}
// 	t.Logf("Привет, все хорошо!")
// }

// package main

// import "testing"

// func TestSum(t *testing.T) {
// 	tests := []struct { // добавляем слайс тестов
// 		name   string
// 		values []int
// 		want   int
// 	}{
// 		{
// 			name:   "simple test #1", // описываем каждый тест:
// 			values: []int{1, 2},      // значения, которые будет принимать функция,
// 			want:   3,                // и ожидаемый результат
// 		},
// 		{
// 			name:   "one",
// 			values: []int{1},
// 			want:   1,
// 		},
// 		{
// 			name:   "with negative values",
// 			values: []int{-1, -2, 3},
// 			want:   0,
// 		},
// 		{
// 			name:   "with negative zero",
// 			values: []int{-0, 3},
// 			want:   3,
// 		},
// 		{
// 			name: "a lot of values",
// 			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
// 				14, 15, 16, 17, 18, 18},
// 			want: 189,
// 		},
// 	}
// 	for _, test := range tests { // цикл по всем тестам
// 		t.Run(test.name, func(t *testing.T) {
// 			if sum := Sum(test.values...); sum != test.want {
// 				t.Errorf("Sum() = %d, want %d", sum, test.want)
// 			}
// 		})
// 	}
// }

// package main

// import "testing"

//	func TestSum(t *testing.T) {
//		if sum := Sum(1, 2); sum != 3 {
//			t.Errorf("Я ожидал %v,  а получил %v", 3, sum)
//		}
//	}

// package main

// import "testing"

// func TestAbs(t *testing.T) {
// 	tests := []struct { // добавляем слайс тестов
// 		name  string
// 		value float64
// 		want  float64
// 	}{
// 		{
// 			name:  "simple test #1", // описываем каждый тест:
// 			value: -1,               // значения, которые будет принимать функция,
// 			want:  1,                // и ожидаемый результат
// 		},
// 		{
// 			name:  "zero",
// 			value: -0,
// 			want:  0,
// 		},
// 		{
// 			name:  "Pi",
// 			value: -3.14,
// 			want:  3.14,
// 		},
// 	}
// 	for _, test := range tests { // цикл по всем тестам
// 		t.Run(test.name,
// 			func(t *testing.T) {
// 				if got := Abs(test.value); got != test.want {
// 					t.Errorf("Abs(%v) = %v, want %v", test.value, got, test.want)
// 				}
// 			})
// 	}
// }

// package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// /*
// 	assert — содержит функции, которые проверяют выполнение условий, сравнивают числа, строки и более сложные объекты (JSON, YAML). Эти функции возвращают значение типа bool. Если проверяемое условие не выполнено, тест выдаёт ошибку, но продолжает свою работу.

// require — содержит тот же набор функций, но они не возвращают значения. Если проверка не пройдена, работа теста прекращается.
// */
// func TestAbs(t *testing.T) {
// 	assert.NotEmpty(t)
// }

// package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// /*
// 	assert — содержит функции, которые проверяют выполнение условий, сравнивают числа, строки и более сложные объекты (JSON, YAML). Эти функции возвращают значение типа bool. Если проверяемое условие не выполнено, тест выдаёт ошибку, но продолжает свою работу.

// require — содержит тот же набор функций, но они не возвращают значения. Если проверка не пройдена, работа теста прекращается.
// */
// func TestAbs(t *testing.T) {
// 	assert.Equal(t, 8.0, Abs(-8))
// 	assert.True(t, true, retTrue())
// 	assert.ElementsMatch(t, []int{1, 2, 3, 4}, []int{2, 3, 4, 1})
// }

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 8.0, Abs(-8))
}
func TestRetTrue(t *testing.T) {
	assert.True(t, true, RetTrue())
}

//go test -coverprofile=cover.out

//go tool cover -html=cover.out
