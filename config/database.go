package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func (p *config) ConnectDB() *sql.DB {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s sslrootcert=%s sslkey=%s sslcert=%s",
		p.Server.Database.Host,
		p.Server.Database.Port,
		p.Server.Database.User,
		p.Server.Database.Password,
		p.Server.Database.DbName,
		p.Server.Database.SslMode,
		p.Server.Database.SslRootCert,
		p.Server.Database.SslKey,
		p.Server.Database.SslCert,
	)

	db, err := sql.Open(p.Server.Database.Driver, url)

	if err != nil {
		p.error(err)
	}

	if err := db.Ping(); err != nil {
		p.error(err)
	}

	return db
}

func (p *config) error(err error) {
	log.Fatalf("failed to connect to database, %v://%v:****@%v:%v/%v?sslmode=%s;sslrootcert=%s;sslkey=%s;sslcert=%s, error %v",
		p.Server.Database.Driver,
		p.Server.Database.User,
		p.Server.Database.Host,
		p.Server.Database.Port,
		p.Server.Database.DbName,
		p.Server.Database.SslMode,
		p.Server.Database.SslRootCert,
		p.Server.Database.SslKey,
		p.Server.Database.SslCert,
		err.Error())
}
