module github.com/oo-pp307/gen/tests

go 1.16

require (
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	gorm.io/driver/mysql v1.5.7
	gorm.io/driver/sqlite v1.4.4
	github.com/oo-pp307/gen v0.3.19
	gorm.io/gorm v1.25.12
	gorm.io/hints v1.1.1 // indirect
	gorm.io/plugin/dbresolver v1.5.3
)

replace github.com/oo-pp307/gen => ../
