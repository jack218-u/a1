package main
import "fmt"
func main(){
	var a1,a2,a3,a4,a5 int
	// var b1,b2,b3,b4,b5 int
	// var c1,c2,c3,c4,c5 int
	// var d int
	for grade :=1; grade <=3; grade++ {
		if grade == 1 {
			fmt.Println("分别输入班级一a1到a5的成绩,用空格隔开")
			fmt.Scanf("%v, %v, %v, %v, %v",&a1,&a2,&a3,&a4,&a5)
			// d = a1 + a2 + a3 + a4 + a5
			fmt.Println(a1 + a2 + a3 + a4 + a5)
			// fmt.Printf("班级:%d,的平均分是",grade,d/5)
		// } else if grade == 2 {
		// 	fmt.Println("分别输入班级二b1到b5的成绩,用空格隔开")
		// 	fmt.Scanf("%v,%v,%v,%v,%v",&b1,&b2,&b3,&b4,&b5)
		// 	e = b1 + b2 + b3 + b4 + b5
		// 	fmt.Println("班级:%d,的平均分是",grade,e/5)
		// } else if grade == 3 {
		// 	fmt.Println("分别输入班级三c1到c5的成绩,用空格隔开")
		// 	fmt.Scanf("%v,%v,%v,%v,%v",&c1,&c2,&c3,&c4,&c5)
		// 	f = c1 + c2 + c3 + c4 + c5
		// 	fmt.Println("班级:%d,的平均分是",grade,f/5)
		// }
	}
}
}