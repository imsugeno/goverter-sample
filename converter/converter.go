package converter

import (
	"time"

	gen "github.com/imsugeno/query-generator-sample/api/gen"
	"github.com/imsugeno/query-generator-sample/query"
)

// goverter:converter
// goverter:output:file ./gen/converter.gen.go
// goverter:output:package github.com/imsugeno/query-generator-sample/converter/gen
// goverter:extend TimeToTimePtr
// goverter:useZeroValueOnPointerInconsistency
type Converter interface {
	// goverter:map ID Id
	ConvertTask(source query.TaskResult) gen.Task
	ConvertTaskList(source []query.TaskResult) []gen.Task
}

// TimeToTimePtr は time.Time → *time.Time の変換ヘルパー。
func TimeToTimePtr(t time.Time) *time.Time {
	return &t
}
