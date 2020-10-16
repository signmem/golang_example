

func DoLogRoate(logFilePath string, force bool) bool {
	// logFilePath := g.Config().Logfile
	backupSize := 5

	info, err := os.Stat(logFilePath)
	if err != nil {
		return false
	}
	fileSize := info.Size()

	if force == false {
		if fileSize < 10000000 {
			return true
		}
	}

	for i:= backupSize ; i > 0 ; i-- {
		if i > 1 {
			backupNumber := strconv.Itoa(i)
			originInt := i - 1
			originNumber := strconv.Itoa(originInt)
			backupFile := logFilePath + "." + backupNumber
			originFile := logFilePath + "." + originNumber
			_, err := os.Stat(originFile)

			if err != nil {
				continue
			}
			err = os.Rename(originFile, backupFile)
			if err != nil {
				log.Println("backup file error", err)
				return  false
			}
		}
	}

	backupFile := logFilePath + ".1"

	input, err := ioutil.ReadFile(logFilePath)
	if err != nil {
		log.Println("logfile read err ", err)
		return false
	}
	err = ioutil.WriteFile(backupFile, input, 0664)
	if err != nil {
		log.Println("create logfile write err ")
		return false
	}

	err = os.Truncate(logFilePath, 0)
	if err != nil {
		log.Println("truncate logfile error")
		return false
	}
	return true
}
