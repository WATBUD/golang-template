package mongodb

import "os"

func defaultOrEnv() (uri, dbName string) {
	// [MONGODB_URI]
	uri = "mongodb+srv://williamchuang:hkdlOv3cRPlHUTR1@maitoday.kfzdeqh.mongodb.net/?retryWrites=true&w=majority&appName=MaiToday"

	// [MONGODB_DB_NAME]
	dbName = "mai_dev"

	if v, ok := os.LookupEnv("MONGODB_URI"); ok {
		uri = v
	}

	if v, ok := os.LookupEnv("MONGODB_DB_NAME"); ok {
		dbName = v
	}

	return
}

func defaultUriOrEnv() (uri string) {
	// [MONGODB_URI]
	uri = "mongodb+srv://williamchuang:hkdlOv3cRPlHUTR1@maitoday.kfzdeqh.mongodb.net/?retryWrites=true&w=majority&appName=MaiToday"

	if v, ok := os.LookupEnv("MONGODB_URI"); ok {
		uri = v
	}

	return
}

func defaultDbNameOrEnv() (dbName string) {
	// [MONGODB_DB_NAME]
	dbName = "mai_dev"

	if v, ok := os.LookupEnv("MONGODB_DB_NAME"); ok {
		dbName = v
	}

	return
}
