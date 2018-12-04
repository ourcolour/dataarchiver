package utils

import (
	"container/list"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func FindMatchedGroupByGroupIndex(text string, pattern string, groupIndex int) (*list.List, error) {
	// 定义
	var (
		matchedList *list.List

		result *list.List
		err    error
	)

	// 参数
	if groupIndex < 0 {
		err = errors.New("Invalid parameter groupIndex, should >=0.")
		return result, err
	}

	// 匹配
	matchedList, err = FindMatchedGroups(text, pattern)
	if nil != err {
		return result, err
	}

	// 提取结果
	result = list.New()
	for item := matchedList.Front(); item != nil; item = item.Next() {
		if nil == item {
			continue
		}

		var curMap map[int]string = item.Value.(map[int]string)
		if len(curMap) <= groupIndex {
			err = errors.New(fmt.Sprintf("Index out of range, parameter groupIndex value is %d.", groupIndex))
			return result, err
		}

		result.PushBack(curMap[groupIndex])
	}

	return result, err
}
func FindMatchedGroups(text string, pattern string) (*list.List, error) {
	// 定义
	var (
		matches [][]string

		result *list.List
		err    error
	)

	// 参数
	if "" == strings.TrimSpace(text) {
		err = errors.New("Invalid parameter text.")
		return result, err
	}
	if "" == strings.TrimSpace(pattern) {
		err = errors.New("Invalid parameter pattern.")
		return result, err
	}

	// 匹配
	matches = regexp.MustCompile(pattern).FindAllStringSubmatch(text, -1)

	// 添加
	result = list.New()
	for _, row := range matches {
		if nil == row {
			continue
		}

		var item map[int]string = make(map[int]string)
		for cIdx, col := range row {
			if "" == strings.TrimSpace(col) {
				continue
			}
			item[cIdx] = col
		}
		result.PushBack(item)
	}

	return result, err
}
