package model

import "time"

type Project struct {
	ID                   string     `json:"id"`
	AuthCtxCode          string     `json:"authCtxCode"`
	Code                 string     `json:"code"`
	Name                 string     `json:"name"`
	Principal            string     `json:"principal"`
	PrincipalPersonPhone string     `json:"principalPersonPhone"`
	Manager              string     `json:"manager"`
	ManagerPhone         string     `json:"managerPhone"`
	Description          string     `json:"description"`
	Status               int        `json:"status"`
	IsDeleted            int        `json:"isDeleted"`
	Region               string     `json:"region"`
	RegionId             string     `json:"regionId"`
	Longitude            float64    `json:"longitude"`
	Latitude             float64    `json:"latitude"`
	StartAt              string     `json:"startAt"`
	CreateAt             *time.Time `json:"createAt"`
	UpdateAt             *time.Time `json:"updateAt"`
	Logo                 string     `json:"Logo"`
	SamplePic            string     `json:"samplePic"`
	IotProjID            string     `json:"iotProjId"  ` // iot项目ID
}
