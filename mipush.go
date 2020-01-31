package mipush

import (
    "fmt"
    "os"
    "strconv"
    "time"
    
    "go.uber.org/zap"
    
    "github.com/houseme/mipush/api"
    "github.com/houseme/mipush/builder"
    "github.com/houseme/mipush/result"
    "github.com/houseme/mipush/util"
)

func SendMessage(builder *builder.Builder, params builder.CommonParams) (*result.Result, error) {
    return api.SendMessageByRegIds(builder, params)
}

func main() {
    fmt.Println("os.Args: ", os.Args)
    for _, v := range os.Environ() {
        //输出系统所有环境变量的值
        fmt.Println(v)
    }
    // 获取单个变量
    home := os.Getenv("HOME")
    fmt.Println("Home:", home)
    fmt.Println("pwd:", os.Getenv("PWD"))
    util.SetLevel(util.DebugLevel)
    util.Info("测测谁---")
    timeStart := time.Now().UnixNano()
    fmt.Println("timeStart ", timeStart)
    for i := 0; i < 1; i++ {
        fileName := "MiPush-" + strconv.Itoa(i%10)
        fmt.Println(fileName)
        logger := util.Logger(fileName)
        logger.Info("log 初始化成功")
        logger.Info("无法获取网址 ",
            zap.String("url", "http://www.baidu.com"),
            zap.Int("attempt", 3),
            zap.Duration("backoff", time.Second),
            zap.Int("I: ", i))
    }
    endTime := time.Now().UnixNano()
    fmt.Println("endTime ", endTime)
    fmt.Println("end - start time :", endTime-timeStart)
}

//timeStart  1580467384980164000
//endTime  1580467384986640000
//end - start time : 6476000
//end - start time : 5589000
//end - start time : 529780000 10000
//end - start time : 611710000
//end - start time : 910518000
