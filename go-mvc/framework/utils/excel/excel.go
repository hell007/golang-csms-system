package excel

//https://www.jianshu.com/p/764ceb9315eb
//https://www.jianshu.com/p/a3a6847756aa

import (
	"github.com/kataras/golog"
	"github.com/tealeg/xlsx"
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
	row := sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell := row.AddCell()
	cell.Value = "测试测试"

	err = file.Save(outFile)
	if err != nil {
		golog.Errorf("导出保存出错：[%s]", err)
	}
	golog.Info("导出成功！")
}
