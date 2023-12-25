package service

import (
	"gin-bubble/common"
	"gin-bubble/model"
	"gorm.io/gorm"
	"time"
)

type IBubbleService interface {
	GetBubbleList(userId string) ([]*model.Todo, error)
	AddBubble(bubble *model.Todo) error
	DelBubble(id string) error
	UpdateBubble(id string, status string) error
}

type BubbleService struct {
	DB *gorm.DB
}

func (b BubbleService) UpdateBubble(id string, status string) error {
	// 获取当前时间
	// 设置新增时间--修改时间
	time := time.Now().Format("2006-01-02 15:04:05")
	err := b.DB.Debug().Model(&model.Todo{}).Where("id = ?", id).
		Update("status", status).Update("time", time).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BubbleService) DelBubble(id string) error {
	err := b.DB.Where("id = ?", id).Delete(&model.Todo{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BubbleService) AddBubble(bubble *model.Todo) error {
	err := b.DB.Debug().Create(bubble).Error
	if err != nil {
		return err
	}
	return nil
}

func (b BubbleService) GetBubbleList(userId string) ([]*model.Todo, error) {
	var bubbleList []*model.Todo
	err := b.DB.Where("user_id=?", userId).Find(&bubbleList).Error
	if err != nil {
		return nil, err
	}
	return bubbleList, nil
}

func NewBubbleService() IBubbleService {
	return BubbleService{DB: common.GetDB()}
}
