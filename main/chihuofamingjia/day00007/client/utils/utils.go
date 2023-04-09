package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"test/main/chihuofamingjia/day00007/common/message"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //缓冲

}

func (this *Transfer) ReadPag() (mes message.Message, err error) {
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("服务器端——长度读取失败", err)
		//err = errors.New("服务器端——长度读取失败")
		return
	}
	//fmt.Println("服务器端——长度读取：", this.Buf[:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:])
	if n != int(pkgLen) || err != nil {
		fmt.Println("服务器端——长度读取  不相等：", err)
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	return
}

// 发送 数据
func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	_, err = this.Conn.Write(this.Buf[:4])
	if err != nil {
		fmt.Println("客户端——发送长度失败", err)
		return
	}
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("客户端——发送长度失败", err)
		return
	}
	return
}
