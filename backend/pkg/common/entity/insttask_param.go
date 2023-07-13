package entity

import "time"

// InstTaskParamResult
// @Description: 实例任务参数返回结果
type InstTaskParamResult struct {
	ID            int64     // 唯一id
	InstTaskID    string    // 实例任务id
	ParamID       string    // 参数id;参数的唯一标识
	ParamName     string    // 参数名称;参数的名称
	ParamValue    string    // 参数值
	CreateTime    time.Time // 创建时间
	UpdateTime    time.Time // 更新时间
	ParamDataType string    // 参数数据类型【string：字符串；int ：整形数值；float：浮点型数值；object：对象；array：数组；decimal：金额；long：长整型；table：表格；boolean：布尔】
	ParamBinary   []byte    // 参数二进制值;可存二进制值对象
}
