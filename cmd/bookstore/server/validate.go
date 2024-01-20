package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)


func ValidatePageToken(value string) error {
	// token 为 ""，则默认第一页，1
	if len(value) == 0 {
		return nil
	} 
	
	// token 若不为空，校验是否合规
	if Token(value).Decode().InValid() {
		return nil 	
	}

	// 不合规矩
	return errors.New("参数无效")
}

func ValidateShelfTheme(value string) error {
	return ValidateString(value, 2, 20)
}

func ValidateShelfID(value int64) error {
	// 无效状态，命中
	if value <= 0 {
		// 提示，每与操反
		return fmt.Errorf("id 要是正整数")
	}

	return nil 
}

func ValidateString(value string, minLength int, maxLength int) error {
	// n := len(value)  // 字节长度
	
	n := utf8.RuneCountInString(value)
	
	if n < minLength || n > maxLength {
		return fmt.Errorf("字符长度必须在 %d-%d 之间", minLength, maxLength)
	}
	return nil
}
