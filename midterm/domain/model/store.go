package store

// store key and value of strings
var Map = make(map[string]string)

// provides creating mock pair
func Create() {
	Map["test"] = "Already created test"
}

// provides setting key and value pair to the map
func Set(key string, value string) {
	Map[key] = value
}

// provides getting value by key
func Get(key string) string {
	return Map[key]
}

// provides updating key and value pair
func Update(key string, value string) {
	Map[key] = value
}
