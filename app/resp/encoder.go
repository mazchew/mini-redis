package resp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	
	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)

type Encoder struct {
	w *bufio.Writer
}

func NewEncoder(conn net.Conn) *Encoder {
	return &Encoder{w: bufio.NewWriter(conn)}
}

func (e *Encoder) Encode(respType *protocol.RESPType) string {
	fmt.Println("===== here: ", respType)
	switch respType.DataType {
	case protocol.SimpleString:
		return fmt.Sprintf("+%s\r\n", respType.Data[0].(string))
	case protocol.BulkString:
		data, ok := respType.Data[0].(string)
		if !ok {
            return "$-1\r\n"  // Fallback to null bulk string if type assertion fails or data is nil
        }
        return fmt.Sprintf("$%d\r\n%s\r\n", len(data), data)
	case protocol.Array:
		var res strings.Builder
		res.WriteString(fmt.Sprintf("*%d\r\n", len(respType.Data)))
		for _, item := range respType.Data {
			encodedItem := e.Encode(item.(*protocol.RESPType))
			res.WriteString(encodedItem)
		}
		return res.String()
	default:
		return ""
	}
}

func (e *Encoder) Write(respType *protocol.RESPType) error {
	data := e.Encode(respType)
	e.w.WriteString(data)
	err := e.w.Flush()
	if err != nil {
		return err
	}
	return nil
}
