Для считывания температуры hdd необходимо подключить модуль ядра 'drivetemp'		
		  sudo modprobe drivetemp
		для проверки используйте: lsmod | grep drivetemp
		Добавить в автозагрузку можно, указав модуль и его параметры в файле внутри каталога /etc/modules-load.d. 
		Файлы должны заканчиваться .confи могут иметь любое имя:
		  /etc/modules-load.d/module_name.conf
		  option module_name parameter=value
		либо просто написать имя модуля
		  'drivetemp'