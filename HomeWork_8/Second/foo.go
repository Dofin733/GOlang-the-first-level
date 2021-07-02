package Second

import (
	"flag"
	"fmt"
)


type Config struct {
	Port string
	Db_url string
	Jaeger_url string
	Sentry_url string
	Kafka_broker string
	Some_app_id string
	Some_app_key string
}

func f (){
	var port = flag.String("port","8080","Port number")
	var db_url = flag.String("db", "db-user:db-password@petstore-db:5432/petstore?sslmode=disable","postgress" )
	var jaeger_url = flag.String("jaeger", "http://jaeger:16686","url")
	var sentry_url = flag.String("sentry","http://sentry:9000","url")
	var kafka_broker = flag.String( "kafka_broker","kafka:9092","")
	var some_app_id = flag.String("app_id","testid","")
	var some_app_key = flag.String("app_key","testkey","")
	flag.Parse()

	config := Config{
		Port: *port,
		Db_url: *db_url,
		Jaeger_url: *jaeger_url,
		Sentry_url: *sentry_url,
		Kafka_broker: *kafka_broker,
		Some_app_id: *some_app_id,
		Some_app_key: *some_app_key,
	}
	fmt.Println(config)
	//func foo () *Config{
	//	return & Config{}
	//}
}
func NewConfig () *Config{
	return & Config{}
}
