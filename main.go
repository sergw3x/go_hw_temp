package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Temp struct {
	Type string
	Name string
	Temp int
}

func main() {
	if !checkKernelModule(){
		fmt.Println(		`
		Для считывания температуры hdd необходимо подключить модуль ядра 'drivetemp'		
		  sudo modprobe drivetemp
		для проверки используйте: lsmod | grep drivetemp
		Добавить в автозагрузку можно, указав модуль и его параметры в файле внутри каталога /etc/modules-load.d. 
		Файлы должны заканчиваться .confи могут иметь любое имя:
		  /etc/modules-load.d/module_name.conf
		  option module_name parameter=value
		либо просто написать имя модуля
		  'drivetemp'
		`)
	}

	fmt.Printf("%s", GetTemp())

}

func GetTemp() []Temp {
	dt := []Temp{}

	files, err := filepath.Glob("/sys/class/hwmon/hwmon*")
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range files {
		name, err := ioutil.ReadFile(device + "/name")
		if err != nil {
			continue
		}
		sName := strings.TrimRight(string(name), "\n")
		if sName == "drivetemp" {
			model, err := ioutil.ReadFile(device + "/device/model")
			if err != nil {
				log.Panic(err)
			}
			modelName := fmt.Sprintf("%s", model)
			temp1_input, err := ioutil.ReadFile(device + "/temp1_input")
			if err != nil {
				log.Panic(err)
			}
			dt = append(dt, Temp{"drive", modelName, getFormatedTemperature(temp1_input)})
		} else if sName == "coretemp" {
			temp1_input, err := ioutil.ReadFile(device + "/temp1_input")
			if err != nil {
				log.Panic(err)
			}
			dt = append(dt, Temp{"cpu", "cpu", getFormatedTemperature(temp1_input)})
		} else {
			continue
		}
	}
	return dt
}

func getFormatedTemperature(s []uint8) int {
	temp, err := strconv.Atoi(fmt.Sprintf("%-.2s", string(s)))
	if err != nil {
		log.Panic(err)
	}
	return temp
}

func checkKernelModule() bool {
	cmd := exec.Command("bash", "-c", "lsmod | grep drivetemp")
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}