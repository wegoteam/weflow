package utils

import "github.com/zhongguo168a/stcopy"

/*
*
bean属性src拷贝给dst
*/
func BeanCopy(dst, src interface{}) (err error) {
	return stcopy.New(src).To(dst)
}
