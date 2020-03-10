package models

type Flag = string

const (
	FlagHost    Flag = "host"
	FlagAdapter Flag = "adapter"
	FlagRepeat  Flag = "repeat"
	FlagOutput  Flag = "output"
	FlagAppend  Flag = "append"
	FlagTimeout Flag = "timeout"
	FlagInput   Flag = "input"
	FlagAuth    Flag = "auth"
)