package main
import "fmt"
func main(){
	//分析实现思路
	//1.统计一个班成绩情况，每个班有5名同学，求出该班的平均分【学生的成绩从键盘输入，先易后难】
	//2.学生数就5个，先死后活
	//3.声明一个sum 统计班级的总分

	//分析思路2
	//1.统计3个班的成绩情况，每个班有5名同学，求出每个班的平均分
	//2.j表示第几个班级
	//3.定义一个变量存放总成绩

	//分析思路3
	//1.我们可以把代码做活
	//2.定义两个变量，表示班级个数和班级人数

	//统计3个班的及格人数，每个班5名同学
	//分析思路
	//1.声明变量passCount用于保存及格人数

	//走进代码实现
	var classNum int = 2
	var stuNum int = 5
	var totalSum float64
	var passCount int = 0
	for j :=1; j <=classNum ; j++ {
	sum :=0.0
	var score float64
	for i :=1; i <=stuNum; i++ {
		fmt.Printf("输入第%d个班级第%d个学生成绩是: ",j,i)
		fmt.Scanln(&score)
		//累计总分
		sum += score
		//判断分数是否及格
		if score > 60 {
			passCount++
		}
	}
	fmt.Printf("第%d个班级平均分是%f\n",j,sum/float64(stuNum)) //定义的变量要用起来，此处输入变量名
	//将各个班级的总成绩累计到totalSum  
	totalSum += sum
}
fmt.Printf("3个班级的总成绩是:%v,所有班级平均分是%v\n",totalSum,totalSum/float64(stuNum*classNum))
fmt.Printf("及格人数为:%v",passCount)
}