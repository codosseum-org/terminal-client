package upload

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


func UploadCode(code string) error {
	return nil
}

func GetFileInformation(path string) (name, language string) {
	name = filepath.Base(path)
	extension := filepath.Ext(path)

	language = getLanguageFromExtension(extension)
	return name, language
}

func GetFileContent(path string) (data []byte, error error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file '%s' not found", path)
		}
		if os.IsPermission(err) {
			return nil, fmt.Errorf("file '%s' was denied access to", path)
		}
		return nil, err
	}

	if !info.Mode().IsRegular() {
		return nil, fmt.Errorf("file '%s' is not a regular file", path)
	}

	data, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, err
}

func getLanguageFromExtension(extension string) string {
	languages := map[string]string{
		".go":   "go",
		".java": "java",
		".rs":   "rust",
		".py":   "python",
		".hs":   "haskell",
		".rb":   "ruby",
		".js":   "javascript",
		".ts":   "typescript",
		".sh":   "bash",
		".lua":  "lua",
		".kt":   "kotlin",
		".c":    "c",
		".cpp":  "cpp",
		".cs":   "csharp",
		// add all the language the game will support!
	}

	extension = strings.ToLower(extension)

	if language, ok := languages[extension]; ok {
		return language
	}

	return ""
}
