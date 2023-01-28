---
tags: [1.Markdown语法]
title: 9.图形
created: '2023-01-23T16:47:00.635Z'
modified: '2023-01-23T18:43:17.078Z'
---

# 9.图形
1. 声明格式
   &#96;&#96;&#96;mermaid
      graph LR
      …
   &#96;&#96;&#96;
2. 声明图形方向（TB上到下，LR左到右,BT从下往上，RL从右往左）
   graph LR
3. 声明图标
   source(dataGenerator)
   mysql[mysqlDB]
   kafka[kafka]
   phoenix((phoenix))
   flinkDataSync[flinkDataSync]
   dataServer>dataServer]
   ```mermaid
    graph TB
        A[text] 
        B(text)
        C([text])
        D[(text)]
        E((text))
        F>text]
        G{text}
        H{{text}}
        I[/text/]
        J[\text\]
    ```
4. 自定义样式
   ```mermaid
   graph LR
    id1(Start)-->id2(Stop)
    style id1 fill:#f9f,stroke:#333,stroke-width:4px,fill-opacity:0.5
    style id2 fill:#ccf,stroke:#f66,stroke-width:2px,stroke-dasharray: 10,5
​   ```

4. 连线类型
   命令|形状|添加文本
   ---|:--:|---:
   A1 --> B1|直线箭头|--text-->
   A2 --- B2|直线|--text---
   A3 -.-> B3|虚线箭头|-.text.->
   A4 -.- B4|虚线|-.text.--
   A5 === B5|加粗直线|==text===
   A6 ==> B6|加粗直线箭头|==text==>

    ```mermaid
    graph TB
        A1 --> B1
        A2 --- B2
        A3 -.-> B3
        A4 -.- B4
        A5 === B5
        A6 ==> B6
    ```
    ----
    ```mermaid
    graph TB
    A1 --text--> B1
    A2 ---|text| B2
    A3 -.text.-> B3
    A4 -.-|text| B4
    A5 ==text=== B5
    A6 ==>|text| B6
    ```
   【注】可以使用 & 同时指定多个节点之间的多个连线。
    ```mermaid
    graph TB
    A --> D & E & F
    B & C -.-> F
    ```
5. 图形举例
   图形|符号|备注
   ---|:--:|---:
   矩形|name[label]|
   内凹五边形|name>label]|
   圆角矩形|name(label)|
   圆形|name((label))|
   菱形|name{label}|
   六边形|name{{label}}|
   梯形|name[\label/], name[/label]|
   平行四边形|name[\label], name[/label/]|
6. 描述逻辑
   source --> mysql --> kafka --> flinkDataSync --> phoenix ==> dataServer
   有连接的地方可以单独起一行
   flinkDataSync ==> clikhouse((clikhouse)) --> dataServer
   kafka --> flinkAnalyse[flinkAnalyse] --> kafka
7. e.g-1
   ```mermaid
   graph LR
   source(dataGenerator)
   mysql[mysqlDB]
   kafka[kafka]
   phoenix((phoenix))
   flinkDataSync[flinkDataSync]
   dataServer>dataServer]
   
   source --> mysql --flinkCDC.-> kafka --> flinkDataSync --> phoenix ==> dataServer
   flinkDataSync ==> clikhouse((clikhouse)) --> dataServer
   kafka --> flinkAnalyse[flinkAnalyse] --> kafka
   ```
8. 子流程
    ```mermaid
    graph TB
    subgraph one
        A1 --> B1
    end
    subgraph two
        A2 === B2
    end
    subgraph three
        A3 -.-> B2
    end
    ```
9. 注释
    ```mermaid
    graph LR
        A --> B %%此处为注释
    ```
10. 链接
    点击A会调整百度
    ```mermaid
    graph LR;
        A-->B;
        click A "https://www.baidu.com"
    ```
11. e.g-1
   ```mermaid
   graph TB
   Start --> End
   Start --> abc
   ```
12. e.g-2
   ```mermaid
   graph LR
   A1-1(Graph Type)-->A2-1(1:Round Rectangle)
   A2-1-.-A3-1[表示程序的开始或者结束]

   A1-1-->A2-2[2:Rectangle]
   A2-2-.-A3-2[一般用作要执行的处理]

   A1-1-->A2-3{3:菱形}
   A2-3-.-A3-3[表示决策或判断]

   A1-1-->A2-4>Particular shape]
   A2-4-.-A3-4[Unknown]

   A1-1-->A2-5((圆形))
   A2-5-.-A3-5[Usecase]
   ```
