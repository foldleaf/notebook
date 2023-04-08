# 转成python文件
在文件所在目录执行
```bash
# pyside6-uic <ui文件名> -o <生成的py文件名>
pyside6-uic computer.ui -o computer.py
```
在vscode中配置好Qt for Python插件后，只需要在ui文件处右键选择 `Compile Qt UI File(uic)` 即可转换
可以右键选择 Edit QT UI DILE(designer)，便捷地在qt designer里打开并编辑ui文件
# 使用生成的python文件
```python
from PySide6.QtWidgets import QApplication,QWidget
# 从生成的python文件中导入，导入的名称要与文件中的类名一致
from computer import Ui_Form

class MyWindow(QWidget):
    def __init__(self):
        super().__init__()

        # 使用生成的 ui
        self.ui=Ui_Form()
        self.ui.setupUi(self)

if __name__ == "__main__":
    app=QApplication()
    window=MyWindow()
    window.show()
    app.exec()
```
class的另一种多继承写法，更简洁:
```PY
class MyWindow(QWidget,Ui_Form):
    def __init__(self):
        super().__init__()
        self.setupUi(self)
```
