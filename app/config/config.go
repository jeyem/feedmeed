package config

var (
	DefaultConfig = Config{
		Debug:     true,
		MongoDB:   "feedmeed",
		MongoPort: 27017,
		MongoHost: "127.0.0.1",
		BoltDB:    "feedmeed.db",
		Port:      5000,
		Views:     "views",
		Statics:   "statics",
	}
)

type Config struct {
	MongoDB   string `json:"mongo_db" valid:"required"`
	MongoHost string `json:"mongo_host" required:"required"`
	MongoPort int    `json:"mongo_port" required:"required"`
	BoltDB    string `json:"bolt_db" required:"required"`
	Port      int    `json:"port"`
	Debug     bool   `json:"debug"`
	Views     string `json:"views"`
	Statics   string `json:"statics"`
}
