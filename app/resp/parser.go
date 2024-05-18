package resp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)

type Parser struct {
	r *bufio.Reader
}

func NewParser(conn net.Conn) *Parser {
	return &Parser{r: bufio.NewReader(conn)}
}

func (p *Parser) ParseCommand() (*Command, error) {
	dataType, _ := p.r.ReadByte()
	if dataType != byte(protocol.Array) {
		return nil, fmt.Errorf("Invalid Command")
	}

	p.r.UnreadByte()
	args, err := p.parseArray()
	if err != nil {
		return nil, err
	}

	commandName := CommandName(strings.ToUpper(args[0]))
	command := &Command{Name: commandName, Args: args[1:]}

	return command, nil
}

func (p *Parser) parseArray() ([]string, error) {
	dataType, _ := p.r.ReadByte()
	if dataType != byte(protocol.Array) {
		return nil, fmt.Errorf("Invalid Array")
	}

	chunk := p.readChunk()
	size, _ := strconv.Atoi(chunk)

	args := make([]string, 0)
	for i := 0; i < size; i++ {
		arg, err := p.parseBulkString()
		if err != nil {
			return nil, fmt.Errorf("Invalid Bulk String")
		}
		args = append(args, arg)
	}

	return args, nil
}

func (p *Parser) parseBulkString() (string, error) {
	dataType, _ := p.r.ReadByte()
	if dataType != byte(protocol.BulkString) {
		return "", fmt.Errorf("INVALID BULK STRING")
	}

	bulkString := p.readChunk(2)

	return bulkString, nil
}

func (p *Parser) readChunk(n ...int) string {
	var chunkIndex int
	if len(n) > 0 {
		chunkIndex = n[0] - 1
	} else {
		chunkIndex = 0
	}

	var chunk string
	var err error

	for i := 0; i <= chunkIndex; i++ {
		chunk, err = p.r.ReadString('\n')
		if err != nil {
			return ""
		}
	}

	// Remove \r
	chunk = chunk[:len(chunk)-2]

	return chunk
}
