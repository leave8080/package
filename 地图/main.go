package main

import (
	"fmt"
	"github.com/qichengzx/coordtransform"
)

func main() {
	fmt.Println(coordTransform.BD09toGCJ02(116.404, 39.915))
	fmt.Println(coordTransform.GCJ02toBD09(116.404, 39.915))
	fmt.Println(coordTransform.WGS84toGCJ02(116.404, 39.915))
	fmt.Println(coordTransform.GCJ02toWGS84(116.404, 39.915))
	fmt.Println(coordTransform.BD09toWGS84(116.404, 39.915))
	fmt.Println(coordTransform.WGS84toBD09(116.404, 39.915))
}
