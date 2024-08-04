package utils

import "github.com/bwmarrin/snowflake"

var snowflakeNode *snowflake.Node

func init() {
	var err error
	snowflakeNode, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

func NewID() snowflake.ID {
	return snowflakeNode.Generate()
}
