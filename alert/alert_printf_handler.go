package alert

import (
	"fmt"
	"github.com/DWHengr/aurora/internal/alertcore"
)

func PrintfHandler(message *alertcore.AlertMessage, ctx *alertcore.Context) {
	fmt.Println("----------------------start----------------------")
	fmt.Println(message)
	fmt.Println("-----------------------end-----------------------")
}
