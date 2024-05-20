package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nome      string         `json:"nome"`
	CPF       string         `json:"cpf"`
	Email     string         `json:"email"`
	Senha     string         `json:"senha"`
	IsAtivo   int            `json:"is_ativo"`
	CreatedAt *time.Time     `gorm:"column:created_at;autoCreateTime" json:"CreatedAt"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;autoUpdateTime" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"DeletedAt"`
}

func (u *Usuario) BeforeSave(tx *gorm.DB) (err error) {
	if u.Senha != "" {
		hashedSenha, err := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Senha = string(hashedSenha)
	}
	return nil
}

func (Usuario) TableName() string {
	return "usuario"
}
