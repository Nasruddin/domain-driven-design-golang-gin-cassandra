package cassandra

import (
	"github.com/gocql/gocql"
)

/*var (
	cluster *gocql.ClusterConfig
)
*/

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "user"
	cluster.Consistency = gocql.Quorum
	// return cluster.CreateSession()

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}