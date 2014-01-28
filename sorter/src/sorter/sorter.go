package main

import "bufio" 
import "flag" 
import "fmt" 
import "io" 
import "os" 
import "strconv" 
import "time"
import "algorithms/bubblesort" 
import "algorithms/qsort"

var infile *string = flag.String("i", "infile", "File contains values for sorting") 
var outfile *string = flag.String("o", "outfile", "File to receive sorted values") 
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
    flag.Parse()

    if infile != nil {
        fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
    }

    values, err := readValues(*infile)
    if err == nil {
    	t1 := time.Now()
    	switch *algorithm {
    		case "qsort":
    			qsort.QuickSort(values)
    		case "bubblesort":
    			bubblesort.BubbleSort(values)
    		default:
    			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
    	}
    	t2 := time.Now()

    	fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")

    	writeValues(values, *outfile)
    } else {
    	fmt.Println(err)
    }
}

//创建并写入文件
func writeValues(values []int, outfile string) error {
    file, err := os.Create(outfile)
    if err != nil {
        fmt.Println("Failed to create the output file ", outfile)
        return err
    }
    defer file.Close()

    for _,value := range values {
        str := strconv.Itoa(value)
        file.WriteString(str + "\n")
    }
    return nil
}

//从文件读取数据
func readValues(infile string) (values []int, err error) {
    file, err := os.Open(infile)
    if err != nil {
        fmt.Println("Failed to open the input file ", infile)
        return 
    }

    defer file.Close()

    br := bufio.NewReader(file)
    
    for {
        //每次读取一行
        line, isPrefix, err1 := br.ReadLine()

        if err1 != nil {
            if err1 == io.EOF {
                break
            } else {
                err = err1
            }
        }

        //没找到行尾标记
        if isPrefix {
            fmt.Println("A too long line, seems unexpected.")
            return 
        }

        // 转换字符数组为字符串
        str := string(line) 
        
        //转为整型
        value, err2 := strconv.Atoi(str)
        if err2 != nil {
            err = err2
            return
        }
        values = append(values, value)
    }
    return
}