package golug_env

import (
	"os"
	"strings"
)

func init() {
	// 环境变量处理, key转大写, 同时把-./转换为_
	replacer := strings.NewReplacer("-", "_", ".", "_", "/", "_")
	for _, env := range os.Environ() {
		if envs := strings.SplitN(env, "=", 2); len(envs) == 2 && trim(envs[0]) != "" {
			_ = os.Unsetenv(envs[0])
			key := replacer.Replace(upper(trim(envs[0])))
			_ = os.Setenv(key, envs[1])
		}
	}
}
