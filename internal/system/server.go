package system

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"log"
	"time"
)

type MemData struct {
	Name       string  // 名称
	Value      float64 // 数值
	Percentage float64 // 百分比
}

// 获取内存信息
func GetMemoryInfo() {
	v, err := mem.VirtualMemory()
	if err != nil {
		memMsg := fmt.Sprintf("[get memory data] error=%v", err)
		log.Fatalln(memMsg)
	}
	fmt.Println(v)

	// 将内存单位转为 G
	memTotalG := float64(v.Total/(1<<30) + 1)
	fmt.Printf("总内存: %vG \n", memTotalG)

	used := MemData{
		Name:       "已使用内存",
		Value:      memTotalG * v.UsedPercent / 100,
		Percentage: v.UsedPercent,
	}

	freePercent := 100.0 - used.Percentage
	free := MemData{
		Name:       "未使用内存",
		Value:      memTotalG * freePercent / 100,
		Percentage: freePercent,
	}

	fmt.Println(used, free)

}

// 获取cpu信息
func GetCpuInfo() {
	cpuList, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		memMsg := fmt.Sprintf("[get cpu data] error=%v", err)
		log.Fatalln(memMsg)
	}
	fmt.Println(cpuList)
}
