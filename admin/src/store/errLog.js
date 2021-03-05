const errLog = {
  state: {
    errLog: []
  },
  pushLog(log) {
    //向数组开头添加一个新元素并返回新的素组
    this.state.errLog.unshift(log)
  },
  clearLog() {
    this.state.errLog = []
  }
}

export default errLog