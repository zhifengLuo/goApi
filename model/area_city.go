package model

type AreaCity struct {
	// ID uint   `json:"id" gorm:"primary_key"`
	Code string
	Name string
}

// var AreaCityModel = new(AreaCity)

// 指定表名
// func (u AreaCity) TableName() string {
// 	return "are"
// }

func GetAllProvince() (data []*AreaCity) {
	db.Where("code like '__0000'").Order("code").Find(&data)
	return
}

func GetByCode() {

}
