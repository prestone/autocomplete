package autocomplete

import "github.com/OneOfOne/xxhash"

func hash(s []byte) int {
	return int(xxhash.Checksum32(s))
}
