package app

func init() {
	SetUpConf()
	SetUpDB()
}

func Close() {
	_ = DB.Close()
}
