package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Person struct {
	Magic   [8]byte
	NameLen int32
	Name    string
	Age     int32
}

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func (self *Person) Encode() ([]byte) {
	var err error
	buf := new(bytes.Buffer)

	// binary 的写入方法只运行fixed的类型
	err = binary.Write(buf, binary.BigEndian, self.Magic)
	check_err(err)

	err = binary.Write(buf, binary.BigEndian, self.NameLen)
	check_err(err)

	_, err = buf.Write([]byte(self.Name))
	check_err(err)

	err = binary.Write(buf, binary.BigEndian, self.Age)
	check_err(err)

	return buf.Bytes()
}

func (self *Person) Decode(data []byte) {
	var err error
	buf := bytes.NewBuffer(data)

	err = binary.Read(buf, binary.BigEndian, &self.Magic)
	check_err(err)

	err = binary.Read(buf, binary.BigEndian, &self.NameLen)
	check_err(err)

	tmpName := buf.Next(int(self.NameLen))
	check_err(err)

	self.Name = string(tmpName)

	err = binary.Read(buf, binary.BigEndian, &self.Age)
	check_err(err)
}

func main() {
	p := Person{
		//Magic: [8]byte{'m', 'a', 'g', 'i', 'c', 0, 0},
		Name: "redchen",
		Age: 29,
	}
	copy(p.Magic[:], []byte("per_v1"))
	p.NameLen = int32(len(p.Name))

	b := p.Encode()
	fmt.Printf("%s\n", b)

	np := new(Person)
	np.Decode(b)

	fmt.Printf("%#v\n", np)

}
