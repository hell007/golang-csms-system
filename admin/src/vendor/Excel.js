/*
 * @Descripttion: 基于xlsx的Excel封装
 * 参考地址：https://github.com/SheetJS/sheetjs
 * @Author: zenghua.wang
 * @Date: 2019-08-04 16:14:40
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2021-07-22 10:34:12
 */
import fs from 'file-saver'
import XLSX from 'xlsx'


const s2ab = s => {
  if (typeof ArrayBuffer !== 'undefined') {
    var buf = new ArrayBuffer(s.length)
    var view = new Uint8Array(buf)
    for (var i = 0; i != s.length; ++i) view[i] = s.charCodeAt(i) & 0xff
    return buf
  } else {
    var buf = new Array(s.length);
    for (var i = 0; i != s.length; ++i) buf[i] = s.charCodeAt(i) & 0xFF;
    return buf;
  }
}


/**
 * 将json数据转换为excel
 * @param {object} fields { name: "姓名", mobile: "手机号" };
 * @param {array}  data   [{name: '张三', mobile: '138xxx'}]
 * @param {string} fileName 
 */
export function json2Excel(fields, json, fileName = "测试数据") {
  let autoWidth = true
  //删除多余不显示字段,表头显示中文
  let data = json
  data.forEach(item => {
    for (let i in item) {
      if (fields.hasOwnProperty(i)) {
        item[fields[i]] = item[i];
      }
      delete item[i];
    }
  })

  let sheetName = "sheet1";
  //工作簿对象包含一SheetNames数组，以及一个表对象映射表名称到表对象。XLSX.utils.book_new实用函数创建一个新的工作簿对象。
  let wb = XLSX.utils.book_new();
  //将JS对象数组转换为工作表。
  let ws = XLSX.utils.json_to_sheet(data, { header: Object.values(fields) });

  //单元格合并
  // 设置工作表的记录范围
  // [列号][行号]，A1 则代表 A 列的第1行
  // 列数一般是已知的（未知时可以设置为ZZ）
  // 行数则以 data 内容的长度结束即可
  // ws['!ref'] = `A1:AI${data.length}`;
  // s 意为 start ，即开始的单元格
  // r 是 row ，表示行号，从 0 计起
  // c 是 col ，表示列号，从 0 计起
  // var merge = [
  //   // 纵向合并，范围是第1列的行1到行2
  //   { s: { r: 0, c: 0 }, e: { r: 1, c: 0 } },
  //   // 横向合并，范围是第1行的列30到列35
  //   { s: { r: 0, c: 29 }, e: { r: 0, c: 34 } }
  // ];
  // ws['!ref'] = `A1:AI${data.length + 1}`;
  // var merge = [{
  //   s: { r: 1, c: 0 },
  //   e: { r: 3, c: 0 }
  // }];
  // ws['!merges'] = merge;

  //设置每列的最大宽度
  if (autoWidth) {
    const colWidth = Object.values(json[0]).map((val, index) => {
      if (!val) {
        return {
          'wch': 10
        };
      } else if (val.toString().charCodeAt(0) > 255) {
        return {
          'wch': val.toString().length * 3
        };
      } else {
        let len = Object.values(fields)[index].toString().length || 1;
        return {
          'wch': val.toString().length > len ? val.toString().length * 3 : len * 3,
        };
      }
    });
    ws['!cols'] = colWidth;
  }

  //设置表格的样式
  const sheetStyle = {
    font: { name: "宋体", sz: 14, color: "#333", bold: true },
    alignment: {
      horizontal: "center",
      vertical: "center"
    },
    fill: { fgColor: { rgb: "green" } }
  };

  //将sheet添加到工作簿
  wb.SheetNames.push(sheetName);
  wb.Sheets[sheetName] = ws;

  let wopts = { bookType: 'xlsx', bookSST: false, type: 'binary', cellStyles: true, defaultCellStyle: sheetStyle, showGridLines: false };
  let wbout = XLSX.write(wb, wopts);
  let blob = new Blob([s2ab(wbout)], { type: 'application/octet-stream' });
  fs.saveAs(blob, fileName + ".xlsx");
}

/**
 * 将<table></table>内容导入excel
 * @param {*} dom 
 * @param {*} sheetName 
 */
export function table2Excel(dom, fileName = "测试数据") {
  let table = document.querySelector(dom)
  let wb = XLSX.utils.book_new();
  let ws = XLSX.utils.table_to_sheet(table);
  XLSX.utils.book_append_sheet(wb, ws, "Sheet1");
  XLSX.writeFile(wb, fileName + ".xlsx");
}