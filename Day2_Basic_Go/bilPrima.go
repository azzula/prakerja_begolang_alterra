package bilPrima
import "fmt"

func bilPrima() {
	for bil := 1; bil <= 10; bil++ {
		i := 0
		for bagi := 1; bagi <= 10; bagi++ {
			if bil % bagi == 0 {
				i++
			}
		}
		if (i == 2) && (bil != 1) {
			fmt.Println(bil)
		}
	}
}