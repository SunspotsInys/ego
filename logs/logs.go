package logs

import (
	"io"
	"os"
	"runtime"
	"time"
)

type Level uint8

const (
	TRACE Level = iota
	DEBUG
	INFO
	ERROR
)

type logger struct {
	io.Writer
	name  string
	debug string
	info  string
	err   string
}

func (l *logger) write(info string, msg string) (n int, err error) {
	bs := bpl.Get().(*buf)
	bs.buf = bs.buf[:0]
	bs.buf = time.Now().AppendFormat(bs.buf, "[2006-01-02 15:04:05.000]") // 写入时间

	// 写入等级及服务名
	bs.buf = append(bs.buf, info...)

	// 写入调用信息
	appendCaller(bs)

	// 写入 goid
	bs.buf = append(bs.buf, "[goid:"...)
	appendNum(bs, runtime.Goid())

	// 写入日志信息
	bs.buf = append(bs.buf, msg...)
	bs.buf = append(bs.buf, '\n')

	// 写入
	n, err = l.Write(bs.buf)
	if len(bs.buf) < maxLen<<2 { // 太大就抛弃了
		bpl.Put(bs)
	}
	return
}

func New(name string, w io.Writer) (l *logger) {
	name = "[" + name + "]"
	if w == nil {
		return &logger{
			Writer: os.Stdout,
			name:   name,
			debug:  name + "[DEBUG][",
			info:   name + "[INFO][",
			err:    name + "[ERROR][",
		}
	}
	return &logger{
		Writer: w,
		name:   name,
		debug:  name + "[DEBUG][",
		info:   name + "[INFO][",
		err:    name + "[ERROR][",
	}
}
