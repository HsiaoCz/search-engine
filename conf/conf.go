package conf

import "os"

func GetPort(key string) string {
	port := os.Getenv(key)
	if port == "" {
		return ":4001"
	}
	return port
}

func GetMongoUrl(key string) string {
	mongoUrl := os.Getenv(key)
	if mongoUrl == "" {
		return "mongodb://localhost:27017"
	}
	return mongoUrl
}

func GetMongoDBName(key string) string {
	dbname := os.Getenv(key)
	if dbname == "" {
		return "search"
	}
	return dbname
}

func GetUserColl(key string) string {
	userColl := os.Getenv(key)
	if userColl == "" {
		return "users"
	}
	return userColl
}
