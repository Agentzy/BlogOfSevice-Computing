package main

import ("fmt"
		"bufio"
		"io"
		"os"
		"os/exec"
		"strconv"
		"github.com/spf13/pflag"
)

type selpg_args struct{
	start_page int
	end_page int   //开始页、结束页
	length_page int 			//一页长度
	page_type bool				//-l  -f ,输入文本类型
	file_in string
	file_out string	//输入输出文件
}

var progarm string 	//程序名

//参数值绑定
func analy_args(arg *selpg_args){
	pflag.IntVarP(&arg.start_page, "startPage", "s",0, "Start Page:")
	pflag.IntVarP(&arg.end_page, "endPage", "e",0, "end Page:")
	pflag.IntVarP(&arg.length_page, "lenPage", "l",20, "Line in each page:")
	pflag.StringVarP(&arg.file_out,"fileout","d","","output file name:")
	pflag.BoolVar(&(arg.page_type), "f", false, "type of print")
	//出错格式帮助
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"\nUSAGE: %s -s start_page -e end_page [ -f | -l lines_per_page ]" +
	 		" [ -d dest ] [ in_filename ]\n", progarm)
		pflag.PrintDefaults()
	}

    pflag.Parse()

	other := pflag.Args()
	if len(other) > 0{
		arg.file_in = other[0]//设置文件名
	}else{
		arg.file_in = ""
	}

	//fmt.Println(arg.start_page, arg.end_page, arg.file_in, arg.page_type)
	//参数少于3个
	if(len(os.Args) < 3){
		fmt.Fprintf(os.Stderr,"%s: need 3 args at least\n", progarm)
		pflag.Usage()
		os.Exit(1)
	}

	//第一个参数必须是-s
	if os.Args[1] != "-s"{
		fmt.Fprintf(os.Stderr,"%s:first arg need to be -s start_page\n", progarm)
		pflag.Usage()
		os.Exit(2)
	}

	//起始页码不能小与1
	sp,_ := strconv.Atoi(os.Args[2])
	if sp < 1 {
		fmt.Fprintf(os.Stderr, "%s: start page need to large than 0 %d\n", progarm, sp)
		pflag.Usage()
		os.Exit(3)
	}	
		//第三个参数检查
	if os.Args[3] != "-e" {
		fmt.Fprintf(os.Stderr, "%s: second arg need to be -e end_page\n", progarm)
		pflag.Usage()
		os.Exit(4)
	}
	
	ep,_ := strconv.Atoi(os.Args[4])
	if ep < 1  || ep < sp{
		fmt.Fprintf(os.Stderr, "%s: invalid end page %d\n", progarm, ep)
		pflag.Usage()
		os.Exit(5)
	}	

	//判断其他参数
	argindex := 5
	for{
		if argindex > len(os.Args)-1 || os.Args[argindex][0] != '-'{
			break
		}
		switch os.Args[argindex][1] {
		case 'l':
			pl, _ := strconv.Atoi(os.Args[argindex+1])
			fmt.Println(pl)
			if pl < 1 {
				fmt.Fprintf(os.Stderr, "%s: invaild page length %d\n", progarm, pl)
				pflag.Usage()
				os.Exit(3)
			}
			argindex += 2
		case 'f':
			if len(os.Args[argindex]) > 2 {
				fmt.Fprintf(os.Stderr, "%s: option should be -f\n", progarm)
				pflag.Usage()
				os.Exit(2)
			}
			argindex++

		case 'd':
			if len(os.Args[argindex]) <= 2 {
				fmt.Fprintf(os.Stderr, "%s: -d option requires a printer destinationination\n", progarm)
				os.Exit(1)
			}
			argindex += 2
		default:
			fmt.Fprintf(os.Stderr, "%s: unknown option", progarm)
			pflag.Usage()
			os.Exit(2)
		
		}
	}
}

func process_input(arg* selpg_args){
	fin := os.Stdin
	fout := os.Stdout//首先定义为标准输入输出
	var inpipe io.WriteCloser
    var err error
	//输入文件是否能成功打开
	if arg.file_in != ""{
		in, err := os.Open(arg.file_in)
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		defer in.Close()
		fin = in	
	}else{
		//fin := os.Stdin
	}//标准输入

	//检查输出
	if arg.file_out != ""{
		cmd := exec.Command("cat")
		inpipe, err = cmd.StdinPipe() //读取pipe
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		defer inpipe.Close()
		cmd.Stdout = fout
		cmd.Start()
	}else{
		//保持标准输出
	}

	current_page := 1
	current_line := 0
	//-l
	if arg.page_type == false{
		rf := bufio.NewScanner(fin)
		//读取文件
		for rf.Scan(){
			if current_page>=arg.start_page && current_page<=arg.end_page{
				fout.Write([]byte(rf.Text()+"\n"))
				if arg.file_out != ""{
					inpipe.Write([]byte(rf.Text()+"\n"))
				}
			}
			current_line ++
			//读完一页
			if current_line%arg.length_page == 0{
				current_line = 0
				current_page++
			}
		}
	}else{
		current_page = 1
		//fmt.Println("ffs", arg.page_type)
		rf := bufio.NewReader(fin)
		for {
			page, err := rf.ReadString('\f')
			if err != nil {
				break
			}
			if (current_page >= arg.start_page) && (current_page <= arg.end_page) {
				_, err := fout.Write([]byte(page))
				if err != nil {
					os.Exit(1)
				}
			}
			current_page++
		}
	}
}

func main()  {
	sa := new(selpg_args) //参数初始化
	analy_args(sa)
	progarm = os.Args[0]
	process_input(sa)
}
