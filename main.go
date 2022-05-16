package main

import (
	"fmt"

	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/routers"
	"github.com/metabloxStaking/settings"
)

func main() {
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

	err = dao.InitSql()
	if err != nil {
		fmt.Println(err)
		return
	}

	routers.Setup()
}
