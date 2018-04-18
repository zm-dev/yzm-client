package distinguish

import (
	"io"
	"bufio"
	"strconv"
	"sort"
	"strings"
)

type MappingLines map[uint16]string

func (ml MappingLines) ToSortedString() string {
	var keys []int
	for k := range ml {
		keys = append(keys, int(k))
	}
	sort.Ints([]int(keys))

	stringBuilder := &strings.Builder{}

	for _, k := range keys {
		stringBuilder.WriteString(ml[uint16(k)])
		stringBuilder.WriteByte('\n')
	}
	return stringBuilder.String()
}

func LoadMappingLines(mappingReader io.Reader) MappingLines {
	mappingLines := make(MappingLines)
	bufReader := bufio.NewReader(mappingReader)
	for {
		b, _, err := bufReader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			// todo error handing
			continue
		}
		filename, err := strconv.Atoi(string(b[:4]))
		if err != nil {
			// todo error handing
			continue
		}
		mappingLines[uint16(filename)] = string(b)
	}

	return mappingLines
}
