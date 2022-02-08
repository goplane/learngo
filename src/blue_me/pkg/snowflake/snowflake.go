package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

//以下是写了一个id生成器，这个项目用不上分布式生成器服务，就写了一个模块
// sf = snowflake
var node *sf.Node //全局节点
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000 //毫秒值
	node, err = sf.NewNode(machineID)  //指定机器的id
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
