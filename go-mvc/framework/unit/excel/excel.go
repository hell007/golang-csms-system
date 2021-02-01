package main

//https://www.jianshu.com/p/764ceb9315eb
//https://www.jianshu.com/p/a3a6847756aa

import (
	"github.com/kataras/golog"
	"github.com/tealeg/xlsx"
	"go-mvc/framework/conf"
	"strconv"
)

// 导入
func Import(inFile string) {
	// 打开文件
	xlFile, err := xlsx.OpenFile(inFile)
	if err != nil {
		golog.Errorf("excel 打开文件出错：[%s]", err)
		return
	}

	// 遍历sheet页读取
	for _, sheet := range xlFile.Sheets {
		golog.Infof("sheet name: [%s]", sheet.Name)
		//遍历行读取
		for _, row := range sheet.Rows {
			// 遍历每行的列读取
			for _, cell := range row.Cells {
				text := cell.String()
				golog.Infof("[%s]", text)
			}
			golog.Info("\n")
		}
	}
	golog.Info("导入成功")
}

// 导出
func Export(outFile string, title string) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(title)
	if err != nil {
		golog.Errorf("导出出错：[%s]", err)
	}

	// 结构体
	students := getStudents()
	for _, stu := range students {
		row := sheet.AddRow()
		row.SetHeightCM(1) //设置每行的高度

		nameCell := row.AddCell()
		nameCell.Value = stu.Name

		avatarCell := row.AddCell()
		avatarCell.Value = stu.Avatar

		ageCell := row.AddCell()
		ageCell.Value = strconv.Itoa(stu.age)

		phoneCell := row.AddCell()
		phoneCell.Value = stu.Phone

		genderCell := row.AddCell()
		genderCell.Value = stu.Gender

		mailCell := row.AddCell()
		mailCell.Value = stu.Mail
	}

	err = file.Save(outFile)
	if err != nil {
		golog.Errorf("导出保存出错：[%s]", err)
	}
	golog.Info("导出成功！")
}

type Student struct {
	Name   string
	Avatar string
	age    int
	Phone  string
	Gender string
	Mail   string
}

func getStudents() []Student {
	students := make([]Student, 0)
	for i := 0; i < 10; i++ {
		stu := Student{}
		stu.Name = "name" + strconv.Itoa(i+1)
		stu.Avatar = ""
		stu.Mail = stu.Name + "@chairis.cn"
		stu.Phone = "1380013800" + strconv.Itoa(i)
		stu.age = 20
		stu.Gender = "男"
		students = append(students, stu)
	}
	return students
}

func main() {
	//导入
	inFile := conf.GetUploadFile() + "test.xlsx"
	Import(inFile)

	//导出
	outFile := conf.GetUploadFile() + "test.xlsx"
	Export(outFile, "学生表导出测试")
}
