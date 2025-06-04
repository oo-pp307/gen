package main

import (
	"github.com/oo-pp307/gen"
	"github.com/oo-pp307/gen/examples/conf"
	"github.com/oo-pp307/gen/examples/dal"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()

	prepare(dal.DB) // prepare table for generate
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "/tmp/gentest/query",
	})

	g.UseDB(dal.DB)

	g.GenerateAllTable()

	g.Execute()
}
