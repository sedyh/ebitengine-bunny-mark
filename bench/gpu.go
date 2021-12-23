package bench

import (
	"github.com/jaypipes/ghw"
)

func GpuInfo() (gpu string) {
	gpu = "Unknown GPU"
	info, err := ghw.GPU()
	if err != nil {
		return
	}
	if len(info.GraphicsCards) == 0 {
		return
	}
	return info.GraphicsCards[0].DeviceInfo.Product.Name
}
