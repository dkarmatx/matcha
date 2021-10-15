package install

import (
	"matcha/config"
	"strings"
)

func getSqlParamReplacer() *strings.Replacer {
	return strings.NewReplacer(
		"[SCHEMA_NAME]", config.GetDBConfig().Schema,
		"[SCHEMA_OWNER]", config.GetDBConfig().User,
	)
}
