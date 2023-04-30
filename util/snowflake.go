package util

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

/*
生成全局唯一的雪花ID
*/
var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time

	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

func GenSnowflakeID() int64 {
	return node.Generate().Int64()
}
