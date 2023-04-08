在此之前，你应该保证已经安装好python了
# pyside6
```py
pip install pyside6
```
# vscode
插件 : Qt for Python

设置-扩展-Qt for Python :
`designer path`: D:\anaconda3\Lib\site-packages\PySide6\designer.exe
`rcc path`: D:\anaconda3\Scripts\pyside6-rcc.exe
`uic path`: D:\anaconda3\Scripts\pyside6-uic.exe

(因为我这个环境是有用anaconda的，选择pip install的位置(以上我的安装路径为D:\anaconda3)即可,)

也可以直接在setting.json里直接编辑:
```json
"qtForPython.designer.path": "D:\\anaconda3\\Lib\\site-packages\\PySide6\\designer.exe",
"qtForPython.rcc.path": "D:\\anaconda3\\Scripts\\pyside6-rcc.exe",
"qtForPython.uic.path": "D:\\anaconda3\\Scripts\\pyside6-uic.exe",
```
在vscode中的左侧导航栏右键选择Create Qt UI File(designer),可以打开qt designer

设置-扩展-Python:设置自动补全的额外路径
```json
"python.analysis.extraPaths": [
        "D:\\anaconda3\\Lib\\site-packages",
        "D:\\anaconda3\\Scripts"
    ],
```
# 最简可运行代码
```py
# 导入包
from PySide6.QtWidgets import QApplication, QWidget
# sys 包仅用于访问命令行参数，图形化界面一般用不上
import sys

# 每个应用程序只需要一个 QApplication 实例
# 传入 sys.argv 以允许应用程序使用命令行参数(也可以忽略，图形化界面一般用不上)
app = QApplication(sys.argv)

# 创建一个 Qt 部件, 作为我们的窗口.
window = QWidget()

# 显示窗口，窗口默认是隐藏的
window.show()  

# 启动事件循环
app.exec_()

# 当退出程序后事件结束，程序结束

```
# 自定义窗口
```py
from PySide6.QtWidgets import QApplication, QMainWindow, QPushButton, QLabel

# 创建 QMainWindow 的子类以自定义窗口
class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        # 设置窗口标题
        self.setWindowTitle("My App")

        # 按钮控件
        button = QPushButton("点我点我",self)
        # 接着就可以设置这个控件的属性
        button.setGeometry(0,0,200,200)
        button.setToolTip("这是一个按钮")
        button.setText("我被点了")

        # 同理，标签控件
        label = QLabel("welcome",self)
        label.setGeometry(200,200,300,300)
        label.setText("什么")


app = QApplication()

window = MainWindow()
window.show()

app.exec()
```
