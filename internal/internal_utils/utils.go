package internal_utils

func GetUUIDFromString(uuid string) [16]byte {
	uuidBytes := [16]byte{}
	copy(uuidBytes[:], []byte(uuid))
	return uuidBytes
}
