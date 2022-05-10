package main

import "gorm.io/gorm"

type Epub struct {
	gorm.Model
	Title    string
	Descript string
	Path     string
	ISBN     string
}

func (e *Epub) Insert() error {
	return store.Create(e).Error
}
