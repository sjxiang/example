package main

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

// 基于游标的分页

type Page struct {
	NextID        string `json:"next_id"`           // cursor
	NextTimeAtUTC int64  `json:"next_time_at_utc"`  // 分页发生的时间戳，e.g. 20 min 内有效
	PageSize      int64  `json:"page_size"`         // 每一页的元素个数
}

func NewPage(nextID string, pageSize int64) Page {
	return Page{	
		NextID:        nextID,
		NextTimeAtUTC: time.Now().Unix(),
		PageSize:      pageSize,
	}
}

// Encode 返回分页 token
func (p Page) Encode() Token {
	b, err := json.Marshal(p)
	if err != nil {
		return Token("")
	}
	return Token(base64.StdEncoding.EncodeToString(b))
}

// 发生以下情况之一，属于无效参数
func (p Page) InValid() bool {
	return p.NextID == "" || p.NextTimeAtUTC == 0 || p.NextTimeAtUTC > time.Now().Unix() || p.PageSize <= 0
}

type Token string

// Decode 解析分页信息
func (t Token) Decode() Page {
	var result Page
	if len(t) == 0 {
		return result
	}

	bytes, err := base64.StdEncoding.DecodeString(string(t))
	if err != nil {
		return result
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result
	}

	return result
}
