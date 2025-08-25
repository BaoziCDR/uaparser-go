package uaparser

import (
	"strconv"
	"strings"
)

// Comparable 接口
type Comparable interface {
	Compare(other Comparable) int
}

// IntComparable 实现
type IntComparable int

func (i IntComparable) Compare(other Comparable) int {
	o := other.(IntComparable)
	if i < o {
		return -1
	}
	if i > o {
		return 1
	}
	return 0
}

// VersionComparable 实现
type VersionComparable string

func (v VersionComparable) Compare(other Comparable) int {
	o := other.(VersionComparable)
	parts1 := strings.Split(string(v), ".")
	parts2 := strings.Split(string(o), ".")
	length := len(parts1)
	if len(parts2) > length {
		length = len(parts2)
	}
	for i := 0; i < length; i++ {
		var num1, num2 int // 默认为0
		if i < len(parts1) {
			num1, _ = strconv.Atoi(parts1[i])
		}
		if i < len(parts2) {
			num2, _ = strconv.Atoi(parts2[i])
		}
		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
	}
	return 0
}

// 解析值
func parseComparable(str string, value Comparable) Comparable {
	str = strings.Trim(str, "'")
	switch value.(type) {
	case IntComparable:
		intVal, _ := strconv.Atoi(str)
		return IntComparable(intVal)
	case VersionComparable:
		return VersionComparable(str)
	default:
		return nil
	}
}

// MatchRange 通用范围匹配
func MatchRange(rangeStr string, value Comparable) bool {
	// 去掉空格
	rangeStr = strings.TrimSpace(rangeStr)
	// 检查区间的开闭类型
	lowerClosed := strings.HasPrefix(rangeStr, "[")
	upperClosed := strings.HasSuffix(rangeStr, "]")
	// 去掉括号，获取纯数字部分
	rangeStr = strings.Trim(rangeStr, "[]()")
	// 分割范围的上下界
	parts := strings.Split(rangeStr, ",")
	// 解析下界
	var lowerBound Comparable
	var lowerBoundExists bool
	if parts[0] != "" {
		lowerBoundExists = true
		lowerBound = parseComparable(parts[0], value)
	}
	// 解析上界
	var upperBound Comparable
	var upperBoundExists bool
	if len(parts) > 1 && parts[1] != "" {
		upperBoundExists = true
		upperBound = parseComparable(parts[1], value)
	}
	// 检查值是否在范围内
	if lowerBoundExists {
		comp := value.Compare(lowerBound)
		if lowerClosed {
			if comp < 0 {
				return false
			}
		} else {
			if comp <= 0 {
				return false
			}
		}
	}
	if upperBoundExists {
		comp := value.Compare(upperBound)
		if upperClosed {
			if comp > 0 {
				return false
			}
		} else {
			if comp >= 0 {
				return false
			}
		}
	}
	return true
}
