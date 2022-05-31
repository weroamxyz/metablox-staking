package main

import (
	"fmt"
	"github.com/metabloxStaking/interest"

	"github.com/go-playground/validator/v10"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/routers"
	"github.com/metabloxStaking/settings"
)

func main() {
	validate := validator.New()

	err := settings.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	/*err = contract.Init()
	if err != nil {
		fmt.Println(err)
		return
	}*/

	err = dao.InitSql(validate)
	if err != nil {
		fmt.Println(err)
		return
	}

	interest.StartHourlyTimer()

	routers.Setup()
}
