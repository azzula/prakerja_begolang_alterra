package bilKelipatan7
import "fmt"

func bilKelipatan7() {
	for bil := 0; bil <= 70; bil++ {
		if bil % 7 == 0 {
			fmt.Println(bil)
		}
	}
}