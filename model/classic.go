package model

import (
	"github.com/Away0x/7yue_api_server/constant"
	"github.com/Away0x/7yue_api_server/constant/errno"
	"github.com/jinzhu/gorm"
)

type Classic struct {
	gorm.Model
	Content string `gorm:"type:varchar(255);not null;column:content"`
	Image   string `gorm:"type:varchar(128);not null;column:image"`
	Url     string `gorm:"type:varchar(128);not null;column:url"`
	Index   int    `gorm:"type:integer;not null;column:index"`
	Title   string `gorm:"type:varchar(128);not null;column:title"`
	Type    int    `gorm:"type:integer;not null;column:type"`
}

func (Classic) TableName() string {
	return "classic"
}

func (c *Classic) Create() error {
	return DB.Create(&c).Error
}

// 获取 index 最后的 classic
func LatestClassic() (*Classic, error) {
	c := &Classic{}
	d := DB.Order("index").Find(&c).Limit(1)
	return c, d.Error
}

// 获取指定 index 的下一期 classic
func FindNextClassic(index int) (*Classic, error) {
	c := &Classic{Index: index + 1}
	d := DB.Where(c).First(&c)
	return c, d.Error
}

// 获取指定 index 的前一期 classic
func FindPreviousClassic(index int) (*Classic, error) {
	if index == 1 {
		return nil, errno.NoClassicError
	}
	c := &Classic{Index: index - 1}
	d := DB.Where(c).First(&c)
	return c, d.Error
}

// 获取指定 id 和 type 的数据
func GetClassicDetail(id int, _type int) (*Classic, error) {
	c := &Classic{}
	d := DB.Where("id = ? AND type = ?", id, _type).First(&c)
	return c, d.Error
}

// 获取我喜欢的期刊
func GetFavorClassices(user_key string, page int, count int) ([]*Classic, error) {
	classices := make([]*Classic, count)
	types := []int{constant.ClASSIC_TYPE_MOVIE, constant.ClASSIC_TYPE_MUSIC, constant.ClASSIC_TYPE_SENTENCE}
	_, ids, err := GetUserFavorList(user_key, types)
	if err != nil {
		return classices, err
	}

	offset_count := (page - 1) * count
	d := DB.Where("id in (?)", ids).Limit(count).Offset(offset_count).Find(&classices)
	return classices, d.Error
}

type ClassicSerializer struct {
	Id         uint   `json:"id"`
	Pubdate    string `json:"pubdate"`
	Content    string `json:"content"`
	FavNums    int    `json:"fav_nums"`
	Image      string `json:"image"`
	Url        string `json:"url"`
	Index      int    `json:"index"`
	LikeStatus int    `json:"like_status"`
	Title      string `json:"title"`
	Type       int    `json:"type"`
}

func (c *Classic) Serializer(path string, like_status int, fav_nums int) ClassicSerializer {
	s := ClassicSerializer{
		Id:         c.ID,
		Pubdate:    c.CreatedAt.Format("2006-01-02"),
		Content:    c.Content,
		FavNums:    fav_nums,
		Image:      path + "/static/images/" + c.Image,
		Url:        c.Url,
		Index:      c.Index,
		LikeStatus: like_status,
		Title:      c.Title,
		Type:       c.Type,
	}

	return s
}
