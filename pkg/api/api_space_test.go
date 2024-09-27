package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	pkg "github.com/tslsmartztc/cube_project_cli/pkg/api"
)

func Test_GetSpaceAncestors(t *testing.T) {
	p, err := pkg.NewClient("10.37.126.117", 10315, pkg.ClientOpt{BaseUrl: "cubeProject"}).
		GetSpaceAncestors(pkg.CommonHeader{
			Authorization:   "a82f8d5ff377487788b255f8962771f9",
			AuthContextCode: "1",
		}, 1, 50, "TSL20230822A0001", "1608783896353185792", 10)

	assert.NoError(t, err)
	assert.NotEmpty(t, p)
}
