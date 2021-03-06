package httphoneypot

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RequestMySQLModel struct {
	gorm.Model
	Method  string         `gorm:"size:256"`
	URL     string         `gorm:"size:2048"`
	Proto   string         `gorm:"size:256"`
	Headers datatypes.JSON `json:"headers"`
	Body    string         `json:"body"`
	IP      string         `gorm:"size:256"`
}

func (model *RequestMySQLModel) TableName() string {
	return "http_requests"
}

var _ RequestRepository = (*MySQLRequestRepository)(nil)

type MySQLRequestRepository struct {
	db *gorm.DB
}

func NewMySQLRequestRepository(conf *MySQLConfig) (*MySQLRequestRepository, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Hostname, conf.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&RequestMySQLModel{})

	return &MySQLRequestRepository{db: db}, err
}

func (repo *MySQLRequestRepository) Create(req *Request) error {
	headers, err := json.Marshal(req.Headers)
	if err != nil {
		return errors.WithStack(err)
	}
	model := RequestMySQLModel{
		Method:  req.Method,
		URL:     req.URL,
		Proto:   req.Proto,
		Headers: headers,
		Body:    req.Body,
		IP:      req.IP,
	}
	return repo.db.Create(&model).Error
}
