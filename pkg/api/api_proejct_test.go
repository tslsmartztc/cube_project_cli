package pkg_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	pkg "github.com/tslsmartztc/cube_project_cli/pkg/api"
)

func Test_GetProjWithCode(t *testing.T) {
	p, err := pkg.NewClient("10.37.126.117", 10315, pkg.ClientOpt{BaseUrl: "cubeProject"}).GetProjWithCode(pkg.CommonHeader{
		Authorization:   "a82f8d5ff377487788b255f8962771f9",
		AuthContextCode: "1",
	}, "TSL20240531A0621")
	assert.NoError(t, err)
	assert.NotEmpty(t, p)
}

func Test_GetProjWithID(t *testing.T) {
	p, err := pkg.NewClient("10.37.126.117", 10315, pkg.ClientOpt{BaseUrl: "cubeProject"}).GetProjWithID(pkg.CommonHeader{
		Authorization:   "a82f8d5ff377487788b255f8962771f9",
		AuthContextCode: "1",
	}, "1796467071186378752")
	assert.NoError(t, err)
	assert.NotEmpty(t, p)
}

func Test_GetProjWithName(t *testing.T) {
	p, err := pkg.NewClient("10.37.126.117", 10315, pkg.ClientOpt{BaseUrl: "cubeProject"}).GetProjsWithName(pkg.CommonHeader{
		Authorization:   "a82f8d5ff377487788b255f8962771f9",
		AuthContextCode: "1",
	}, "Smart")
	assert.NoError(t, err)
	assert.NotEmpty(t, p)
}

func Test_GetProjWithAuthContext(t *testing.T) {
	p, err := pkg.NewClient("10.37.126.117", 10315, pkg.ClientOpt{BaseUrl: "cubeProject"}).GetProjWithAuthCtx(pkg.CommonHeader{
		Authorization:   "a82f8d5ff377487788b255f8962771f9",
		AuthContextCode: "1",
	}, []string{"proj2g8TGPYxeMJ4pJULBfsU6raMhOm"})
	assert.NoError(t, err)
	assert.NotEmpty(t, p)
}
