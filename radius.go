type RADIUS struct {
	// Connection
	Next     httpserver.Handler
	SiteRoot string
	Config   radiusConfig
	db       *bolt.DB  // a bolt db
}
