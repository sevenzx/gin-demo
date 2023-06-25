// Author:      xuan
// Date:        2023/6/24
// Description:	操作数据库

package main

import (
	"fmt"
	"gin-demo/study/gen/common"
	"gin-demo/study/gen/enum"
	"gin-demo/study/gen/model/vo"
	"gin-demo/study/gen/query"
	"github.com/goccy/go-json"
	"github.com/jinzhu/copier"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"

func queryByPage(pr common.PageRequest) vo.PageVO[vo.UserVO] {
	// 获取基础翻页信息
	current := pr.Current
	pageSize := pr.PageSize
	offset := (current - 1) * pageSize
	u := query.User

	// 总数
	var total int64
	if pr.NeedTotal {
		count, _ := u.Count()
		fmt.Println("count = ", count)
		total = count
	}

	// 查询数据库
	users, _ := u.Offset(offset).Limit(pageSize).Find()

	// 使用UserVO进行信息脱敏
	userVOList := make([]vo.UserVO, 0)
	for _, user := range users {
		var userVO vo.UserVO
		_ = copier.Copy(&userVO, &user)
		userVOList = append(userVOList, userVO)
	}

	pageVO := vo.PageVO[vo.UserVO]{
		Records:  userVOList,
		Current:  current,
		PageSize: pageSize,
		Total:    total,
	}
	return pageVO
}

func main() {
	db, _ := gorm.Open(mysql.Open(dsn))
	// 设置默认数据库
	query.SetDefault(db)
	u := query.User

	// 软删除测试成功
	// info, _ := u.Where(u.ID.Eq(1)).Delete()
	// fmt.Println(info)

	// 指定字段查询
	users, _ := u.Select(u.UserAccount, u.Username, u.CreateTime, u.UpdateTime).Find()
	for _, user := range users {
		bytes, _ := json.Marshal(user)
		fmt.Println(string(bytes))
	}

	// 多条件查询
	user, _ := u.Where(u.Gender.Eq(enum.Male), u.Username.Eq("mars")).First()
	fmt.Println(user)

	// 更新
	_, _ = u.Where(u.ID.Eq(2)).Update(u.Username, "Ease")

	// 翻页查询
	pr := common.GetDefaultPageRequest()
	pr.Current = 2
	pr.PageSize = 2
	pageVO := queryByPage(pr)
	fmt.Println(pageVO)
}
