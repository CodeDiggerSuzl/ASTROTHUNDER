package packaging

import (
	"encoding/binary"
	"errors"
	"io"
)

const MsgHeader = "12345678"

func Encode(bytesBuffer io.Writer, content string) error {
	//  消息格式 msgHeader + content_len + content
	// 8 + 4 + content_len(字节长度）
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(MsgHeader)); err != nil {
		return err
	}
	contentLen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, contentLen); err != nil {
		return err
	}
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

// decode
func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	magicBuf := make([]byte, len(MsgHeader))
	if _, err = io.ReadFull(bytesBuffer, magicBuf); err != nil {
		return nil, err
	}
	if string(magicBuf) != MsgHeader {
		return nil, errors.New("msg_header error")
	}
	lenBuf := make([]byte, 4)
	if _, err = io.ReadFull(bytesBuffer, lenBuf); err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lenBuf)
	bodyBuf = make([]byte, length)
	if _, err = io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, err
}
