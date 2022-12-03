package main

func deleteTest() {

	//把所有name为空的和age为0的删除
	var users []TestUser
	DB.Where("age = ?", 0).Delete(users) //会进行软删除，也就是五条记录不会删除，
	// 记录还在 ，但是会记录delete_at标记，可以用下面的硬删除

	//把所有name为空的和age为0的删除
	var user []TestUser
	DB.Unscoped().Where("age = ?", 0).Delete(user)

}
