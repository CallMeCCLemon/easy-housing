package cmd

import "latentlab.cc/easyhousing/pkg/service"

func main() {
	s := service.NewEasyHousingService()
	s.Start()
}
