拉取 - git clone:有无权限均可拉取,从0到1，本地无代码才能用
```bash
git clone <远程仓库地址.git>
git clone -b <远程分支名> <远程仓库地址.git>
```
拉取 - git pull:有权限才能拉取，更新代码时使用
```bash
git init
# 关联远程仓库，可关联多个仓库，origin会设置为仓库名
# 跳过这一步执行 git pull 也可以拉取
git remote add origin <远程仓库地址.git>
# 可选：设置多个远程地址，origin为仓库名
git remote set-url --add origin <远程仓库地址2.git>
# 拉取远程分支到某个本地分支
git pull <远程仓库地址.git> <远程分支名>:<本地分支名>
# 本地分支名可忽略，默认为拉取到当前分支
git pull <远程仓库地址.git> <远程分支名>
# 可选:移除远程仓库关联
git remote remove <远程仓库地址.git>
```
提交 - git push
```bash
# git初始化，添加 .git 文件
git init
# 将文件提交到暂存区
git add .
# commit
git commit -m <commit的内容>
# 可选:新建分支
git branch <分支名>
# 可选:切换到分支
git checkout <分支名>
# 提交前，如未与远程仓库关联则先与远程仓库关联
# 仓库名一般用 origin 
git remote add <仓库名> <远程仓库地址.git>
# 以当前分支提交到远程分支
git push -u <仓库名> <远程分支名>
# 可选:指定本地分支提交到远程分支
git push -u <仓库名> <本地分支名>:<远程分支名>
# 如出现冲突则先git pull拉取，解决冲突再提交
# git 拉取，没有指定分支则拉取默认分支
git pull <远程仓库地址.git> <远程分支名>
# 允许拉取非关联的仓库
git pull <远程仓库地址.git> <远程分支名> --allow-unrelated-histories
# 根据提示手动修改冲突文件
# 从git add开始重来 git 提交流程
```

删除
```bash
# 查看文件
dir
# 删除指定文件
git rm  <文件>
# 删除指定文件夹
git rm -r <文件夹>
git commit -m ""
git push
```
git add
```bash
#保存所有的修改,包括删除
git add -A  
# 保存新的添加和修改，但是不包括删除
git add . 
#保存修改和删除，但是不包括新建文件
git add -u  
```
覆盖
```bash
# 以远程仓库为准
# 拉取所有更新，不同步
git fetch --all
# 本地代码同步线上最新版本(会覆盖本地所有与远程仓库上同名的文件)；
git reset --hard origin/master

# 以本地仓库为准
# 强制推送到远程仓库，并覆盖远程代码库
git push -f --set-upstream origin master:master
```

合并
```bash
git merge <远程分支名> --allow-unrelated-histories
# 出现冲突，解决冲突
# 添加冲突文件
git add .
# 然后继续合并
git merge --continue
# 如果放弃合并则
git merge --abort
# 再推送
git push
```

迁移
```bash
# 裸仓库
git clone --bare xxx
# 或镜像仓库
git clone --mirror xxx

# 关联另一个仓库
git remote add remote2 xxx 
# 推送到另一个仓库
$ git push --mirror remote2
```
