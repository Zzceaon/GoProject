## DOS
- 查看目录 dir
- 新建目录 md
    md test1 test2
- 删除目录 rd
    re /q/s test
        /q: 不用询问
        /s: 删除包括这个目录的所有子目录
- 写入文本 echo >
    echo > a.txt: 新建不写入
    echo hello > d:\test\a.txt
        d:\test\a.txt: 写入到其他目录
    如果没有文件,则新建
- 拷贝文件 copy
    copy a.txt d:\test
        d:\test: 拷贝到其他目录
    copy a.txt d:\test\b.txt
        \b.txt: 重命名
- 移动/剪切 move
    move a.txt d:\test
- 删除 del
    del a.txt
    del *.txt: 删除所有后缀伪txt的文件
- 清屏 cls
- 退出DOS exit