package protocol

type DataType byte

const (
	SimpleString DataType = '+'
	Array        DataType = '*'
	BulkString   DataType = '$'
)

const (
	OpCodeModuleAux    byte = 247 /* Module auxiliary data. */
	OpCodeIdle         byte = 248 /* LRU idle time. */
	OpCodeFreq         byte = 249 /* LFU frequency. */
	OpCodeAux          byte = 250 /* RDB aux field. */
	OpCodeResizeDB     byte = 251 /* Hash table resize hint. */
	OpCodeExpireTimeMs byte = 252 /* Expire time in milliseconds. */
	OpCodeExpireTime   byte = 253 /* Old expire time in seconds. */
	OpCodeSelectDB     byte = 254 /* DB number of the following keys. */
	OpCodeEOF          byte = 255
)

type RESPType struct {
	DataType DataType
	Data     []interface{}
}
