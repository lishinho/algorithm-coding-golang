package bili

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type csvData struct {
	Mid  string
	Aid  string
	Time string
	Bid  string
	Cid  string
}

func generateCsv() {
	f, err := os.Create("test1.csv") //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	var datas [][]string
	for j := 32431240; j < 32541343; j++ {
		var d []string
		csv := &csvData{
			Mid:  strconv.Itoa(j),
			Aid:  strconv.Itoa(j - 385943),
			Time: strconv.FormatInt(time.Now().Unix(), 10),
			Bid:  "12454",
			Cid:  "222345",
		}
		d = append(d, csv.Mid, csv.Aid, csv.Time, csv.Bid, csv.Cid)
		datas = append(datas, d)
		time.Sleep(time.Millisecond * 1)
	}

	w := csv.NewWriter(f) //创建一个新的写入文件流
	w.WriteAll(datas)     //写入数据
	w.Flush()
}

func ReadCsv() {
	//准备读取文件
	fileName := "testCsv.csv"
	file, _ := os.Open(fileName)
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	popularWatchTimeMap := make(map[int64][]*PopularWatchTime)
	for _, line := range records {
		if len(line) <= 0 {
			continue
		}
		fmt.Println(line)
		mid, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			continue
		}
		aid, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			continue
		}
		timeStamp, err := strconv.ParseInt(line[2], 10, 64)
		if err != nil {
			continue
		}
		popularWatchTimeMap[mid] = append(popularWatchTimeMap[mid], &PopularWatchTime{
			Aid:  aid,
			Time: timeStamp,
		})
	}
	fmt.Println("--------------")
	for k, v := range popularWatchTimeMap {
		fmt.Printf("k %v :", k)
		for _, val := range v {
			fmt.Printf(" %v", val)
		}
		fmt.Println()
	}
}
