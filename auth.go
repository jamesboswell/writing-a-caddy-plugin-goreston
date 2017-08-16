
// auth generates a RADIUS authentication request
func auth(r radiusConfig, username string, password string) (bool, error) {
	
	// Create a new RADIUS packet for Access-Request
	packet := radius.New(radius.CodeAccessRequest, []byte(r.Secret))
	packet.Add("User-Name", username)
	packet.Add("User-Password", password)

	reply, err := client.Exchange(packet, r.Server)
	if err != nil {
		return false, fmt.Errorf("RADIUS failure: %s", err)
	}
	// RADIUS Access-Accept is a successful authentication
	if reply.Code == radius.CodeAccessAccept {
		return true, nil
	}
	// Any other reply is a failed authentication
	return false, nil
}
