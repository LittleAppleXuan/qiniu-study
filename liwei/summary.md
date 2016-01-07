2016-01-04

学习Golang 编程基础:
- Golang 编程单元测试用例学习
- Golang 包引用学习
- github工作流使用、git基础命令学习掌握
- 发布Golang library 到 github
- SublimeText golang 开发环境搭建
- 学习Golang 单元测试库 assert的用法

---------------------------------

2016-01-05

学习Golang 
- 字符串 操作
- 数字类型
- 数组
     - 无符号数: uint8, uint16, uint32, uint64
     - 有符号数: int8, int16, int32, int64
     - 浮点数: float32, float64
     - 复数: complex64, complex128
     - 其它: byte(uint8 的别名), rune(int32 的别名，用来存储 Unicode)
- 切片
     - 底层为数组，但是可变长
     - 比数组更为实用
- map
     - make() 内建方法只能用于创建 slice, map 和 channel 
     - slice, map, channel 是指向底层数据结构的引用，使用前必须被初始化
- 指针
     - Go 有指针，但是没有指针运算。你不能 用指针变量遍历字符串的各个字节。
     - Go 指针只是一种引用。
     - 取址操作符 & 获取变量存放地址，可以赋给一个指针。
- 结构体
     * 创建自己定义的类型
     * 方法定义
     * 匿名字段
     * 方法继承
     * 方法覆盖
     * 给类型 B 添加一个方法覆盖类型 S 中的 String()
     * 字节对齐
2016-01-05
------------------

#### 学习golang 分支语句、循环控制语句、文件操作、异常处理等基础知识
* goto
忽略标签而不是后面的语句

* if

作用域

* switch

支持初始化

默认 break

fallthough

- 错误处理 panic recover

- 限速传输、文件读写、无写权限：

