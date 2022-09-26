package main
 
import (
    "bufio"
    "fmt"
    "io"
    "net"
    "time"
)
 
var port = "0.0.0.0:9001"
func echoMessage(conn net.Conn) {
    r := bufio.NewReader(conn)
    for {
        message, err := r.ReadBytes(byte('\n'))
        switch err {
        case nil:
            break
        case io.EOF:
        default:
            fmt.Println("Error ", err)
 
        }
        t := time.Now()
        myTime := t.Format(time.RFC3339)
        conn.Write([]byte("Received at server on "))
        conn.Write([]byte(myTime))
        conn.Write([]byte(" : "))
        conn.Write(message)
        fmt.Println(conn.RemoteAddr(), " : ", string(message))
    }
 
}
func main() {
    fmt.Println("Start the server ")
    ln, err := net.Listen("tcp", port)
    for {
        conn, _ := ln.Accept()
        if err != nil {
            fmt.Println("Error ", err)
            continue
        }
 
        go echoMessage(conn)
 
    }
 
}