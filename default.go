package pond

// defaultPool is the default pool used by the package-level functions.
var defaultPool = newPool(0, nil)

// Submit submits a task to the default pool and returns a future that can be used to wait for the task to complete.
func Submit(task func()) Task { _ = "STUB: not implemented"; return *new(Task) }

// SubmitErr submits a task to the default pool and returns a future that can be used to wait for the task to complete.
func SubmitErr(task func() error) Task { _ = "STUB: not implemented"; return *new(Task) }

// NewGroup creates a new task group with the default pool.
func NewGroup() TaskGroup { _ = "STUB: not implemented"; return *new(TaskGroup) }

// NewSubpool creates a new subpool with the default pool.
func NewSubpool(maxConcurrency int) Pool { _ = "STUB: not implemented"; return *new(Pool) }
