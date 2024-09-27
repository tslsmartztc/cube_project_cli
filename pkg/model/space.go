package model

type GetSpaceAnCestorsRespEle struct {
	AncestorsSpace []*GetSpaceAnCestorsRespEle   `json:"ancestorsSpace"`
	Ps             GetSpaceAnCestorsRespEleSpace `json:"PS"`
}

type GetSpaceAnCestorsRespEleSpace struct {
	Code            string `json:"code"`
	CreateAt        string `json:"createAt"`
	Description     string `json:"description"`
	Extra           string `json:"extra"`
	GenceVertex     string `json:"genceVertex"`
	Height          int    `json:"height"`
	ID              string `json:"id"`
	IndexId         int    `json:"indexId"`
	IsDeleted       int    `json:"isDeleted"`
	Latitude        float64    `json:"latitude"`
	Longitude       float64    `json:"longitude"`
	MapServiceParam string `json:"mapServiceParam"`
	Name            string `json:"name"`
	ParentID        string `json:"parentID"`
	Physical        int    `json:"physical"`
	ProjID          string `json:"projID"`
	Type            string `json:"type"`
	UpdateAt        string `json:"updateAt"`
}
