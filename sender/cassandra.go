package Sendr

import (
	"fmt"
	"log"

	cassandra "github.com/gocql/gocql"
)

//Cassandra :
type Cassandra struct {
	ServerList []string
	Session    *cassandra.Session
}

//NewCassandra :
func NewCassandra(dbConfig DBConfig) *Cassandra {
	var err error
	c := new(Cassandra)
	cluster := cassandra.NewCluster(dbConfig.ServerList[0])
	cluster.Keyspace = dbConfig.KeySpace
	cluster.Consistency = cassandra.Quorum
	c.Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("fail:", err)
		return nil
	}
	return c
}

//InsertKV :
func (c *Cassandra) InsertKV(table, key, value string) error {
	qStr := fmt.Sprintf("INSERT INTO %s (uuid, key, value) VALUES (?, ?, ?)", table)
	if err := c.Session.Query(qStr, cassandra.TimeUUID(), key, value).Exec(); err != nil {
		log.Fatal(err)
	}
	return nil
}
