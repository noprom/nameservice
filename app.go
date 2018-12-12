package app

import (
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/x/auth"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "nameservice"
)

type nameserviceApp struct {
	*bam.BaseApp
}

func NewnameserviceApp(logger log.Logger, db dbm.DB) *nameserviceApp {
	cdc := MakeCodec()
  bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))
  
  var app = &nameserviceApp{
    BaseApp: bApp,
    cdc: cdc
  }

  return app
}


