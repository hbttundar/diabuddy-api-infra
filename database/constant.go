package database

const (
	SelectOperation       = "select"
	InsertOperation       = "insert"
	InsertOperationWithId = "insert_with_id"
	UpdateOperation       = "update"
	SoftDeleteOperation   = "soft_delete"
	HardDeleteOperation   = "hard_delete"
	RestoreOperation      = "restore"
	ContextTransactionKey = "transaction"
)
