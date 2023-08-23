package model

// do migration
func migration() {
	// 自动迁移模式
	_ = DB.AutoMigrate(&Message{}, &Consumer{})
}
