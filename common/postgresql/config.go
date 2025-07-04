package postgresql

type Config struct {
	Host                  string
	Port                  string
	UserName              string
	Password              string
	DatabaseName          string
	MaxConnections        string
	MaxConnectionIdleTime string
}
