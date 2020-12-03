package models

import (
  "time"
)

type Book struct {
  ID	              uint	    `json:"id" gorm:"primary_key"`
  Title	            string    `json:"title"`
  Author            string    `json:"author"`
  LastModifiedTime  time.Time `json:"lastModifiedTime"` 
  LastModifiedUID   uint      `json:"lastModifiedUID"`
  CreateTime        time.Time `json:"createTime"`
  CreateUID         uint      `json:"createUID"`
  IsDeleted         bool      `json:"isDeleted"`
}