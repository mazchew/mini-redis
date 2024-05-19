package protocol

type DataType byte

const (
	SimpleString DataType = '+'
	Array        DataType = '*'
	BulkString   DataType = '$'
)

type RESPType struct {
	DataType DataType
	Data     []interface{}
}