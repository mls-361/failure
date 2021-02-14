/*
------------------------------------------------------------------------------------------------------------------------
####### failure ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package failure

import (
	"errors"
	"fmt"

	"github.com/mls-361/buffer"
	"github.com/mls-361/logfmt"
)

const (
	_bufferSize = 256
)

type (
	// Failure AFAIRE.
	Failure struct {
		err    error
		errMsg string
		ctx    map[string]interface{}
		msg    string
	}
)

var (
	// ErrNotImplemented AFAIRE.
	ErrNotImplemented = errors.New("not implemented")
)

// New AFAIRE.
func New(err error) *Failure {
	var errMsg string

	if err != nil {
		errMsg = err.Error()
	}

	return &Failure{
		err:    err,
		errMsg: errMsg,
		ctx:    make(map[string]interface{}),
	}
}

// Wrap AFAIRE.
func Wrap(err error) *Failure {
	return &Failure{
		err: err,
	}
}

// NotImplemented AFAIRE.
func NotImplemented() *Failure {
	return New(ErrNotImplemented)
}

// Set AFAIRE.
func (f *Failure) Set(key string, value interface{}) *Failure {
	f.ctx[key] = value
	return f
}

// Setf AFAIRE.
func (f *Failure) Setf(key, format string, args ...interface{}) *Failure {
	return f.Set(key, fmt.Sprintf(format, args...))
}

// Msg AFAIRE.
func (f *Failure) Msg(msg string) *Failure {
	f.msg = msg
	return f
}

// Msgf AFAIRE.
func (f *Failure) Msgf(format string, args ...interface{}) *Failure {
	return f.Msg(fmt.Sprintf(format, args...))
}

// Error AFAIRE.
func (f *Failure) Error() string {
	buf := buffer.New(_bufferSize)
	buf.AppendString(f.msg)

	if len(f.ctx) > 0 {
		if buf.Len() > 0 {
			buf.AppendString(": ")
		}

		logfmt.EncodeMap(buf, f.ctx)
	}

	if f.errMsg != "" {
		if buf.Len() > 0 {
			buf.AppendString(" >>> ")
		}

		buf.AppendString(f.errMsg)
	}

	return buf.String()
}

// Unwrap AFAIRE.
func (f *Failure) Unwrap() error {
	return f.err
}

/*
######################################################################################################## @(°_°)@ #######
*/
