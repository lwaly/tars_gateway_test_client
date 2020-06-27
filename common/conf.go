package common

import "ohplay_gateway/gateway/common/conf"

var Conf *goconfig.ConfigFile

func init() {
	var err error
	Conf, err = goconfig.LoadConfigFile("./conf/conf")
	if err != nil {
		panic(err)
	}
}
