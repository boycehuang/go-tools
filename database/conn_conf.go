package database

import "fmt"

type ConnectionConfig struct {
	Dialect  string
	Filepath string //sqlite only
	Host     string
	Port     string
	User     string
	Pass     string
	Dbname   string
	Ssl      bool
}

func (cc ConnectionConfig) URI() (uri string) {
	switch dialect := cc.Dialect; dialect {
	case MYSQL:
		if cc.Port != "" {
			cc.Port = "3306"
		}
		uri = fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local", cc.User, cc.Pass, cc.Host, cc.Port, cc.Dbname)
	case POSTGRES:
		if cc.Port != "" {
			cc.Port = "5432"
		}
		uri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", cc.Host, cc.Port, cc.User, cc.Dbname, cc.Pass)
	case MSSQL:
		if cc.Port != "" {
			cc.Port = "1443"
		}
		uri = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", cc.User, cc.Pass, cc.Host, cc.Port, cc.Dbname)
	case COCKROACHDB:
		if cc.Port != "" {
			cc.Port = "26257"
		}
		uri = fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable", cc.User, cc.Host, cc.Port, cc.Dbname)
	case SQLITE:
		uri = cc.Filepath
	}

	return
}
