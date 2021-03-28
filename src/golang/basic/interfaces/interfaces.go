package main

import (
	"bytes"
	"fmt"
	//"io"
)

/*
Go implicitly implement the interface, unlike Java or C#, explicitly do it
Best practices
Use many, small interfaces
	Single method interfaces are some of the most powerful and flexible
		io.Writer, io.Reader, interface{}
	Don't export interfaces for types that will be consumed
	Do export interfaces for types that will be used by package
	Design functions and methods to receive interfaces whenever possible

Type conversion
var wc WriterCloser = NewBufferedWriterCloser()
bwc := wc.(*BufferedWriterCloser)  // inside the parentheses we put the type we want to cast to

The empty interface and type switches
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
		break
		// fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}

Implementing with values vs. pointers
	Method set of value is all methods with value receivers
	Method set of pointer is all methods, regardless of receiver type

*/

func main() {
	var w Writer = ConsoleWriter{}
	_, _ = w.Write([]byte("Hello consoleWriter"))
	fmt.Println("")

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	_, _ = wc.Write([]byte("Hello YouTube listeners, this is a test"))
	_ = wc.Close()

	// r, ok := wc.(io.Reader)
	r, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Embedding interface with interface
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}


type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}


type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}