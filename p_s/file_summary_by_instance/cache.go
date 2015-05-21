package file_summary_by_instance

import (
	"errors"
)

// provide a mapping from filename to table.schema etc

var (
	mappedName     map[string]string
	total, matched int
)

func init() {
	// setup on startup
	mappedName = make(map[string]string)
}

func getFromCache(key string) (result string, err error) {
	total++
	if result, ok := mappedName[key]; ok {
		matched++
		//		lib.Logger.Println("matched/total:", matched, total)
		return result, nil
	}
	//		lib.Logger.Println("matched/total:", matched, total)
	return "", errors.New("Not found")
}

func saveToCache(key, value string) string {
	mappedName[key] = value
	return value
}
