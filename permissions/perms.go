package permissions

import (
	"github.com/plally/dgcommand"
	"strings"
)

type PermStorage interface {
	Set(snowflake string, perm string, set bool)
	Get(snowflake string, perm string, set bool)
}

func SetPermCommand(p PermStorage) *dgcommand.Command {
	return dgcommand.NewCommand("setperm <snowflake> <perm> <set>", func(ctx dgcommand.Context) {
		objId := normalizeSnowflake( ctx.Args[0] )
		p.Set(objId, ctx.Args[1], ctx.Args[2] == "true")
	})
}

func GetPermCommand(p PermStorage) *dgcommand.Command {
	return dgcommand.NewCommand("setperm <snowflake> <perm> <set>", func(ctx dgcommand.Context) {
		objId := normalizeSnowflake( ctx.Args[0] )
		p.Set(objId, ctx.Args[1], ctx.Args[2] == "true")
	})
}

func normalizeSnowflake(id string) string {
	objID := strings.ReplaceAll(id, "!", "")
	objID = strings.ReplaceAll(objID, "<", "")
	objID = strings.ReplaceAll(objID, ">", "")
	objID = strings.ReplaceAll(objID, "@", "")

	return id
}