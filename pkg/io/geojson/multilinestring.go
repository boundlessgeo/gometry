/*
MIT License

Copyright (c) 2018 Boundless

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package geojson

import (
	"encoding/json"

	"github.com/boundlessgeo/gometry/pkg/geom"
)

type multiLineString struct {
	Typ        geoJSONType          `json:"type"`
	Coordinate [][]*geom.Coordinate `json:"coordinates"`
}

func marshalMultiLineString(l *geom.MultiLineString) ([]byte, error) {
	coords := make([][]*geom.Coordinate, len(l.LineStrings))
	for i, v := range l.LineStrings {
		coords[i] = v.Coordinates
	}
	return json.Marshal(&multiLineString{multiLineStringType, coords})
}
