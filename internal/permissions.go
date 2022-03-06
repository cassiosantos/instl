package internal

import (
	"os"
)

// CheckPermissionsToDir checks if instl can write to a directory
func CheckPermissionsToDir(dirpath string) bool {
	// Create temporary test file
	err := os.WriteFile(dirpath+"/instl_write_test.tmp", []byte("test"), 0644)
	if err != nil {
		return false
	}

	// Remove temporary test file
	err = os.Remove(dirpath + "/instl_write_test.tmp")
	if err != nil {
		return false
	}

	return true
}
