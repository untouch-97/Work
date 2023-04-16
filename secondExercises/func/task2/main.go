package main

func reads(m map[string]string, key string) {
	value, _ := m[key]
}

func main() {
	m := make(map[string]string)
	m["Alik"] = "alchemist"
}
