package modelconfig

type Config struct {
	General General `toml:"general"`
}

type General struct {
	URL string `toml:"url"`
}
