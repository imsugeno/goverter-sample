package query

import "time"

// TaskResult はクエリ層のタスク出力型。
// API型 (gen.Task) との意図的な差異:
//   - ID (string)        vs Api型の Id        → フィールド名の違い
//   - Description (string)  vs Api型の *string   → ポインタの違い
//   - CreatedAt (time.Time) vs Api型の *time.Time → ポインタの違い
//   - Status (string)       vs Api型の TaskStatus → named type の違い
type TaskResult struct {
	ID          string
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
}
