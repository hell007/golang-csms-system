package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kataras/golog"
	"go-mvc/framework/conf"
)

var myTableName = "成绩表"
var p1 = `{
	"border":[
		{
			"type":"left",
			"color":"000000",
			"style":1
		},
		{
			"type":"top",
			"color":"000000",
			"style":1
		},
		{
			"type":"bottom",
			"color":"000000",
			"style":1
		},
		{
			"type":"right",
			"color":"000000",
			"style":1
		}
	]
}`
var p2 = `{
	"alignment":{
		"horizontal":"center"
	},
	"border":[
		{
			"type":"left",
			"color":"000000",
			"style":1
		},
		{
			"type":"top",
			"color":"000000",
			"style":1
		},
		{
			"type":"bottom",
			"color":"000000",
			"style":1
		},
		{
			"type":"right",
			"color":"000000",
			"style":1
		}
	],
	"fill":{
		"type":"gradient",
		"color":[
			"#FFFFFF",
			"#FF4040"
		],
		"shading":0
	}
}`

func main() {
	//1.创建一个文件Excel文件
	f := excelize.NewFile()
	f.AddPicture()

	//2.创建一个工作表，表名为myTableName的内容
	index := f.NewSheet(myTableName)

	//3.设置第一个样式
	style, err := f.NewStyle(p1)
	//3.1.设置第二个样式,fill的颜色如果两个都是一样的，就是纯色了，type查看文档可以调出很多渐变样式
	style2, err := f.NewStyle(p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	//将第一个样式给行A1到C1，你也可以设置A1 A10这样，看需求
	f.SetCellStyle(myTableName, "A1", "C1", style)
	f.SetCellValue(myTableName, "A1", "员工工号")
	f.SetCellValue(myTableName, "B1", "员工姓名")

	//合并单元格
	f.MergeCell(myTableName, "D1", "E1")
	f.SetCellValue(myTableName, "D1", "合并的单元格")
	//设置合并单元格，每一格宽度是10,D+E=20
	f.SetColWidth(myTableName, "D", "E", 10)
	//设置C单元格20的宽度
	f.SetColWidth(myTableName, "C", "C", 20)
	//将第一行设置为20高度
	f.SetRowHeight(myTableName, 1, 20)

	//将第一个样式给行A2到C2，你也可以设置A1 A10这样，看需求
	f.SetCellStyle(myTableName, "A2", "C2", style)
	//将第二个样式给行A2到C2，你也可以设置A1 A10这样，看需求，从这里可以看出，可以设置多个style组合，以便少定义一些样式
	f.SetCellStyle(myTableName, "A2", "C2", style2)

	f.SetCellValue(myTableName, "A2", "XC10746")
	f.SetCellValue(myTableName, "B2", "春卷虎")
	f.SetCellValue(myTableName, "C2", "表头都没有")

	//4.设置当前索引的工作表为默认显示
	f.SetActiveSheet(index)
	//将文件另存为myTableName名.xlsx
	outFile := conf.GetUploadFile() + "360.xlsx"
	if err := f.SaveAs(outFile); err != nil {
		println(err.Error())
	}
	golog.Info("导出成功！")
}