package utils

import (
	"bytes"
	"encoding/gob"
	beanUtil "github.com/wegoteam/wepkg/bean"
)

// BeanCopy
// @Description: bean属性src拷贝给dst
// @param: dst
// @param: src
// @return err
func BeanCopy(dst, src interface{}) (err error) {
	beanCopyErr := beanUtil.BeanCopy(dst, src)
	if len(beanCopyErr) > 0 {
		return beanCopyErr[0]
	}
	return nil
}

// DeepCopy
// @Description: 小写变量，函数拷贝成功
// @param: dst
// @param: src
// @return error
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
