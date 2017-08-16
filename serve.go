func (a RADIUS) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	// Pass-through if not protected path
	if a.Config.requestFilter != nil && !a.Config.requestFilter.shouldAuthenticate(r) {
		return a.Next.ServeHTTP(w, r)
	}
	
	// Check for HTTP Basic Authorization Headers and valid username, password
	username, password, ok := r.BasicAuth()

	// send username, password to RADIUS server(s) for authentication
	// returns isAuthenticated if authentication successful
	// err if no RADIUS servers respond
	isAuthenticated, err := auth(a.Config, username, password)

	// err handling removed for presentation

	// if RADIUS authentication failed, return 401
	if !isAuthenticated {
		w.Header().Set("WWW-Authenticate", realm)
		return http.StatusUnauthorized, nil
	}
	return a.Next.ServeHTTP(w, r)
}