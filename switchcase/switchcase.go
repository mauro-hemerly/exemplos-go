package switchcase

import (
	"fmt"
	"time"
)

func Switchcase() {

	i := 2
	fmt.Print("Escrever ", i, " como ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("tres")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Fim de semana...")
	default:
		fmt.Println("Dia de semana...")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Período da manhã...")
	default:
		fmt.Println("Período da tarde...")
	}

}
