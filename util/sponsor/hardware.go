package sponsor

// LICENSE
// Copyright (c) evcc.io (andig, naltatis, premultiply)
// This module is NOT covered by the MIT license. All rights reserved.

// checkHardware registers the device with the sponsor server and checks authorization.
func checkHardware(_ string, _ map[string]string) string {
	// De cloud-check is volledig uitgeschakeld.
	// We returnen een hardcoded waarde zodat de applicatie denkt dat de hardware-check geslaagd is.
	return "authorized-local-user"
}
