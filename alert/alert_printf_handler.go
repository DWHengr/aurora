package alert

import (
	"fmt"
	"github.com/DWHengr/aurora/internal/alertcore"
)

func PrintfHandler(message *alertcore.AlertMessage) {
	fmt.Println("----------------------start----------------------")
	fmt.Println(message)
	fmt.Println("-----------------------end-----------------------")
}
