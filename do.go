package main
import (
    "os"
    "log"
    "fmt"
    "github.com/garyburd/redigo/redis"
)

func main() {
    
     file, err := os.OpenFile("bbb.log",  os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    //file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE, 666)
    //file, err := os.Create("./log.txt");
    options := redis.DialPassword("yizhiqingxie")
    c, err := redis.Dial("tcp", "47.105.32.99:6378",options)
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }
    defer c.Close()
    v, err := c.Do("DEL", "time")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(v)
    defer file.Close()
    logger := log.New(file, "", log.LstdFlags) 
    logger.Println("将日志写入文件")
}
