package main
import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

var sign = false

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func feedArray(a [9][9] int)[9][9] int  {
	a[0] = Random()
	return a
}


func Random() [9] int {
	var arr = [9]int{1,2,3,4,5,6,7,8,9}
	rand.Seed(time.Now().Unix())
	for i := len(arr) - 1; i >= 0; i-- {
		num := rand.Intn(len(arr))
		arr[i], arr[num] = arr[num], arr[i]
	}
 
	return arr
}

func printSudoku(resolve [9][9]int)  {
	fmt.Println("-------------------------------")
	for i := 0; i < 9; i++ {
		var v = resolve[i]
		var str string = ""
		str = "|" + str
		for n :=0; n < 9; n++ {
			s := strconv.Itoa(v[n])
			str = str + " " + s + " "
			if (n + 1)%3 == 0 {
				str = str + "|"
			}
		}
        fmt.Println(str)
		if (i + 1)%3 == 0 {
			fmt.Println("-------------------------------")
		}
    }
}

func sudoku()  {
	var array [9] [9] int
	var Sudoku = feedArray(array)
	DFS(0, Sudoku)
	// Output(Sudoku)
}

func Output(sudoku [9][9]int) {
    for i := 0; i < 81; i++ {
        fmt.Printf("%2d ", sudoku[i/9][i%9])
        if i != 0 && (i+1)%9 == 0 {
            fmt.Println("")
        }
    }
}

func Check(n int, key int, _Sudoku [9][9] int) bool {
    for i := 0; i < 9; i++ {
        r := n / 9
        if _Sudoku[r][i] == key {
            return false
        }
    }

    for i := 0; i < 9; i++ {
        c := n % 9
        if _Sudoku[i][c] == key {
            return false
        }
    }

    x := n / 9 / 3 * 3
    y := n % 9 / 3 * 3

    for i := x; i < x+3; i++ {
        for j := y; j < y+3; j++ {
            if _Sudoku[i][j] == key {
                return false
            }
        }
    }

    return true
}

func DFS(n int, _Sudoku [9][9] int) {
    if n > 80 {
        sign = true
		printSudoku(_Sudoku)
        return
    }

    if _Sudoku[n/9][n%9] != 0 {
        DFS(n + 1, _Sudoku)
    } else {
        for v := 1; v <= 9; v++ {
            if Check(n, v, _Sudoku) {
                _Sudoku[n/9][n%9] = v
                DFS(n + 1, _Sudoku)
                if sign {
                    return
                }
                _Sudoku[n/9][n%9] = 0
            }
        }
    }
}

func main() {
	sudoku()
}