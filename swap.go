/*
* @Time    : 2021-02-05 19:40
* @Author  : CoderCharm
* @File    : swap.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/

package common

import "encoding/json"

//通过json tag 进行结构体赋值
func SwapTo(request, category interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	return json.Unmarshal(dataByte, category)
}
