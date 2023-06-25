// Author:      xuan
// Date:        2023/6/25
// Description:  VO

package vo

import (
	"time"
)

type TodoVO struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`
	Title      string    `gorm:"column:title;not null;comment:标题" json:"title"`
	Status     bool      `gorm:"column:status;not null;comment:状态(是否完成)" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
}
