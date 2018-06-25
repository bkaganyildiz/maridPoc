package conf

func ReadLocalConfiguration(path string) (conf map[string]string) {
	 return readConfigurations(path)
}