package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"shopee/usermanager/userdbproxyserver/g"
	pbdb "shopee/usermanager/userdbproxyserver/proto"
	"log"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("mysql", g.Config().MysqlMaster.Database)
	if err != nil {
		log.Fatalln("open db fail:", err)
	}

	DB.SetMaxOpenConns(g.Config().MysqlMaster.MaxConns)
	DB.SetMaxIdleConns(g.Config().MysqlMaster.MaxIdle)

	err = DB.Ping()
	if err != nil {
		log.Fatalln("ping db fail:", err)
	}

	log.Println("open db success")
}

func getTableNum(username string) int {
	//return 0  // for test
    	var value int
    	for _, c := range []rune(username) {
        	value = value + int(c)
    	}

	return value % 10
}

func AddDBUser(userInfo *pbdb.AddDBUserRequest) error {
	sql := fmt.Sprintf("insert into t_user_%d(user_name, password, nick_name, user_picture, create_time) values ('%s', '%s', '%s', '%s', %d)",
			getTableNum(userInfo.Name),
			userInfo.Name,
			userInfo.Password,
			userInfo.Nick,
			userInfo.Picture,
			userInfo.Regtime,
		)

	log.Println("AddDBUser|sql|", sql)
	_, err := DB.Exec(sql)
	if err != nil {
		log.Println("exec", sql, "fail", err)
	}

	return err
}

func UpdateDBUser(userInfo *pbdb.UpdateDBUserRequest) error {
	sql := fmt.Sprintf("update t_user_%d set nick_name='%s', user_picture='%s' where user_name='%s' ",
			getTableNum(userInfo.Name),
			userInfo.Nick,
			userInfo.Picture,
			userInfo.Name,
		)

	log.Println("UpdateDBUser|sql|", sql)
	_, err := DB.Exec(sql)
	if err != nil {
		log.Println("exec", sql, "fail", err)
	}

	return err
}

func GetDBUser(name string) (*pbdb.GetDBUserReply, error) {
	ret := pbdb.GetDBUserReply{Errmsg: "ok", Retcode: 1}
	sql := fmt.Sprintf("select * from t_user_%d where user_name='%s' ",
			getTableNum(name),
			name,
		)

	log.Println("GetDBUser|sql|", sql)
	rows, err := DB.Query(sql)
	if err != nil {
		log.Println("ERROR|GetDBUser|", err)
		return &ret, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&ret.Name,
			&ret.Password,
			&ret.Nick,
			&ret.Picture,
			&ret.Regtime,
		)

		if err != nil {
			log.Println("WARN|GetDBUser|", err)
			ret.Retcode = 2
			ret.Errmsg = "error"
			return &ret, err
		}

		break;
	}

	return &ret, nil
}
