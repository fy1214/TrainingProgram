package model

import . "TrainingProgram/resource"

//执行数据迁移
func migration() {
	// 自动迁移模式
	_ = DB.AutoMigrate(&Member{}, &Course{}, &StudentCourse{})
}
