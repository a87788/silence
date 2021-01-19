package main

import (
	"archive/zip"
	"bufio"
	"database/sql"
	"fmt"
	"github.com/go-ini/ini"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//zip文件解压
func UnZip(dst1, src string) (err error) {
	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	dst := strings.Join([]string{dst1, ":\\"}, "")
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		return
	}

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	// 遍历 zr ，将文件写入到磁盘
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}

		// 获取到 Reader
		fr, err := file.Open()
		if err != nil {
			return err
		}

		// 创建要写出的文件对应的 Write
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		n, err := io.Copy(fw, fr)
		if err != nil {
			return err
		}

		// 将解压的结果输出
		fmt.Printf("成功解压 %s ，共写入了 %d 个字符的数据\n", path, n)

		// 因为是在循环中，无法使用 defer ，直接放在最后
		// 不过这样也有问题，当出现 err 的时候就不会执行这个了，
		// 可以把它单独放在一个函数中，这里是个实验，就这样了
		_ = fw.Close()
		_ = fr.Close()

	}
	return nil
}

//修改配置文件
func modify(dst, port string) {
	//拼接安装路径
	dir := strings.Join([]string{dst, ":\\", "mysql-5.7.26-winx64\\my.ini"}, "")
	fmt.Print(dir)
	cfg, err := ini.Load(dir)
	if err != nil {

		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	cfg.Section("mysqld").Key("port").SetValue(port)
	if dst != "D" {
		cfg.Section("mysqld").Key("basedir").SetValue(strings.Join([]string{dst, ":\\mysql-5.7.26-winx64"}, ""))
		cfg.Section("mysqld").Key("datadir").SetValue(strings.Join([]string{dst, ":\\mysql-5.7.26-winx64\\data"}, ""))
	}
	//fmt.Println("mysql的安装目录为:", cfg.Section("mysqld").Key("basedir").String())
	//fmt.Println("mysql的数据目录为:", cfg.Section("mysqld").Key("datadir").String())
	//fmt.Println("mysql端口为:", cfg.Section("mysqld").Key("port").String())
	_ = cfg.SaveTo(dir)
}

//mysql初始化
func mysql(dst string) {
	//初始化
	scr := strings.Join([]string{dst, ":\\mysql-5.7.26-winx64\\bin\\mysqld "}, "")
	cmd := exec.Command("cmd", "/c", scr, "--initialize")
	err := cmd.Run()
	fmt.Println("Error: ", err)
	cmd2 := exec.Command("cmd", "/c", scr, "install")
	err = cmd2.Run()
	fmt.Println("Error: ", err)
	//获取随机密码文件
	passrc := strings.Join([]string{dst, ":\\mysql-5.7.26-winx64\\data"}, "")
	dir_list, e := ioutil.ReadDir(passrc)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	var errname string
	for _, v := range dir_list {
		if strings.Contains(v.Name(), ".err") == true {
			errname = v.Name()
			break
		}
	}

	//获取随机密码
	src := strings.Join([]string{passrc, errname}, "")
	fmt.Println(src)
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		aa := scanner.Text()
		if strings.Contains(aa, "root@localhost:") == true {
			res := strings.Split(aa, "root@localhost:")
			//随机密码
			restul := strings.Replace(res[1], " ", "", 1)
			fmt.Print(restul)
		}
	}
	//连接数据库
	db, err := sql.Open("mysql", "root:123456@(1localhost:3306)?charset=utf8")
	if err != nil {
		panic(err)
	}
	fmt.Println(db.Ping()) //检查是否连接成功数据库
}

//执行
func main() {

	//文件路径
	var src string
	fmt.Print("请将文件夹拖入到此窗口：")
	_, _ = fmt.Scan(&src)
	//
	// 解压后保存的位置
	var dst string
	fmt.Print("请输入要安装到哪个盘：")
	_, _ = fmt.Scan(&dst)
	_ = UnZip(dst, src)
	// 修改配置文件
	var port string
	fmt.Print("请输入数据库安装端口：")
	_, _ = fmt.Scan(&port)
	modify(dst, port)
	//初始化
	mysql(dst)

}
