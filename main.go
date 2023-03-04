package main

import (
	_ "openai-proxy-backend/internal/packed"

	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"

	_ "openai-proxy-backend/modules"

	"github.com/gogf/gf/v2/os/gctx"

	"openai-proxy-backend/internal/cmd"
)

func main() {
	// gres.Dump()
	cmd.Main.Run(gctx.New())
}
