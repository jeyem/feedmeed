package config

var (
	DefaultConfig = Config{
		Debug:          true,
		MongoDB:        "feedmeed",
		MongoPort:      27017,
		MongoHost:      "127.0.0.1",
		Port:           5000,
		Views:          "views",
		Statics:        "statics",
		GoogleClientID: "469400100676-qp17a7tub5o9dm1psjbmrejo27sakagj.apps.googleusercontent.com",
		GoogleSecret:   "3amTv6jPVHrLW0Y7dnBTQsgk",
	}
)

type Config struct {
	MongoDB        string `json:"mongo_db" valid:"required"`
	MongoHost      string `json:"mongo_host" required:"required"`
	MongoPort      int    `json:"mongo_port" required:"required"`
	Port           int    `json:"port"`
	Debug          bool   `json:"debug"`
	Views          string `json:"views"`
	Statics        string `json:"statics"`
	GoogleClientID string `json:"google_client_id"`
	GoogleSecret   string `json:"google_secret"`
}
