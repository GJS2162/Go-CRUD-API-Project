package main

import "os"

const DbName = "inventory"
const DbUser = "root"

var DbPassword = os.Getenv("DB_PASSWORD")
