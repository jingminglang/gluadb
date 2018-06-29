package main

import (
	"github.com/BixData/gluasocket"
    "github.com/BixData/gluabit32"
	"github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	gluasocket.Preload(L)
    gluabit32.Preload(L)
	defer L.Close()
	if err := L.DoString(`

    local mysql = require "db.mysql"
    local db = mysql:new()

    local ok, err, errcode, sqlstate = db:connect{
        host = "127.0.0.1",
        port = 3306,
        database = "mysql",
        user = "root",
        password = "",
        charset = "utf8",
        max_packet_size = 1024 * 1024,
    }

    local res, err, errcode, sqlstate =
        db:query("select * from db", 10)

    print(#res)
    print(res[1].Host)

    db:close()

    local redis = require "db.redis"
    local red = redis:new()
    local ok, err = red:connect("127.0.0.1", 6379)
    ok, err = red:set("dog", "an animal")
    local res, err = red:get("dog")
    print(res)

    `); err != nil {
		panic(err)
	}
}
