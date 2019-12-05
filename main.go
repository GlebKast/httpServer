package main

import (
	_ "database/sql"
	"github.com/hubaxis/jwt-auth/tests"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	tests.StartDBTests()

}
