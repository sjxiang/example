package main

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	defaulShelfSize = 100
)


// 定义 Model

// Shelf 书架
type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

func (s Shelf) TableName() string {
	return "shelf"
}


// Book 图书
type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

func (b Book) TableName() string {
	return "book"
}


// 数据库操作
type Store interface {
	CreateShelf(ctx context.Context, data Shelf) (*Shelf, error)
	GetShelf(ctx context.Context, id int64) (*Shelf, error)
	ListShelves(ctx context.Context) ([]*Shelf, error)
	DeleteShelf(ctx context.Context, id int64) error
	GetBookListByShelfID(ctx context.Context, shelfID int64, cursor string, pageSize int) ([]*Book, error)
}

type StoreImpl struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return &StoreImpl{
		db: db,
	}
}


// CreateShelf 创建书架
func (impl *StoreImpl) CreateShelf(ctx context.Context, data Shelf) (*Shelf, error) {
	// 缺省设置
	size := data.Size
	if size <= 0 {
		size = defaulShelfSize
	}
	now := time.Now()
	item := Shelf{Theme: data.Theme, Size: size, CreateAt: now, UpdateAt: now}
	err := impl.db.WithContext(ctx).Create(&item).Error
	return &item, err
}

// GetShelf 获取书架
func (impl *StoreImpl) GetShelf(ctx context.Context, id int64) (*Shelf, error) {
	item := Shelf{}
	err := impl.db.WithContext(ctx).First(&item, id).Error
	return &item, err
}

// ListShelves 书架列表
func (impl *StoreImpl) ListShelves(ctx context.Context) ([]*Shelf, error) {
	var items []*Shelf
	err := impl.db.WithContext(ctx).Find(&items).Error
	return items, err
}

// DeleteShelf 删除书架
func (impl *StoreImpl) DeleteShelf(ctx context.Context, id int64) error {
	return impl.db.WithContext(ctx).Delete(&Shelf{}, id).Error
}

// GetBookListByShelfID 根据书架id查询图书
func (impl *StoreImpl) GetBookListByShelfID(ctx context.Context, shelfID int64, cursor string, pageSize int) ([]*Book, error) {
	var items []*Book
	err := impl.db.Debug().WithContext(ctx).Where("shelf_id = ? and id > ?", shelfID, cursor).Order("id asc").Limit(pageSize).Find(&items).Error
	return items, err
}
