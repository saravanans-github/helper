package fs

import (
	"log"
	"os"
)

// MakeDirectory create a new directory specified at newDir with file permisions perm. if forceCreate = true then the directory (and its contents) are overwritten.
func MakeDirectory(newDir string, perm os.FileMode, forceCreate bool) error {
	// if force create then remove the directory and it's contents if it exists
	if forceCreate {
		log.Printf("Force create = true. Removing directory [%s]...", newDir)
		err := os.RemoveAll(newDir)
		if err != nil {
			log.Printf("Force create = true. Removing directory [%s]... FAILED [%s]", newDir, err.Error())
		}
	}
	log.Printf("Force create = true. Removing directory [%s]... Done", newDir)

	// then create the directory
	log.Printf("Creating directory [%s]...", newDir)
	err := os.Mkdir(newDir, perm)
	if err != nil {
		log.Printf("Creating directory [%s]... FAILED [%s]", newDir, err.Error())
	}
	log.Printf("Creating directory [%s]... Done", newDir)

	return nil
}

// WriteToFile is an async go routine which writes to the file specifed by dst. Bytes are written as they are written to the channel data.
// The method will set writeComplete to true once the operation is complete. this can be used to block the routine. any errors will be written in err.
func WriteToFile(dst string, data chan []byte, err *error, writeComplete *bool) {
	log.Printf("Creating file [%s]... ", dst)
	var file *os.File
	file, *err = os.Create(dst)
	if *err != nil {
		log.Printf("Creating file [%s]... FAILED [%s]", dst, (*err).Error())
	}
	log.Printf("Creating file [%s]... Done", dst)

	log.Printf("Writing to file [%s]... ", dst)
	go func() {
		defer file.Close()
		for p := range data {
			var n int
			n, *err = file.Write(p)
			if *err != nil {
				log.Printf("Writing to file [%s]... FAILED [%s]", dst, (*err).Error())
				return
			}
			if n != len(p) {
				log.Printf("Writing to file [%s]... WARNING [%d out of %d bytes were written]", dst, n, len(p))
			}
		}
		log.Printf("Writing to file [%s]... Done", dst)
		*writeComplete = true
	}()
}
