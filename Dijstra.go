package main

import "fmt"

var Nodes = 6

var FindMaps = [9][3]int{
	{1,2,1},
	{1,3,12},
	{2,3,9},

	{2,4,3},
	{3,5,15},
	{4,3,4},

	{4,5,13},
	{4,6,15},
	{5,6,4},



}

var inf =999  //最大值代表不可联通

var min int

func main()  {
	var Maps[6][6]int

	for i:=0;i<Nodes;i++{
		for j:=0;j<Nodes;j++{
			if i==j{
				Maps[i][j]=0  //自己到自己的距离为0

			}else {
				Maps[i][j]=inf
			}
		}
	}


	for _,v := range FindMaps{
		Maps[v[0]-1][v[1]-1]=v[2]
	}

	for i:=0;i<Nodes;i++{
		fmt.Printf("%3d\n",Maps[i])
	}

	var dis[6]int //定义一个数组保存距离
	for k,v := range Maps[0]{
		dis[k]=v

	}
	fmt.Println(dis)

	var book[6]int
	book[0]=1
	var u int
	for i:=0;i<Nodes-1;i++{
		//寻找距离1号最近的点
		min = inf //设置当前最大距离
		for j:=0;j<Nodes;j++{
			if book[j]==0&&dis[j]<min{
				min = dis[j]
				u=j
			}
		}
		book[u]=1
		for v :=0 ; v<Nodes;v++{
			if Maps[u][v]!=inf && Maps[u][v]!=0{
				if dis[v] > dis[u] +Maps[u][v]{
					dis[v] = dis[u] +Maps[u][v] //保存最短距离
				}
			}
		}
	}
	fmt.Println(dis)
	fmt.Println(min)
}