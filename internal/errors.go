package internal

import "errors"

var (
	ErrNoAssetFound = errors.New("could not find an asset which fits to the system")
)
