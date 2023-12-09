package helpers

import (
	"fmt"
	"path/filepath"
)

func Include(path string) []string {
	files, err := filepath.Glob("admin/views/templates/*.html")
	if err != nil {
		fmt.Println(err)
	}

	pathFiles, e := filepath.Glob("admin/views/" + path + "/*.html") //admin/views/dashboard/list/*.html
	if e != nil {
		fmt.Println(e)
	}

	for _, file := range pathFiles {
		files = append(files, file)
	}
	return files
}
