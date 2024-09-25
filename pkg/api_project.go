package pkg

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/tslsmartztc/cube_project_cli/internal"
	"github.com/tslsmartztc/cube_project_cli/pkg/model"
)

func (s *Client) GetProjWithCode(header CommonHeader, code string) (*model.Project, error) {
	url := s.httpUrl("getProjWithCode")
	url = s.appendUrlParams(url, "projCode", code)
	resp, err := internal.NewSimpleHttpRequest[*model.Project]().
		SetMethod(http.MethodGet).
		SetHeaders([][]string{{"Authorization", header.Authorization}}).
		SetUrl(url).
		Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get proj by id: %w", err)
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf(resp.Message)
	}
	return resp.Data, nil // return proj pointer if no error
}

func (s *Client) GetProjWithID(header CommonHeader, pid string) (*model.Project, error) {
	url := s.httpUrl("getProject")
	url = s.appendUrlParams(url, "id", pid)
	resp, err := internal.NewSimpleHttpRequest[*model.Project]().
		SetMethod(http.MethodGet).
		SetHeaders([][]string{{"Authorization", header.Authorization}}).
		SetUrl(url).
		Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get proj by id: %w", err)
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf(resp.Message)
	}
	return resp.Data, nil // return proj pointer if no error
}

func (s *Client) GetProjsWithName(header CommonHeader, name string) ([]*model.Project, error) {
	url := s.httpUrl("getProjWithName")
	url = s.appendUrlParams(url, "projName", name)

	resp, err := internal.NewSimpleHttpRequest[[]*model.Project]().
		SetMethod(http.MethodGet).
		SetHeaders([][]string{{"Authorization", header.Authorization}}).
		SetUrl(url).
		Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get proj by id: %w", err)
	}
	if resp.Code == 401009 && resp.Message == "获取的项目为空" {
		return nil, errors.New("no data found")
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf(resp.Message)
	}
	return resp.Data, nil // return proj pointer if no error
}

func (s *Client) GetProjWithAuthCtx(header CommonHeader, authCtx []string) ([]*model.Project, error) {
	// url := fmt.Sprintf("http://%s%s", client, s.urlGetProjWithAuthCtx)
	url := s.httpUrl("projects/getWithAuthContextCodes")

	resp, err := internal.NewSimpleHttpRequest[[]*model.Project]().
		SetMethod(http.MethodPost).
		SetHeaders([][]string{{"Authorization", header.Authorization}}).
		SetUrl(url).
		SetJsonBody(authCtx).
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
