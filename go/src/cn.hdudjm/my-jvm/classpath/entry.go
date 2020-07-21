package classpath

import "os"
import "strings"

// :(linux/unix) or ;(windows)
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 寻找加载class文件
	readClass(className string) ([]byte, Entry, error)
	// 相当于java的toString()
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}

	return newDirEntry(path)
}
