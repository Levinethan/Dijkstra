package main

import "fmt"

//矩阵实现迪捷斯卡拉算法

const MAX_SIZE int =5
const MAX_VALUE int = 9999

type Graph struct {
	vexs []string
	vexnum int
	edgnum int
	matrx [MAX_SIZE][MAX_SIZE]int  //临阶矩阵
	
}

type Edge struct {
	start string
	end  string
	weight int //权重
	
}

func Dij(g *Graph,start int)  {
	var dist[MAX_SIZE] int //路径长度数组
	var path[MAX_SIZE]bool  //最短路径节点集合
	dist = g.matrx[start]  //获取第一行


	path[start] = true //假定开始都都是最短的路线

	dist[start] = 0
	//从第一步开始

	for i:=0;i<g.vexnum;i++{
		k := 0 //备份数据
		min := MAX_VALUE
		for j:=0;j<len(dist);j++{
			//循环路径 贪心算法
			if path[j]==false && dist[j]!=MAX_VALUE && dist[j] < min{
				//不走回头路
				min = dist[j]
				k = j
			}
		}

		path[k] = true
		for u :=0 ; u <g.vexnum;u++{
			if g.matrx[k][u]!=MAX_VALUE && path[u] == false{
				weight := min + g.matrx[k][u]
				if weight<dist[u]{
					dist[u] = weight
				}
			}
		}

		for i:=0;i<g.vexnum;i++{
			fmt.Printf("path %s -> %s = %d\n",g.vexs[start],g.vexs[i],dist[i])
		}


	}
}

func initGraph(g *Graph)  {
	var vexs=[]string{"A","B","C","D","E"}
	g.vexnum = len(vexs)

	g.vexs = vexs  //设置顶点边长
	for i:=0;i<len(vexs);i++{
		for j:=0;j<len(vexs);j++{
			if i==j{
				g.matrx[i][j]=0
			}else {
				g.matrx[i][j]=MAX_VALUE
			}
		}
	}
	g.matrx[0][1] = 5
	g.matrx[0][2] = 3

	g.matrx[1][0] = 5
	g.matrx[1][3] = 99
	g.matrx[1][4] = 4

	g.matrx[2][0] = 3
	g.matrx[2][3] = 6
	g.matrx[3][1] = 7

	g.matrx[3][2] = 6
	g.matrx[3][4] = 1
	g.matrx[4][1] = 4
	g.matrx[4][3] = 5

	g.edgnum = 6 //设定边长为6
}
func PrintGraph(g Graph,length int)  {
	vlist := []string{" A "," B "," C "," D "," E "}
	fmt.Printf("%5s\n",vlist)
	for i:=0;i<length;i++{
		fmt.Printf("%5s",vlist[i])
		fmt.Printf("%5d\n",g.matrx[i])   //打印矩阵
	}
}
func main()  {
	var g Graph
	initGraph(&g)
	Dij(&g,1)
	PrintGraph(g,MAX_SIZE)
}