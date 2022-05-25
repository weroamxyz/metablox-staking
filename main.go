package main

import (
	"fmt"

	foundationContract "github.com/MetaBloxIO/metablox-foundation-services/contract"
	"github.com/metabloxStaking/contract"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/foundationdao"
	"github.com/metabloxStaking/routers"
	"github.com/metabloxStaking/settings"
)

func main() {
	err := settings.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = contract.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = dao.InitSql()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = foundationdao.InitSql()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = foundationContract.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	routers.Setup()
}
