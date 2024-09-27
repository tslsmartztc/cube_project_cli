package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/tslsmartztc/cube_project_cli/internal"
	"github.com/tslsmartztc/cube_project_cli/pkg/model"
)

func (s *Client) GetSpaceAncestors(header CommonHeader, page, pageSize int, projCode, spaceID string, ttl int) ([]*model.GetSpaceAnCestorsRespEle, error) {
	url := s.httpUrl("spaceGetAncestors")

	bd := map[string]any{
		"projCode": projCode,
		"spaceId":  spaceID,
		"ttl":      ttl,
		"page":     page,
		"pageSize": pageSize,
	}

	resp, err := internal.NewSimpleHttpRequest[[]*model.GetSpaceAnCestorsRespEle]().
		SetMethod(http.MethodPost).
		SetHeaders([][]string{
			{"Authorization", header.Authorization},
			{"AuthContextCode", header.AuthContextCode}}).
		SetUrl(url).
		SetJsonBody(bd).
		Do()
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf(resp.Message)
	}
	if resp.Data == nil || len(resp.Data) == 0 {
		return nil, errors.New("no data found")
	}
	return resp.Data, nil
}
