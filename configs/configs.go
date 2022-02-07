package configs

import (
	"io/ioutil"

	"github.com/inyscc/ego/jsons"
)

// 读取 json 配置文件
func InitConfig[T any](cfgPath string, cfg *T, callback func(t *T)) (err error) {
	bs, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return err
	}
	err = jsons.Unmarshal(bs, cfg)
	if err != nil {
		return err
	}
	if callback != nil {
		callback(cfg)
	}
	return err
}
