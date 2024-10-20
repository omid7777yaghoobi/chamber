package bucket

import (
	"time"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Bucket struct {
	Id   		int64 `gorm:"column:id"`
	Name 		string `gorm:"column:name"`
	Created_at 	time.Time `gorm:"column:created_at"`
	Owner		string `gorm:"column:owner"`

	// bucketPolicy     BucketPolicy
	// bucketVersioning BucketVersioning
	// BucketEncryption BucketEncryption
}

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})


func (b *bucket) ServeHTTP(r http.ResponseWriter, w *http.Request) {

}

func (b *bucket) CreateBucket() {
	db.Create(&Bucket{})
}

func (b *bucket) ListBuckets() []bucket {

}
