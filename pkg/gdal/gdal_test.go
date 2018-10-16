package gdal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGdalInfo(t *testing.T) {
	fname := "/vsis3/landsat-pds/L8/210/244/LC82102442014232LGN00/LC82102442014232LGN00_B1.TIF"
	gir, err := GDALInfo(fname)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, fname, gir.Description)
	assert.Equal(t, "GTiff", gir.DriverShortName)
	assert.Equal(t, "GeoTIFF", gir.DriverLongName)
	assert.NotEmpty(t, gir.Files)
}
