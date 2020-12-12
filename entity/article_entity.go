package entity

// ArticleEntity 文章数据实体
type ArticleEntity struct {
	ID         int    `gorm:"id"`
	Title      string `gorm:"title"`
	Tags       string `gorm:"tags"`
	Short      string `gorm:"short"`
	Content    string `gorm:"content"`
	Author     string `gorm:"author"`
	CreateTime int64  `gorm:"create_time"`
	UpdateTime int64  `gorm:"updata_time"`
	Status     int    `gorm:"status"` //Status=0为正常，1为删除，2为冻结
}
