package tasks

const QUEUE_NAME = "vultisigner"

const (
	TypeKeyGeneration     = "key:generation"
	TypeKeySign           = "key:sign"
	TypeReshare           = "key:reshare"
	TypeKeyGenerationDKLS = "key:generationDKLS"
	TypeKeySignDKLS       = "key:signDKLS"
	TypeReshareDKLS       = "key:reshareDKLS"
	TypeMigrate           = "key:migrate"
	TypeImport            = "key:import"
	TypeCreateMldsa       = "key:createMldsa"
	TypeKeygenBatch       = "key:keygenBatch"
	TypeReshareBatch      = "key:reshareBatch"
	TypeImportBatch       = "key:importBatch"
)
