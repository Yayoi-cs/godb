package dbg

import (
	"errors"
	"io"
)

func (dbger *TypeDbg) SendLine(payload []byte) error {
	if dbger.stdinWriter == nil {
		return errors.New(stWrong)
	}
	if payload[len(payload)-1] != '\n' {
		payload = append(payload, '\n')
	}

	_, err := dbger.stdinWriter.Write(payload)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) Send(payload []byte) error {
	if dbger.stdinWriter == nil {
		return errors.New(stWrong)
	}
	_, err := dbger.stdinWriter.Write(payload)
	if err != nil {
		return err
	}
	return nil
}

func (dbger *TypeDbg) Recv() ([]byte, error) {
	buf := make([]byte, 1024)
	n, err := dbger.stdoutReader.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf[:n], nil
}
